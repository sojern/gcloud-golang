// Code generated by protoc-gen-gogo.
// source: table.proto
// DO NOT EDIT!

/*
	Package google_bigtable_admin_v2 is a generated protocol buffer package.

	It is generated from these files:
		table.proto

	It has these top-level messages:
		Table
		ColumnFamily
		GcRule
*/
package google_bigtable_admin_v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import sojern_protobuf "github.com/sojern/gcloud-golang/protobuf/duration"

import errors "errors"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// Possible timestamp granularities to use when keeping multiple versions
// of data in a table.
type Table_TimestampGranularity int32

const (
	// The user did not specify a granularity. Should not be returned.
	// When specified during table creation, MILLIS will be used.
	Table_TIMESTAMP_GRANULARITY_UNSPECIFIED Table_TimestampGranularity = 0
	// The table keeps data versioned at a granularity of 1ms.
	Table_MILLIS Table_TimestampGranularity = 1
)

var Table_TimestampGranularity_name = map[int32]string{
	0: "TIMESTAMP_GRANULARITY_UNSPECIFIED",
	1: "MILLIS",
}
var Table_TimestampGranularity_value = map[string]int32{
	"TIMESTAMP_GRANULARITY_UNSPECIFIED": 0,
	"MILLIS": 1,
}

func (x Table_TimestampGranularity) String() string {
	return proto.EnumName(Table_TimestampGranularity_name, int32(x))
}
func (Table_TimestampGranularity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTable, []int{0, 0}
}

// Defines a view over a table's fields.
type Table_View int32

const (
	// Uses the default view for each method as documented in its request.
	Table_VIEW_UNSPECIFIED Table_View = 0
	// Only populates `name`.
	Table_NAME_ONLY Table_View = 1
	// Only populates `name` and fields related to the table's schema.
	Table_SCHEMA_VIEW Table_View = 2
	// Populates all fields.
	Table_FULL Table_View = 4
)

var Table_View_name = map[int32]string{
	0: "VIEW_UNSPECIFIED",
	1: "NAME_ONLY",
	2: "SCHEMA_VIEW",
	4: "FULL",
}
var Table_View_value = map[string]int32{
	"VIEW_UNSPECIFIED": 0,
	"NAME_ONLY":        1,
	"SCHEMA_VIEW":      2,
	"FULL":             4,
}

func (x Table_View) String() string {
	return proto.EnumName(Table_View_name, int32(x))
}
func (Table_View) EnumDescriptor() ([]byte, []int) { return fileDescriptorTable, []int{0, 1} }

