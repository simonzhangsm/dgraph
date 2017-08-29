// Code generated by protoc-gen-gogo.
// source: schema.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SchemaUpdate_Directive int32

const (
	SchemaUpdate_NONE    SchemaUpdate_Directive = 0
	SchemaUpdate_INDEX   SchemaUpdate_Directive = 1
	SchemaUpdate_REVERSE SchemaUpdate_Directive = 2
)

var SchemaUpdate_Directive_name = map[int32]string{
	0: "NONE",
	1: "INDEX",
	2: "REVERSE",
}
var SchemaUpdate_Directive_value = map[string]int32{
	"NONE":    0,
	"INDEX":   1,
	"REVERSE": 2,
}

func (x SchemaUpdate_Directive) String() string {
	return proto.EnumName(SchemaUpdate_Directive_name, int32(x))
}
func (SchemaUpdate_Directive) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorSchema, []int{3, 0}
}

type SchemaRequest struct {
	GroupId    uint32   `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Predicates []string `protobuf:"bytes,2,rep,name=predicates" json:"predicates,omitempty"`
	// fields can be on of type, index, reverse or tokenizer
	Fields []string `protobuf:"bytes,3,rep,name=fields" json:"fields,omitempty"`
}

func (m *SchemaRequest) Reset()                    { *m = SchemaRequest{} }
func (m *SchemaRequest) String() string            { return proto.CompactTextString(m) }
func (*SchemaRequest) ProtoMessage()               {}
func (*SchemaRequest) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{0} }

func (m *SchemaRequest) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *SchemaRequest) GetPredicates() []string {
	if m != nil {
		return m.Predicates
	}
	return nil
}

func (m *SchemaRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type SchemaResult struct {
	Schema []*SchemaNode `protobuf:"bytes,1,rep,name=schema" json:"schema,omitempty"`
}

func (m *SchemaResult) Reset()                    { *m = SchemaResult{} }
func (m *SchemaResult) String() string            { return proto.CompactTextString(m) }
func (*SchemaResult) ProtoMessage()               {}
func (*SchemaResult) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{1} }

func (m *SchemaResult) GetSchema() []*SchemaNode {
	if m != nil {
		return m.Schema
	}
	return nil
}

type SchemaNode struct {
	Predicate string   `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Type      string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Index     bool     `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Tokenizer []string `protobuf:"bytes,4,rep,name=tokenizer" json:"tokenizer,omitempty"`
	Reverse   bool     `protobuf:"varint,5,opt,name=reverse,proto3" json:"reverse,omitempty"`
	Count     bool     `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	List      bool     `protobuf:"varint,7,opt,name=list,proto3" json:"list,omitempty"`
}

func (m *SchemaNode) Reset()                    { *m = SchemaNode{} }
func (m *SchemaNode) String() string            { return proto.CompactTextString(m) }
func (*SchemaNode) ProtoMessage()               {}
func (*SchemaNode) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{2} }

func (m *SchemaNode) GetPredicate() string {
	if m != nil {
		return m.Predicate
	}
	return ""
}

func (m *SchemaNode) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SchemaNode) GetIndex() bool {
	if m != nil {
		return m.Index
	}
	return false
}

func (m *SchemaNode) GetTokenizer() []string {
	if m != nil {
		return m.Tokenizer
	}
	return nil
}

func (m *SchemaNode) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

func (m *SchemaNode) GetCount() bool {
	if m != nil {
		return m.Count
	}
	return false
}

func (m *SchemaNode) GetList() bool {
	if m != nil {
		return m.List
	}
	return false
}

