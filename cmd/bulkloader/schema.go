package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/types"
	wk "github.com/dgraph-io/dgraph/worker"
)

type schemaState struct {
	strict bool
	*protos.SchemaUpdate
}

type schemaStore struct {
	mu sync.Mutex
	m  map[string]schemaState
}

func newSchemaStore(initial []*protos.SchemaUpdate) *schemaStore {
	s := &schemaStore{
		mu: sync.Mutex{},
		m:  map[string]schemaState{},
	}
	for _, sch := range initial {
		p := sch.Predicate
		sch.Predicate = ""
		s.m[p] = schemaState{true, sch}
	}
	return s
}

func (s *schemaStore) getSchema(pred string) *protos.SchemaUpdate {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.m[pred].SchemaUpdate
}

func (s *schemaStore) fixEdge(de *protos.DirectedEdge, objectIsUID bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if objectIsUID {
		de.ValueType = uint32(protos.Posting_UID)
	}

	sch, ok := s.m[de.Attr]
	if !ok {
		sch = schemaState{false, &protos.SchemaUpdate{ValueType: de.ValueType}}
		s.m[de.Attr] = sch
	}

	schTyp := types.TypeID(sch.ValueType)
	err := wk.ValidateAndConvert(de, schTyp)
	if sch.strict && err != nil {
		// TODO: It's unclear to me as to why it's only an error to have a bad
		// conversion if the schema was established explicitly rather than
		// automatically.
		//
		// TODO: Better error message
		fmt.Printf("RDF doesn't match schema: %v\n", err) // TODO: bubble back properly
		os.Exit(1)
	}
}