// A collection of user data indexed by row, column, and timestamp.
// Each table is served using the resources of its parent cluster.
type Table struct {
	// The unique name of the table. Values are of the form
	// projects/<project>/instances/<instance>/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*
	// Views: NAME_ONLY, SCHEMA_VIEW, REPLICATION_VIEW, FULL
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The column families configured for this table, mapped by column family ID.
	// Views: SCHEMA_VIEW, FULL
	// @CreationOnly
	ColumnFamilies map[string]*ColumnFamily `protobuf:"bytes,3,rep,name=column_families,json=columnFamilies" json:"column_families,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// The granularity (e.g. MILLIS, MICROS) at which timestamps are stored in
	// this table. Timestamps not matching the granularity will be rejected.
	// If unspecified at creation time, the value will be set to MILLIS.
	// Views: SCHEMA_VIEW, FULL
	// @CreationOnly
	Granularity Table_TimestampGranularity `protobuf:"varint,4,opt,name=granularity,proto3,enum=google.bigtable.admin.v2.Table_TimestampGranularity" json:"granularity,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptorTable, []int{0} }

func (m *Table) GetColumnFamilies() map[string]*ColumnFamily {
	if m != nil {
		return m.ColumnFamilies
	}
	return nil
}

// A set of columns within a table which share a common configuration.
type ColumnFamily struct {
	// Garbage collection rule specified as a protobuf.
	// Must serialize to at most 500 bytes.
	//
	// NOTE: Garbage collection executes opportunistically in the background, and
	// so it's possible for reads to return a cell even if it matches the active
	// GC expression for its family.
	GcRule *GcRule `protobuf:"bytes,1,opt,name=gc_rule,json=gcRule" json:"gc_rule,omitempty"`
}

func (m *ColumnFamily) Reset()                    { *m = ColumnFamily{} }
func (m *ColumnFamily) String() string            { return proto.CompactTextString(m) }
func (*ColumnFamily) ProtoMessage()               {}
func (*ColumnFamily) Descriptor() ([]byte, []int) { return fileDescriptorTable, []int{1} }

func (m *ColumnFamily) GetGcRule() *GcRule {
	if m != nil {
		return m.GcRule
	}
	return nil
}

// Rule for determining which cells to delete during garbage collection.
type GcRule struct {
	// Types that are valid to be assigned to Rule:
	//	*GcRule_MaxNumVersions
	//	*GcRule_MaxAge
	//	*GcRule_Intersection_
	//	*GcRule_Union_
	Rule isGcRule_Rule `protobuf_oneof:"rule"`
}

func (m *GcRule) Reset()                    { *m = GcRule{} }
func (m *GcRule) String() string            { return proto.CompactTextString(m) }
func (*GcRule) ProtoMessage()               {}
func (*GcRule) Descriptor() ([]byte, []int) { return fileDescriptorTable, []int{2} }

type isGcRule_Rule interface {
	isGcRule_Rule()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GcRule_MaxNumVersions struct {
	MaxNumVersions int32 `protobuf:"varint,1,opt,name=max_num_versions,json=maxNumVersions,proto3,oneof"`
}
type GcRule_MaxAge struct {
	MaxAge *sojern_protobuf.Duration `protobuf:"bytes,2,opt,name=max_age,json=maxAge,oneof"`
}
type GcRule_Intersection_ struct {
	Intersection *GcRule_Intersection `protobuf:"bytes,3,opt,name=intersection,oneof"`
}
type GcRule_Union_ struct {
	Union *GcRule_Union `protobuf:"bytes,4,opt,name=union,oneof"`
}

func (*GcRule_MaxNumVersions) isGcRule_Rule() {}
func (*GcRule_MaxAge) isGcRule_Rule()         {}
func (*GcRule_Intersection_) isGcRule_Rule()  {}
func (*GcRule_Union_) isGcRule_Rule()         {}

func (m *GcRule) GetRule() isGcRule_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *GcRule) GetMaxNumVersions() int32 {
	if x, ok := m.GetRule().(*GcRule_MaxNumVersions); ok {
		return x.MaxNumVersions
	}
	return 0
}

func (m *GcRule) GetMaxAge() *sojern_protobuf.Duration {
	if x, ok := m.GetRule().(*GcRule_MaxAge); ok {
		return x.MaxAge
	}
	return nil
}

func (m *GcRule) GetIntersection() *GcRule_Intersection {
	if x, ok := m.GetRule().(*GcRule_Intersection_); ok {
		return x.Intersection
	}
	return nil
}

func (m *GcRule) GetUnion() *GcRule_Union {
	if x, ok := m.GetRule().(*GcRule_Union_); ok {
		return x.Union
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GcRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GcRule_OneofMarshaler, _GcRule_OneofUnmarshaler, _GcRule_OneofSizer, []interface{}{
		(*GcRule_MaxNumVersions)(nil),
		(*GcRule_MaxAge)(nil),
		(*GcRule_Intersection_)(nil),
		(*GcRule_Union_)(nil),
	}
}

func _GcRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		_ = b.EncodeVarint(1<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MaxAge); err != nil {
			return err
		}
	case *GcRule_Intersection_:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Intersection); err != nil {
			return err
		}
	case *GcRule_Union_:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Union); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GcRule.Rule has unexpected type %T", x)
	}
	return nil
}

