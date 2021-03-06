// Code generated by protoc-gen-gogo.
// source: struct.proto
// DO NOT EDIT!

/*
	Package struct_structpb is a generated protocol buffer package.

	It is generated from these files:
		struct.proto

	It has these top-level messages:
		Struct
		Value
		ListValue
*/
package struct_structpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import errors "errors"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//  The JSON representation for `NullValue` is JSON `null`.
type NullValue int32

const (
	// Null value.
	NullValue_NULL_VALUE NullValue = 0
)

var NullValue_name = map[int32]string{
	0: "NULL_VALUE",
}
var NullValue_value = map[string]int32{
	"NULL_VALUE": 0,
}

func (x NullValue) String() string {
	return proto.EnumName(NullValue_name, int32(x))
}
func (NullValue) EnumDescriptor() ([]byte, []int) { return fileDescriptorStruct, []int{0} }

// `Struct` represents a structured data value, consisting of fields
// which map to dynamically typed values. In some languages, `Struct`
// might be supported by a native representation. For example, in
// scripting languages like JS a struct is represented as an
// object. The details of that representation are described together
// with the proto support for the language.
//
// The JSON representation for `Struct` is JSON object.
type Struct struct {
	// Unordered map of dynamically typed values.
	Fields map[string]*Value `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Struct) Reset()                    { *m = Struct{} }
func (m *Struct) String() string            { return proto.CompactTextString(m) }
func (*Struct) ProtoMessage()               {}
func (*Struct) Descriptor() ([]byte, []int) { return fileDescriptorStruct, []int{0} }

func (m *Struct) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

// `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of that
// variants, absence of any variant indicates an error.
//
// The JSON representation for `Value` is JSON value.
type Value struct {
	// The kind of value.
	//
	// Types that are valid to be assigned to Kind:
	//	*Value_NullValue
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_StructValue
	//	*Value_ListValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptorStruct, []int{1} }

type isValue_Kind interface {
	isValue_Kind()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=sojern.protobuf.NullValue,oneof"`
}
type Value_NumberValue struct {
	NumberValue float64 `protobuf:"fixed64,2,opt,name=number_value,json=numberValue,proto3,oneof"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}
type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}
type Value_StructValue struct {
	StructValue *Struct `protobuf:"bytes,5,opt,name=struct_value,json=structValue,oneof"`
}
type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,6,opt,name=list_value,json=listValue,oneof"`
}

func (*Value_NullValue) isValue_Kind()   {}
func (*Value_NumberValue) isValue_Kind() {}
func (*Value_StringValue) isValue_Kind() {}
func (*Value_BoolValue) isValue_Kind()   {}
func (*Value_StructValue) isValue_Kind() {}
func (*Value_ListValue) isValue_Kind()   {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetNullValue() NullValue {
	if x, ok := m.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return NullValue_NULL_VALUE
}

func (m *Value) GetNumberValue() float64 {
	if x, ok := m.GetKind().(*Value_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetStructValue() *Struct {
	if x, ok := m.GetKind().(*Value_StructValue); ok {
		return x.StructValue
	}
	return nil
}

func (m *Value) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_NullValue)(nil),
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_ListValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.NullValue))
	case *Value_NumberValue:
		_ = b.EncodeVarint(2<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.NumberValue))
	case *Value_StringValue:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.StringValue)
	case *Value_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *Value_StructValue:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StructValue); err != nil {
			return err
		}
	case *Value_ListValue:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.Kind has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // kind.null_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_NullValue{NullValue(x)}
		return true, err
	case 2: // kind.number_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Kind = &Value_NumberValue{math.Float64frombits(x)}
		return true, err
	case 3: // kind.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &Value_StringValue{x}
		return true, err
	case 4: // kind.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_BoolValue{x != 0}
		return true, err
	case 5: // kind.struct_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Struct)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_StructValue{msg}
		return true, err
	case 6: // kind.list_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListValue)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_ListValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.NullValue))
	case *Value_NumberValue:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case *Value_StringValue:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BoolValue:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += 1
	case *Value_StructValue:
		s := proto.Size(x.StructValue)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ListValue:
		s := proto.Size(x.ListValue)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// `ListValue` is a wrapper around a repeated field of values.
