// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bulkloader.proto

/*
	Package protos is a generated protocol buffer package.

	It is generated from these files:
		bulkloader.proto
		facets.proto
		graphresponse.proto
		payload.proto
		schema.proto
		task.proto
		types.proto

	It has these top-level messages:
		DenormalisedPosting
		Facet
		Param
		Facets
		FacetsList
		Function
		FilterTree
		Num
		AssignedIds
		NQuad
		Value
		Mutation
		Request
		Latency
		Property
		Node
		Response
		Check
		Version
		Payload
		ExportPayload
		SchemaRequest
		SchemaResult
		SchemaNode
		SchemaUpdate
		List
		TaskValue
		Query
		ValueList
		Result
		SortMessage
		SortResult
		RaftContext
		Membership
		MembershipUpdate
		DirectedEdge
		Mutations
		Proposal
		KV
		KC
		GroupKeys
		Posting
		PostingList
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DenormalisedPosting struct {
	PostingListKey []byte `protobuf:"bytes,1,opt,name=posting_list_key,json=postingListKey,proto3" json:"posting_list_key,omitempty"`
	// Types that are valid to be assigned to Posting:
	//	*DenormalisedPosting_UidPosting
	//	*DenormalisedPosting_FullPosting
	Posting isDenormalisedPosting_Posting `protobuf_oneof:"posting"`
}

func (m *DenormalisedPosting) Reset()                    { *m = DenormalisedPosting{} }
func (m *DenormalisedPosting) String() string            { return proto.CompactTextString(m) }
func (*DenormalisedPosting) ProtoMessage()               {}
func (*DenormalisedPosting) Descriptor() ([]byte, []int) { return fileDescriptorBulkloader, []int{0} }

type isDenormalisedPosting_Posting interface {
	isDenormalisedPosting_Posting()
	MarshalTo([]byte) (int, error)
	Size() int
}

type DenormalisedPosting_UidPosting struct {
	UidPosting uint64 `protobuf:"fixed64,2,opt,name=uid_posting,json=uidPosting,proto3,oneof"`
}
type DenormalisedPosting_FullPosting struct {
	FullPosting *Posting `protobuf:"bytes,3,opt,name=full_posting,json=fullPosting,oneof"`
}

func (*DenormalisedPosting_UidPosting) isDenormalisedPosting_Posting()  {}
func (*DenormalisedPosting_FullPosting) isDenormalisedPosting_Posting() {}

func (m *DenormalisedPosting) GetPosting() isDenormalisedPosting_Posting {
	if m != nil {
		return m.Posting
	}
	return nil
}

func (m *DenormalisedPosting) GetPostingListKey() []byte {
	if m != nil {
		return m.PostingListKey
	}
	return nil
}

func (m *DenormalisedPosting) GetUidPosting() uint64 {
	if x, ok := m.GetPosting().(*DenormalisedPosting_UidPosting); ok {
		return x.UidPosting
	}
	return 0
}

func (m *DenormalisedPosting) GetFullPosting() *Posting {
	if x, ok := m.GetPosting().(*DenormalisedPosting_FullPosting); ok {
		return x.FullPosting
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DenormalisedPosting) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DenormalisedPosting_OneofMarshaler, _DenormalisedPosting_OneofUnmarshaler, _DenormalisedPosting_OneofSizer, []interface{}{
		(*DenormalisedPosting_UidPosting)(nil),
		(*DenormalisedPosting_FullPosting)(nil),
	}
}

func _DenormalisedPosting_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DenormalisedPosting)
	// posting
	switch x := m.Posting.(type) {
	case *DenormalisedPosting_UidPosting:
		_ = b.EncodeVarint(2<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(uint64(x.UidPosting))
	case *DenormalisedPosting_FullPosting:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FullPosting); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DenormalisedPosting.Posting has unexpected type %T", x)
	}
	return nil
}

func _DenormalisedPosting_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DenormalisedPosting)
	switch tag {
	case 2: // posting.uid_posting
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Posting = &DenormalisedPosting_UidPosting{x}
		return true, err
	case 3: // posting.full_posting
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Posting)
		err := b.DecodeMessage(msg)
		m.Posting = &DenormalisedPosting_FullPosting{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DenormalisedPosting_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DenormalisedPosting)
	// posting
	switch x := m.Posting.(type) {
	case *DenormalisedPosting_UidPosting:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case *DenormalisedPosting_FullPosting:
		s := proto.Size(x.FullPosting)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*DenormalisedPosting)(nil), "protos.DenormalisedPosting")
}
func (m *DenormalisedPosting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenormalisedPosting) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PostingListKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBulkloader(dAtA, i, uint64(len(m.PostingListKey)))
		i += copy(dAtA[i:], m.PostingListKey)
	}
	if m.Posting != nil {
		nn1, err := m.Posting.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *DenormalisedPosting_UidPosting) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x11
	i++
	i = encodeFixed64Bulkloader(dAtA, i, uint64(m.UidPosting))
	return i, nil
}
func (m *DenormalisedPosting_FullPosting) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.FullPosting != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBulkloader(dAtA, i, uint64(m.FullPosting.Size()))
		n2, err := m.FullPosting.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func encodeFixed64Bulkloader(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Bulkloader(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintBulkloader(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DenormalisedPosting) Size() (n int) {
	var l int
	_ = l
	l = len(m.PostingListKey)
	if l > 0 {
		n += 1 + l + sovBulkloader(uint64(l))
	}
	if m.Posting != nil {
		n += m.Posting.Size()
	}
	return n
}

func (m *DenormalisedPosting_UidPosting) Size() (n int) {
	var l int
	_ = l
	n += 9
	return n
}
func (m *DenormalisedPosting_FullPosting) Size() (n int) {
	var l int
	_ = l
	if m.FullPosting != nil {
		l = m.FullPosting.Size()
		n += 1 + l + sovBulkloader(uint64(l))
	}
	return n
}

func sovBulkloader(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBulkloader(x uint64) (n int) {
	return sovBulkloader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenormalisedPosting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBulkloader
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
			return fmt.Errorf("proto: DenormalisedPosting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenormalisedPosting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostingListKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBulkloader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBulkloader
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostingListKey = append(m.PostingListKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PostingListKey == nil {
				m.PostingListKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field UidPosting", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Posting = &DenormalisedPosting_UidPosting{v}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullPosting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBulkloader
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
				return ErrInvalidLengthBulkloader
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Posting{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Posting = &DenormalisedPosting_FullPosting{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBulkloader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBulkloader
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
func skipBulkloader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBulkloader
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
					return 0, ErrIntOverflowBulkloader
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
					return 0, ErrIntOverflowBulkloader
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
				return 0, ErrInvalidLengthBulkloader
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBulkloader
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
				next, err := skipBulkloader(dAtA[start:])
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
	ErrInvalidLengthBulkloader = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBulkloader   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("bulkloader.proto", fileDescriptorBulkloader) }

var fileDescriptorBulkloader = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2a, 0xcd, 0xc9,
	0xce, 0xc9, 0x4f, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53,
	0xc5, 0x52, 0xdc, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x10, 0x41, 0xa5, 0xc5, 0x8c, 0x5c, 0xc2, 0x2e,
	0xa9, 0x79, 0xf9, 0x45, 0xb9, 0x89, 0x39, 0x99, 0xc5, 0xa9, 0x29, 0x01, 0xf9, 0xc5, 0x25, 0x99,
	0x79, 0xe9, 0x42, 0x1a, 0x5c, 0x02, 0x05, 0x10, 0x66, 0x7c, 0x4e, 0x66, 0x71, 0x49, 0x7c, 0x76,
	0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x1f, 0x54, 0xdc, 0x27, 0xb3, 0xb8, 0xc4,
	0x3b, 0xb5, 0x52, 0x48, 0x91, 0x8b, 0xbb, 0x34, 0x33, 0x25, 0x1e, 0x2a, 0x2a, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0xe6, 0xc1, 0x10, 0xc4, 0x55, 0x9a, 0x09, 0x37, 0xcc, 0x84, 0x8b, 0x27, 0xad, 0x34,
	0x27, 0x07, 0xae, 0x86, 0x59, 0x81, 0x51, 0x83, 0xdb, 0x88, 0x1f, 0xe2, 0x84, 0x62, 0x3d, 0xa8,
	0x32, 0x0f, 0x86, 0x20, 0x6e, 0x90, 0x32, 0x28, 0xd7, 0x89, 0x93, 0x8b, 0x1d, 0xaa, 0xc1, 0x49,
	0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58,
	0x8e, 0x21, 0x09, 0xe2, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xe4, 0x31, 0x24,
	0xe7, 0x00, 0x00, 0x00,
}