func _GcRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GcRule)
	switch tag {
	case 1: // rule.max_num_versions
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Rule = &GcRule_MaxNumVersions{int32(x)}
		return true, err
	case 2: // rule.max_age
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(sojern_protobuf.Duration)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_MaxAge{msg}
		return true, err
	case 3: // rule.intersection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Intersection)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Intersection_{msg}
		return true, err
	case 4: // rule.union
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Union)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Union_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GcRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		s := proto.Size(x.MaxAge)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Intersection_:
		s := proto.Size(x.Intersection)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Union_:
		s := proto.Size(x.Union)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A GcRule which deletes cells matching all of the given rules.
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Intersection) Reset()                    { *m = GcRule_Intersection{} }
func (m *GcRule_Intersection) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Intersection) ProtoMessage()               {}
func (*GcRule_Intersection) Descriptor() ([]byte, []int) { return fileDescriptorTable, []int{2, 0} }

func (m *GcRule_Intersection) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A GcRule which deletes cells matching any of the given rules.
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Union) Reset()                    { *m = GcRule_Union{} }
func (m *GcRule_Union) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Union) ProtoMessage()               {}
func (*GcRule_Union) Descriptor() ([]byte, []int) { return fileDescriptorTable, []int{2, 1} }

func (m *GcRule_Union) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*Table)(nil), "google.bigtable.admin.v2.Table")
	proto.RegisterType((*ColumnFamily)(nil), "google.bigtable.admin.v2.ColumnFamily")
	proto.RegisterType((*GcRule)(nil), "google.bigtable.admin.v2.GcRule")
	proto.RegisterType((*GcRule_Intersection)(nil), "google.bigtable.admin.v2.GcRule.Intersection")
	proto.RegisterType((*GcRule_Union)(nil), "google.bigtable.admin.v2.GcRule.Union")
	proto.RegisterEnum("google.bigtable.admin.v2.Table_TimestampGranularity", Table_TimestampGranularity_name, Table_TimestampGranularity_value)
	proto.RegisterEnum("google.bigtable.admin.v2.Table_View", Table_View_name, Table_View_value)
}
func (m *Table) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Table) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintTable(data, i, uint64(len(m.Name)))
		i += copy(data[i:], m.Name)
	}
	if len(m.ColumnFamilies) > 0 {
		for k, _ := range m.ColumnFamilies {
			data[i] = 0x1a
			i++
			v := m.ColumnFamilies[k]
			if v == nil {
				return 0, errors.New("proto: map has nil element")
			}
			msgSize := v.Size()
			mapSize := 1 + len(k) + sovTable(uint64(len(k))) + 1 + msgSize + sovTable(uint64(msgSize))
			i = encodeVarintTable(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintTable(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintTable(data, i, uint64(v.Size()))
			n1, err := v.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n1
		}
	}
	if m.Granularity != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintTable(data, i, uint64(m.Granularity))
	}
	return i, nil
}

func (m *ColumnFamily) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ColumnFamily) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GcRule != nil {
		data[i] = 0xa
		i++
		i = encodeVarintTable(data, i, uint64(m.GcRule.Size()))
		n2, err := m.GcRule.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *GcRule) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GcRule) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Rule != nil {
		nn3, err := m.Rule.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += nn3
	}
	return i, nil
}