//
// The JSON representation for `ListValue` is JSON array.
type ListValue struct {
	// Repeated field of dynamically typed values.
	Values []*Value `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *ListValue) Reset()                    { *m = ListValue{} }
func (m *ListValue) String() string            { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()               {}
func (*ListValue) Descriptor() ([]byte, []int) { return fileDescriptorStruct, []int{2} }

func (m *ListValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Struct)(nil), "sojern.protobuf.Struct")
	proto.RegisterType((*Value)(nil), "sojern.protobuf.Value")
	proto.RegisterType((*ListValue)(nil), "sojern.protobuf.ListValue")
	proto.RegisterEnum("sojern.protobuf.NullValue", NullValue_name, NullValue_value)
}
func (m *Struct) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Struct) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Fields) > 0 {
		for k, _ := range m.Fields {
			data[i] = 0xa
			i++
			v := m.Fields[k]
			if v == nil {
				return 0, errors.New("proto: map has nil element")
			}
			msgSize := v.Size()
			mapSize := 1 + len(k) + sovStruct(uint64(len(k))) + 1 + msgSize + sovStruct(uint64(msgSize))
			i = encodeVarintStruct(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintStruct(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintStruct(data, i, uint64(v.Size()))
			n1, err := v.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n1
		}
	}
	return i, nil
}

func (m *Value) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Value) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Kind != nil {
		nn2, err := m.Kind.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	return i, nil
}

func (m *Value_NullValue) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x8
	i++
	i = encodeVarintStruct(data, i, uint64(m.NullValue))
	return i, nil
}
func (m *Value_NumberValue) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x11
	i++
	i = encodeFixed64Struct(data, i, uint64(math.Float64bits(float64(m.NumberValue))))
	return i, nil
}
func (m *Value_StringValue) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x1a
	i++
	i = encodeVarintStruct(data, i, uint64(len(m.StringValue)))
	i += copy(data[i:], m.StringValue)
	return i, nil
}
func (m *Value_BoolValue) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x20
	i++
	if m.BoolValue {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}
func (m *Value_StructValue) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.StructValue != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintStruct(data, i, uint64(m.StructValue.Size()))
		n3, err := m.StructValue.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *Value_ListValue) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.ListValue != nil {
		data[i] = 0x32
		i++
		i = encodeVarintStruct(data, i, uint64(m.ListValue.Size()))
		n4, err := m.ListValue.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *ListValue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListValue) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			data[i] = 0xa
			i++
			i = encodeVarintStruct(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Struct(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Struct(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStruct(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Struct) Size() (n int) {
	var l int
	_ = l
	if len(m.Fields) > 0 {
		for k, v := range m.Fields {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
			}
			mapEntrySize := 1 + len(k) + sovStruct(uint64(len(k))) + 1 + l + sovStruct(uint64(l))
			n += mapEntrySize + 1 + sovStruct(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Value) Size() (n int) {
	var l int
	_ = l
	if m.Kind != nil {
		n += m.Kind.Size()
	}
	return n
}

func (m *Value_NullValue) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovStruct(uint64(m.NullValue))
	return n
}
func (m *Value_NumberValue) Size() (n int) {
	var l int
	_ = l
	n += 9
	return n
}
func (m *Value_StringValue) Size() (n int) {
	var l int
	_ = l
	l = len(m.StringValue)
	n += 1 + l + sovStruct(uint64(l))
	return n
}
func (m *Value_BoolValue) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}
func (m *Value_StructValue) Size() (n int) {
	var l int
	_ = l
	if m.StructValue != nil {
		l = m.StructValue.Size()
		n += 1 + l + sovStruct(uint64(l))
	}
	return n
}
func (m *Value_ListValue) Size() (n int) {
	var l int
	_ = l
	if m.ListValue != nil {
		l = m.ListValue.Size()
		n += 1 + l + sovStruct(uint64(l))
	}
	return n
}
func (m *ListValue) Size() (n int) {
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovStruct(uint64(l))
		}
	}
	return n
}

func sovStruct(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStruct(x uint64) (n int) {
	return sovStruct(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Struct) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStruct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Struct: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Struct: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthStruct
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapmsglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapmsglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if mapmsglen < 0 {
				return ErrInvalidLengthStruct
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthStruct
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &Value{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.Fields == nil {
				m.Fields = make(map[string]*Value)
			}
			m.Fields[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStruct(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStruct
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
func (m *Value) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStruct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullValue", wireType)
			}
			var v NullValue
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (NullValue(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Kind = &Value_NullValue{v}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberValue", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.Kind = &Value_NumberValue{float64(math.Float64frombits(v))}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = &Value_StringValue{string(data[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Kind = &Value_BoolValue{b}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StructValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Struct{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Kind = &Value_StructValue{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ListValue{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Kind = &Value_ListValue{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStruct(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStruct
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
func (m *ListValue) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStruct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &Value{})
			if err := m.Values[len(m.Values)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStruct(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStruct
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
func skipStruct(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStruct
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStruct
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStruct
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipStruct(data[start:])
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
	ErrInvalidLengthStruct = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStruct   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorStruct = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x8b, 0xda, 0x40,
	0x18, 0xc6, 0x9d, 0x44, 0x43, 0xf3, 0x46, 0x54, 0x52, 0x68, 0xc5, 0x82, 0x2d, 0x7a, 0x91, 0x52,
	0x72, 0xb0, 0x97, 0x52, 0x7b, 0x69, 0xc0, 0x5a, 0x68, 0x90, 0x34, 0xad, 0x16, 0x7a, 0x91, 0x46,
	0x47, 0x49, 0x1d, 0x67, 0x24, 0x7f, 0x5a, 0xfc, 0x26, 0x7b, 0x5c, 0xf6, 0xb8, 0xc7, 0xfd, 0x14,
	0x7b, 0xdc, 0x8f, 0xb0, 0xec, 0x27, 0xd9, 0xf9, 0x93, 0x64, 0x97, 0x15, 0x0f, 0x83, 0xf3, 0x3e,
	0xfe, 0xde, 0x67, 0xde, 0xe7, 0x0d, 0xd4, 0x93, 0x34, 0xce, 0x96, 0xa9, 0xb3, 0x8f, 0x59, 0xca,
	0xec, 0x66, 0xc2, 0xfe, 0xe2, 0x98, 0xaa, 0x2a, 0xcc, 0xd6, 0xbd, 0x33, 0x04, 0xc6, 0x0f, 0x49,
	0xd8, 0x23, 0x30, 0xd6, 0x11, 0x26, 0xab, 0xa4, 0x8d, 0xde, 0xe8, 0x03, 0x6b, 0xd8, 0x77, 0x9e,
	0xc0, 0x8e, 0x02, 0x9d, 0x2f, 0x92, 0x1a, 0xd3, 0x34, 0x3e, 0x04, 0x79, 0x4b, 0xe7, 0x3b, 0x58,
	0x8f, 0x64, 0xbb, 0x05, 0xfa, 0x16, 0x1f, 0xb8, 0x11, 0x1a, 0x98, 0x81, 0xb8, 0xda, 0xef, 0xa0,
	0xf6, 0xef, 0x0f, 0xc9, 0x70, 0x5b, 0xe3, 0x9a, 0x35, 0x7c, 0x71, 0x64, 0x3e, 0x17, 0xff, 0x06,
	0x0a, 0xfa, 0xa8, 0x7d, 0x40, 0xbd, 0x2b, 0x0d, 0x6a, 0x52, 0xe4, 0x93, 0x01, 0xcd, 0x08, 0x59,
	0x28, 0x03, 0x61, 0xda, 0x18, 0x76, 0x8e, 0x0c, 0xa6, 0x1c, 0x91, 0xfc, 0xd7, 0x4a, 0x60, 0xd2,
	0xa2, 0xb0, 0xfb, 0x50, 0xa7, 0xd9, 0x2e, 0xc4, 0xf1, 0xe2, 0xe1, 0x7d, 0xc4, 0x11, 0x4b, 0xa9,
	0x25, 0xc4, 0xf7, 0x14, 0xd1, 0x4d, 0x0e, 0xe9, 0x62, 0x70, 0x01, 0x29, 0x55, 0x41, 0xaf, 0x01,
	0x42, 0xc6, 0x8a, 0x31, 0xaa, 0x1c, 0x79, 0x26, 0x9e, 0x12, 0x9a, 0x02, 0x3e, 0x15, 0xdb, 0xce,
	0x91, 0x9a, 0x8c, 0xfa, 0xf2, 0xc4, 0x1e, 0x73, 0x7b, 0x7e, 0x2b, 0x53, 0x92, 0x28, 0x29, 0x7a,
	0x0d, 0xd9, 0x7b, 0x9c, 0xd2, 0xe3, 0x48, 0x99, 0x92, 0x14, 0x85, 0x6b, 0x40, 0x75, 0x1b, 0xd1,
	0x55, 0x6f, 0x04, 0x66, 0x49, 0xd8, 0x0e, 0x18, 0xd2, 0xac, 0xf8, 0xa2, 0xa7, 0x96, 0x9e, 0x53,
	0x6f, 0x5f, 0x81, 0x59, 0x2e, 0xd1, 0x6e, 0x00, 0x4c, 0x67, 0x9e, 0xb7, 0x98, 0x7f, 0xf6, 0x66,
	0xe3, 0x56, 0xc5, 0xdd, 0x5e, 0xdf, 0x75, 0xd1, 0x0d, 0x3f, 0xb7, 0xfc, 0xc0, 0xf3, 0x25, 0xdb,
	0x39, 0x1b, 0xc6, 0x36, 0x04, 0x97, 0x6e, 0xae, 0xa5, 0x82, 0xf9, 0xa2, 0xf6, 0xd1, 0xef, 0xa6,
	0x4a, 0x37, 0x52, 0x3f, 0xfb, 0xf0, 0x1c, 0xa1, 0x0b, 0x4d, 0x9f, 0xf8, 0xee, 0xa5, 0xd6, 0x9d,
	0xa8, 0x56, 0xbf, 0x18, 0xe4, 0x17, 0x26, 0xe4, 0x1b, 0x65, 0xff, 0xe9, 0xcf, 0xc3, 0x1e, 0x27,
	0xa1, 0x21, 0x3d, 0xdf, 0xdf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x83, 0x5e, 0x9f, 0xbe, 0x02,
	0x00, 0x00,
}