type SchemaUpdate struct {
	Predicate string                 `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
	ValueType uint32                 `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3" json:"value_type,omitempty"`
	Directive SchemaUpdate_Directive `protobuf:"varint,3,opt,name=directive,proto3,enum=protos.SchemaUpdate_Directive" json:"directive,omitempty"`
	Tokenizer []string               `protobuf:"bytes,4,rep,name=tokenizer" json:"tokenizer,omitempty"`
	Count     bool                   `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	List      bool                   `protobuf:"varint,6,opt,name=list,proto3" json:"list,omitempty"`
}

func (m *SchemaUpdate) Reset()                    { *m = SchemaUpdate{} }
func (m *SchemaUpdate) String() string            { return proto.CompactTextString(m) }
func (*SchemaUpdate) ProtoMessage()               {}
func (*SchemaUpdate) Descriptor() ([]byte, []int) { return fileDescriptorSchema, []int{3} }

func (m *SchemaUpdate) GetPredicate() string {
	if m != nil {
		return m.Predicate
	}
	return ""
}

func (m *SchemaUpdate) GetValueType() uint32 {
	if m != nil {
		return m.ValueType
	}
	return 0
}

func (m *SchemaUpdate) GetDirective() SchemaUpdate_Directive {
	if m != nil {
		return m.Directive
	}
	return SchemaUpdate_NONE
}

func (m *SchemaUpdate) GetTokenizer() []string {
	if m != nil {
		return m.Tokenizer
	}
	return nil
}

func (m *SchemaUpdate) GetCount() bool {
	if m != nil {
		return m.Count
	}
	return false
}

func (m *SchemaUpdate) GetList() bool {
	if m != nil {
		return m.List
	}
	return false
}

func init() {
	proto.RegisterType((*SchemaRequest)(nil), "protos.SchemaRequest")
	proto.RegisterType((*SchemaResult)(nil), "protos.SchemaResult")
	proto.RegisterType((*SchemaNode)(nil), "protos.SchemaNode")
	proto.RegisterType((*SchemaUpdate)(nil), "protos.SchemaUpdate")
	proto.RegisterEnum("protos.SchemaUpdate_Directive", SchemaUpdate_Directive_name, SchemaUpdate_Directive_value)
}
func (m *SchemaRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GroupId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.GroupId))
	}
	if len(m.Predicates) > 0 {
		for _, s := range m.Predicates {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Fields) > 0 {
		for _, s := range m.Fields {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *SchemaResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Schema) > 0 {
		for _, msg := range m.Schema {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSchema(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SchemaNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaNode) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Predicate) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Predicate)))
		i += copy(dAtA[i:], m.Predicate)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Index {
		dAtA[i] = 0x18
		i++
		if m.Index {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Tokenizer) > 0 {
		for _, s := range m.Tokenizer {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Reverse {
		dAtA[i] = 0x28
		i++
		if m.Reverse {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Count {
		dAtA[i] = 0x30
		i++
		if m.Count {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.List {
		dAtA[i] = 0x38
		i++
		if m.List {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *SchemaUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SchemaUpdate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Predicate) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Predicate)))
		i += copy(dAtA[i:], m.Predicate)
	}
	if m.ValueType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.ValueType))
	}
	if m.Directive != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSchema(dAtA, i, uint64(m.Directive))
	}
	if len(m.Tokenizer) > 0 {
		for _, s := range m.Tokenizer {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.Count {
		dAtA[i] = 0x28
		i++
		if m.Count {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.List {
		dAtA[i] = 0x30
		i++
		if m.List {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Schema(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Schema(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SchemaRequest) Size() (n int) {
	var l int
	_ = l
	if m.GroupId != 0 {
		n += 1 + sovSchema(uint64(m.GroupId))
	}
	if len(m.Predicates) > 0 {
		for _, s := range m.Predicates {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if len(m.Fields) > 0 {
		for _, s := range m.Fields {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	return n
}

func (m *SchemaResult) Size() (n int) {
	var l int
	_ = l
	if len(m.Schema) > 0 {
		for _, e := range m.Schema {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	return n
}

func (m *SchemaNode) Size() (n int) {
	var l int
	_ = l
	l = len(m.Predicate)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.Index {
		n += 2
	}
	if len(m.Tokenizer) > 0 {
		for _, s := range m.Tokenizer {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.Reverse {
		n += 2
	}
	if m.Count {
		n += 2
	}
	if m.List {
		n += 2
	}
	return n
}

func (m *SchemaUpdate) Size() (n int) {
	var l int
	_ = l
	l = len(m.Predicate)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.ValueType != 0 {
		n += 1 + sovSchema(uint64(m.ValueType))
	}
	if m.Directive != 0 {
		n += 1 + sovSchema(uint64(m.Directive))
	}
	if len(m.Tokenizer) > 0 {
		for _, s := range m.Tokenizer {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.Count {
		n += 2
	}
	if m.List {
		n += 2
	}
	return n
}

func sovSchema(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SchemaRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Predicates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Predicates = append(m.Predicates, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fields = append(m.Fields, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SchemaResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = append(m.Schema, &SchemaNode{})
			if err := m.Schema[len(m.Schema)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SchemaNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Predicate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Predicate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Index = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokenizer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokenizer = append(m.Tokenizer, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reverse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Reverse = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Count = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.List = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SchemaUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SchemaUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SchemaUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Predicate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Predicate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueType", wireType)
			}
			m.ValueType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValueType |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Directive", wireType)
			}
			m.Directive = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Directive |= (SchemaUpdate_Directive(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokenizer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokenizer = append(m.Tokenizer, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Count = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.List = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSchema
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSchema
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSchema(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSchema = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("schema.proto", fileDescriptorSchema) }

var fileDescriptorSchema = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xcd, 0xee, 0x93, 0x40,
	0x14, 0xc5, 0xff, 0x03, 0x05, 0xca, 0x6d, 0x6b, 0xc8, 0xc4, 0x98, 0x31, 0x51, 0x42, 0x58, 0x11,
	0x13, 0xbb, 0xa8, 0x3b, 0xe3, 0xca, 0x94, 0x45, 0x37, 0x98, 0x4c, 0xd5, 0xb8, 0x6b, 0x5a, 0xe6,
	0xaa, 0x13, 0xb1, 0x20, 0x33, 0x34, 0xea, 0xca, 0xc7, 0xf0, 0x49, 0x7c, 0x06, 0x97, 0x3e, 0x82,
	0xa9, 0x2f, 0x62, 0x18, 0x0a, 0xd4, 0xc4, 0xe8, 0x8a, 0x39, 0x67, 0xee, 0xc7, 0xf9, 0x91, 0x81,
	0xb9, 0xca, 0xdf, 0xe2, 0xfb, 0xfd, 0xb2, 0xaa, 0x4b, 0x5d, 0x52, 0xd7, 0x7c, 0x54, 0x7c, 0x80,
	0xc5, 0xd6, 0xf8, 0x1c, 0x3f, 0x34, 0xa8, 0x34, 0xbd, 0x0b, 0xd3, 0x37, 0x75, 0xd9, 0x54, 0x3b,
	0x29, 0x18, 0x89, 0x48, 0xb2, 0xe0, 0x9e, 0xd1, 0x1b, 0x41, 0x43, 0x80, 0xaa, 0x46, 0x21, 0xf3,
	0xbd, 0x46, 0xc5, 0xac, 0xc8, 0x4e, 0x7c, 0x7e, 0xe5, 0xd0, 0x3b, 0xe0, 0xbe, 0x96, 0x58, 0x08,
	0xc5, 0x6c, 0x73, 0x77, 0x51, 0xf1, 0x63, 0x98, 0xf7, 0x3b, 0x54, 0x53, 0x68, 0xfa, 0x00, 0xdc,
	0x2e, 0x0b, 0x23, 0x91, 0x9d, 0xcc, 0x56, 0xb4, 0xcb, 0xa4, 0x96, 0x5d, 0x55, 0x56, 0x0a, 0xe4,
	0x97, 0x8a, 0xf8, 0x1b, 0x01, 0x18, 0x6d, 0x7a, 0x0f, 0xfc, 0x61, 0xa1, 0x89, 0xe7, 0xf3, 0xd1,
	0xa0, 0x14, 0x26, 0xfa, 0x53, 0x85, 0xcc, 0x32, 0x17, 0xe6, 0x4c, 0x6f, 0x83, 0x23, 0x8f, 0x02,
	0x3f, 0x32, 0x3b, 0x22, 0xc9, 0x94, 0x77, 0xa2, 0x9d, 0xa3, 0xcb, 0x77, 0x78, 0x94, 0x9f, 0xb1,
	0x66, 0x13, 0x93, 0x76, 0x34, 0x28, 0x03, 0xaf, 0xc6, 0x13, 0xd6, 0x0a, 0x99, 0x63, 0xba, 0x7a,
	0xd9, 0x4e, 0xcb, 0xcb, 0xe6, 0xa8, 0x99, 0xdb, 0x4d, 0x33, 0xa2, 0xdd, 0x5b, 0x48, 0xa5, 0x99,
	0x67, 0x4c, 0x73, 0x8e, 0xbf, 0x58, 0x3d, 0xf5, 0x8b, 0x4a, 0xb4, 0xe1, 0xfe, 0x1d, 0xfd, 0x3e,
	0xc0, 0x69, 0x5f, 0x34, 0xb8, 0x1b, 0x00, 0x16, 0xdc, 0x37, 0xce, 0xf3, 0x96, 0xe2, 0x09, 0xf8,
	0x42, 0xd6, 0x98, 0x6b, 0x79, 0x42, 0x43, 0x72, 0x6b, 0x15, 0xfe, 0xf9, 0xd7, 0xba, 0x2d, 0xcb,
	0x75, 0x5f, 0xc5, 0xc7, 0x86, 0xff, 0xd0, 0x0e, 0x4c, 0xce, 0xdf, 0x98, 0xdc, 0x2b, 0xa6, 0x87,
	0xe0, 0x0f, 0xf3, 0xe9, 0x14, 0x26, 0xd9, 0xb3, 0x2c, 0x0d, 0x6e, 0xa8, 0x0f, 0xce, 0x26, 0x5b,
	0xa7, 0xaf, 0x02, 0x42, 0x67, 0xe0, 0xf1, 0xf4, 0x65, 0xca, 0xb7, 0x69, 0x60, 0x3d, 0x0d, 0xbe,
	0x9f, 0x43, 0xf2, 0xe3, 0x1c, 0x92, 0x9f, 0xe7, 0x90, 0x7c, 0xfd, 0x15, 0xde, 0x1c, 0xba, 0x57,
	0xf7, 0xe8, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x75, 0xf4, 0x5a, 0x8c, 0x02, 0x00, 0x00,
}