func (m *GcRule_MaxNumVersions) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x8
	i++
	i = encodeVarintTable(data, i, uint64(m.MaxNumVersions))
	return i, nil
}
func (m *GcRule_MaxAge) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.MaxAge != nil {
		data[i] = 0x12
		i++
		i = encodeVarintTable(data, i, uint64(m.MaxAge.Size()))
		n4, err := m.MaxAge.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *GcRule_Intersection_) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.Intersection != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintTable(data, i, uint64(m.Intersection.Size()))
		n5, err := m.Intersection.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *GcRule_Union_) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.Union != nil {
		data[i] = 0x22
		i++
		i = encodeVarintTable(data, i, uint64(m.Union.Size()))
		n6, err := m.Union.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *GcRule_Intersection) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GcRule_Intersection) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, msg := range m.Rules {
			data[i] = 0xa
			i++
			i = encodeVarintTable(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *GcRule_Union) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GcRule_Union) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, msg := range m.Rules {
			data[i] = 0xa
			i++
			i = encodeVarintTable(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Table(data []byte, offset int, v uint64) int {
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
func encodeFixed32Table(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTable(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Table) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTable(uint64(l))
	}
	if len(m.ColumnFamilies) > 0 {
		for k, v := range m.ColumnFamilies {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
			}
			mapEntrySize := 1 + len(k) + sovTable(uint64(len(k))) + 1 + l + sovTable(uint64(l))
			n += mapEntrySize + 1 + sovTable(uint64(mapEntrySize))
		}
	}
	if m.Granularity != 0 {
		n += 1 + sovTable(uint64(m.Granularity))
	}
	return n
}

func (m *ColumnFamily) Size() (n int) {
	var l int
	_ = l
	if m.GcRule != nil {
		l = m.GcRule.Size()
		n += 1 + l + sovTable(uint64(l))
	}
	return n
}

func (m *GcRule) Size() (n int) {
	var l int
	_ = l
	if m.Rule != nil {
		n += m.Rule.Size()
	}
	return n
}

func (m *GcRule_MaxNumVersions) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovTable(uint64(m.MaxNumVersions))
	return n
}
func (m *GcRule_MaxAge) Size() (n int) {
	var l int
	_ = l
	if m.MaxAge != nil {
		l = m.MaxAge.Size()
		n += 1 + l + sovTable(uint64(l))
	}
	return n
}
func (m *GcRule_Intersection_) Size() (n int) {
	var l int
	_ = l
	if m.Intersection != nil {
		l = m.Intersection.Size()
		n += 1 + l + sovTable(uint64(l))
	}
	return n
}
func (m *GcRule_Union_) Size() (n int) {
	var l int
	_ = l
	if m.Union != nil {
		l = m.Union.Size()
		n += 1 + l + sovTable(uint64(l))
	}
	return n
}
func (m *GcRule_Intersection) Size() (n int) {
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovTable(uint64(l))
		}
	}
	return n
}

func (m *GcRule_Union) Size() (n int) {
	var l int
	_ = l
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovTable(uint64(l))
		}
	}
	return n
}

func sovTable(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTable(x uint64) (n int) {
	return sovTable(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Table) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTable
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
			return fmt.Errorf("proto: Table: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Table: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnFamilies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
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
					return ErrIntOverflowTable
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
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthTable
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &ColumnFamily{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.ColumnFamilies == nil {
				m.ColumnFamilies = make(map[string]*ColumnFamily)
			}
			m.ColumnFamilies[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granularity", wireType)
			}
			m.Granularity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Granularity |= (Table_TimestampGranularity(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTable(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTable
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
func (m *ColumnFamily) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTable
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
			return fmt.Errorf("proto: ColumnFamily: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ColumnFamily: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GcRule", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GcRule == nil {
				m.GcRule = &GcRule{}
			}
			if err := m.GcRule.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTable(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTable
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
func (m *GcRule) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTable
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
			return fmt.Errorf("proto: GcRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GcRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNumVersions", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Rule = &GcRule_MaxNumVersions{v}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &sojern_protobuf.Duration{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Rule = &GcRule_MaxAge{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Intersection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &GcRule_Intersection{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Rule = &GcRule_Intersection_{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Union", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &GcRule_Union{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Rule = &GcRule_Union_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTable(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTable
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
func (m *GcRule_Intersection) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTable
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
			return fmt.Errorf("proto: Intersection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Intersection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, &GcRule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTable(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTable
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
func (m *GcRule_Union) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTable
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
			return fmt.Errorf("proto: Union: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Union: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
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
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, &GcRule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTable(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTable
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
func skipTable(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTable
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
					return 0, ErrIntOverflowTable
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
					return 0, ErrIntOverflowTable
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
				return 0, ErrInvalidLengthTable
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTable
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
				next, err := skipTable(data[start:])
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
	ErrInvalidLengthTable = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTable   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorTable = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x93, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xeb, 0xfa, 0xd0, 0xaf, 0xe3, 0x7e, 0xa9, 0xb5, 0xf4, 0x22, 0x44, 0x08, 0x85, 0x48,
	0xa0, 0x0a, 0xa9, 0x8e, 0x94, 0x56, 0xa8, 0x20, 0x04, 0x38, 0xa9, 0x93, 0x5a, 0x72, 0x42, 0xe4,
	0x1c, 0x50, 0x25, 0x24, 0x6b, 0xe3, 0xb8, 0x8b, 0xc1, 0x5e, 0x57, 0x8e, 0x1d, 0xc8, 0x5b, 0xf0,
	0x58, 0xdc, 0xc1, 0x23, 0x20, 0x9e, 0x84, 0x5d, 0x3b, 0x85, 0x14, 0x35, 0x0a, 0xe2, 0x62, 0xe5,
	0xf1, 0xec, 0xfc, 0x7f, 0x3b, 0x33, 0x3b, 0x0b, 0x6a, 0x8a, 0x27, 0xa1, 0xaf, 0x5f, 0x25, 0x71,
	0x1a, 0xa3, 0x32, 0x89, 0x63, 0xc2, 0xfe, 0x26, 0x01, 0x29, 0xdc, 0x78, 0x1a, 0x05, 0x54, 0x9f,
	0x37, 0x2a, 0xaf, 0x48, 0x90, 0xbe, 0xcb, 0x26, 0xba, 0x17, 0x47, 0xf5, 0x59, 0xfc, 0xde, 0x4f,
	0x68, 0x9d, 0x78, 0x61, 0x9c, 0x4d, 0x8f, 0x48, 0x1c, 0x62, 0x4a, 0xea, 0x39, 0x60, 0x92, 0x5d,
	0xd6, 0xa7, 0x59, 0x82, 0xd3, 0x20, 0xa6, 0xbf, 0x8c, 0x82, 0x5d, 0xfb, 0x2a, 0x82, 0x3c, 0xe4,
	0x50, 0x84, 0x40, 0xa2, 0x38, 0xf2, 0xcb, 0x42, 0x55, 0x38, 0xdc, 0x75, 0x72, 0x1b, 0xbd, 0x85,
	0x7d, 0x2f, 0x0e, 0xb3, 0x88, 0xba, 0x97, 0x38, 0x0a, 0xc2, 0xc0, 0x9f, 0x95, 0xc5, 0xaa, 0x78,
	0xa8, 0x36, 0x8e, 0xf5, 0x75, 0x39, 0xe9, 0x39, 0x4d, 0x6f, 0xe5, 0xb2, 0xf6, 0x52, 0x65, 0xd2,
	0x34, 0x59, 0x38, 0x25, 0xef, 0x86, 0x13, 0x8d, 0x41, 0x25, 0x09, 0xa6, 0x59, 0x88, 0x93, 0x20,
	0x5d, 0x94, 0x25, 0x76, 0x70, 0xa9, 0x71, 0xb2, 0x89, 0x3c, 0x0c, 0x22, 0x7f, 0x96, 0xe2, 0xe8,
	0xaa, 0xf3, 0x5b, 0xeb, 0xac, 0x82, 0x2a, 0x01, 0xdc, 0xb9, 0xe5, 0x78, 0xa4, 0x81, 0xf8, 0xc1,
	0x5f, 0x2c, 0xeb, 0xe3, 0x26, 0x7a, 0x0e, 0xf2, 0x1c, 0x87, 0x99, 0x5f, 0xde, 0x66, 0x3e, 0xb5,
	0xf1, 0x68, 0xfd, 0xd1, 0x2b, 0xbc, 0x85, 0x53, 0x88, 0x9e, 0x6d, 0x9f, 0x0a, 0x35, 0x0b, 0x0e,
	0x6e, 0xcb, 0x07, 0x3d, 0x84, 0x07, 0x43, 0xab, 0x6b, 0x0e, 0x86, 0x46, 0xb7, 0xef, 0x76, 0x1c,
	0xa3, 0x37, 0xb2, 0x0d, 0xc7, 0x1a, 0x5e, 0xb8, 0xa3, 0xde, 0xa0, 0x6f, 0xb6, 0xac, 0xb6, 0x65,
	0x9e, 0x69, 0x5b, 0x08, 0x40, 0xe9, 0x5a, 0xb6, 0x6d, 0x0d, 0x34, 0xa1, 0xd6, 0x06, 0x69, 0x1c,
	0xf8, 0x1f, 0xd1, 0x01, 0x68, 0x63, 0xcb, 0x7c, 0xf3, 0x47, 0xe4, 0xff, 0xb0, 0xdb, 0x33, 0xba,
	0xa6, 0xfb, 0xba, 0x67, 0x5f, 0x68, 0x02, 0xda, 0x07, 0x75, 0xd0, 0x3a, 0x37, 0xbb, 0x86, 0xcb,
	0x63, 0xb5, 0x6d, 0xf4, 0x1f, 0x48, 0xed, 0x91, 0x6d, 0x6b, 0x12, 0x4b, 0x69, 0x6f, 0x35, 0x5b,
	0xf4, 0x14, 0x76, 0x88, 0xe7, 0x26, 0x59, 0x58, 0x5c, 0xad, 0xda, 0xa8, 0xae, 0x2f, 0xb3, 0xe3,
	0x39, 0x2c, 0xce, 0x51, 0x48, 0xfe, 0xad, 0x7d, 0x16, 0x41, 0x29, 0x5c, 0xe8, 0x31, 0x68, 0x11,
	0xfe, 0xe4, 0xd2, 0x2c, 0x72, 0xe7, 0x7e, 0x32, 0x63, 0x03, 0x34, 0xcb, 0x71, 0xf2, 0xf9, 0x96,
	0x53, 0x62, 0x3b, 0xbd, 0x2c, 0x1a, 0x2f, 0xfd, 0xe8, 0x04, 0x76, 0x78, 0x2c, 0x26, 0xd7, 0x8d,
	0xbd, 0xab, 0x17, 0xc3, 0xa9, 0x5f, 0x8f, 0xa3, 0x7e, 0xb6, 0x9c, 0x42, 0xa6, 0x56, 0x58, 0xac,
	0x41, 0x7c, 0x34, 0x80, 0xbd, 0x80, 0xa6, 0x8c, 0xe1, 0x7b, 0x7c, 0x87, 0x0d, 0x1a, 0x97, 0x1e,
	0x6d, 0x4a, 0x56, 0xb7, 0x56, 0x44, 0x0c, 0x77, 0x03, 0x82, 0x5e, 0x80, 0x9c, 0x51, 0x4e, 0x93,
	0x36, 0xdd, 0xf0, 0x92, 0x36, 0xa2, 0x05, 0xa6, 0x90, 0x55, 0xda, 0xb0, 0xb7, 0xca, 0x47, 0x4f,
	0x40, 0xe6, 0x9d, 0xe4, 0xb5, 0x8b, 0x7f, 0xd5, 0xca, 0x22, 0xbc, 0xf2, 0x12, 0xe4, 0x9c, 0xfc,
	0xaf, 0x80, 0xa6, 0x02, 0x12, 0x37, 0x9a, 0xa7, 0x5f, 0x7e, 0xdc, 0x17, 0xbe, 0xb1, 0xf5, 0x9d,
	0x2d, 0xb8, 0xc7, 0x1e, 0xfe, 0x5a, 0x42, 0x13, 0xf2, 0x07, 0xd3, 0xe7, 0x3d, 0xef, 0x0b, 0x13,
	0x25, 0x6f, 0xfe, 0xf1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x28, 0x6f, 0x19, 0x5b, 0x04,
	0x00, 0x00,
}
