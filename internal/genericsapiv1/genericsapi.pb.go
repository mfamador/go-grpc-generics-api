// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: protos/genericsapi.proto

package genericsapiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SortDirection int32

const (
	SortDirection_ASCENDING  SortDirection = 0
	SortDirection_DESCENDING SortDirection = 1
)

// Enum value maps for SortDirection.
var (
	SortDirection_name = map[int32]string{
		0: "ASCENDING",
		1: "DESCENDING",
	}
	SortDirection_value = map[string]int32{
		"ASCENDING":  0,
		"DESCENDING": 1,
	}
)

func (x SortDirection) Enum() *SortDirection {
	p := new(SortDirection)
	*p = x
	return p
}

func (x SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_genericsapi_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_protos_genericsapi_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{0}
}

type FilterOperator int32

const (
	FilterOperator_EQUAL              FilterOperator = 0
	FilterOperator_NOTEQUAL           FilterOperator = 1
	FilterOperator_CONTAINS           FilterOperator = 2
	FilterOperator_STARTSWITH         FilterOperator = 3
	FilterOperator_ENDSWITH           FilterOperator = 4
	FilterOperator_GREATERTHAN        FilterOperator = 5
	FilterOperator_GREATERTHANOREQUAL FilterOperator = 6
	FilterOperator_LESSTHAN           FilterOperator = 7
	FilterOperator_LESSTHANOREQUAL    FilterOperator = 8
)

// Enum value maps for FilterOperator.
var (
	FilterOperator_name = map[int32]string{
		0: "EQUAL",
		1: "NOTEQUAL",
		2: "CONTAINS",
		3: "STARTSWITH",
		4: "ENDSWITH",
		5: "GREATERTHAN",
		6: "GREATERTHANOREQUAL",
		7: "LESSTHAN",
		8: "LESSTHANOREQUAL",
	}
	FilterOperator_value = map[string]int32{
		"EQUAL":              0,
		"NOTEQUAL":           1,
		"CONTAINS":           2,
		"STARTSWITH":         3,
		"ENDSWITH":           4,
		"GREATERTHAN":        5,
		"GREATERTHANOREQUAL": 6,
		"LESSTHAN":           7,
		"LESSTHANOREQUAL":    8,
	}
)

func (x FilterOperator) Enum() *FilterOperator {
	p := new(FilterOperator)
	*p = x
	return p
}

func (x FilterOperator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FilterOperator) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_genericsapi_proto_enumTypes[1].Descriptor()
}

func (FilterOperator) Type() protoreflect.EnumType {
	return &file_protos_genericsapi_proto_enumTypes[1]
}

func (x FilterOperator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FilterOperator.Descriptor instead.
func (FilterOperator) EnumDescriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{1}
}

type StatusReply_ServingStatus int32

const (
	StatusReply_UNKNOWN     StatusReply_ServingStatus = 0
	StatusReply_SERVING     StatusReply_ServingStatus = 1
	StatusReply_NOT_SERVING StatusReply_ServingStatus = 2
)

// Enum value maps for StatusReply_ServingStatus.
var (
	StatusReply_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
	}
	StatusReply_ServingStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"SERVING":     1,
		"NOT_SERVING": 2,
	}
)

func (x StatusReply_ServingStatus) Enum() *StatusReply_ServingStatus {
	p := new(StatusReply_ServingStatus)
	*p = x
	return p
}

func (x StatusReply_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusReply_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_genericsapi_proto_enumTypes[2].Descriptor()
}

func (StatusReply_ServingStatus) Type() protoreflect.EnumType {
	return &file_protos_genericsapi_proto_enumTypes[2]
}

func (x StatusReply_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusReply_ServingStatus.Descriptor instead.
func (StatusReply_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{7, 0}
}

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value       float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	FooSpecific string                 `protobuf:"bytes,4,opt,name=foo_specific,json=fooSpecific,proto3" json:"foo_specific,omitempty"`
	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{0}
}

func (x *Foo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Foo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Foo) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Foo) GetFooSpecific() string {
	if x != nil {
		return x.FooSpecific
	}
	return ""
}

func (x *Foo) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value              float64                `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	BarSpecific        string                 `protobuf:"bytes,4,opt,name=bar_specific,json=barSpecific,proto3" json:"bar_specific,omitempty"`
	BarAnotherSpecific string                 `protobuf:"bytes,5,opt,name=bar_another_specific,json=barAnotherSpecific,proto3" json:"bar_another_specific,omitempty"`
	Timestamp          *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{1}
}

func (x *Bar) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bar) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Bar) GetBarSpecific() string {
	if x != nil {
		return x.BarSpecific
	}
	return ""
}

func (x *Bar) GetBarAnotherSpecific() string {
	if x != nil {
		return x.BarAnotherSpecific
	}
	return ""
}

func (x *Bar) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Sorting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName    string        `protobuf:"bytes,1,opt,name=column_name,json=columnName,proto3" json:"column_name,omitempty"`
	SortDirection SortDirection `protobuf:"varint,2,opt,name=sort_direction,json=sortDirection,proto3,enum=genericsapiv1.SortDirection" json:"sort_direction,omitempty"`
}

func (x *Sorting) Reset() {
	*x = Sorting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sorting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sorting) ProtoMessage() {}

func (x *Sorting) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sorting.ProtoReflect.Descriptor instead.
func (*Sorting) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{2}
}

func (x *Sorting) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *Sorting) GetSortDirection() SortDirection {
	if x != nil {
		return x.SortDirection
	}
	return SortDirection_ASCENDING
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName  string         `protobuf:"bytes,1,opt,name=column_name,json=columnName,proto3" json:"column_name,omitempty"`
	Operator    FilterOperator `protobuf:"varint,2,opt,name=operator,proto3,enum=genericsapiv1.FilterOperator" json:"operator,omitempty"`
	StringValue string         `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{3}
}

func (x *Filter) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *Filter) GetOperator() FilterOperator {
	if x != nil {
		return x.Operator
	}
	return FilterOperator_EQUAL
}

func (x *Filter) GetStringValue() string {
	if x != nil {
		return x.StringValue
	}
	return ""
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sorting []*Sorting `protobuf:"bytes,1,rep,name=sorting,proto3" json:"sorting,omitempty"`
	Filters []*Filter  `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	Limit   uint32     `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Cursor  string     `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{4}
}

func (x *ReadRequest) GetSorting() []*Sorting {
	if x != nil {
		return x.Sorting
	}
	return nil
}

func (x *ReadRequest) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ReadRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ReadRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type ReadFooReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foos   []*Foo `protobuf:"bytes,1,rep,name=foos,proto3" json:"foos,omitempty"`
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ReadFooReply) Reset() {
	*x = ReadFooReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFooReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFooReply) ProtoMessage() {}

func (x *ReadFooReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFooReply.ProtoReflect.Descriptor instead.
func (*ReadFooReply) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{5}
}

func (x *ReadFooReply) GetFoos() []*Foo {
	if x != nil {
		return x.Foos
	}
	return nil
}

func (x *ReadFooReply) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type ReadBarReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bars   []*Bar `protobuf:"bytes,1,rep,name=bars,proto3" json:"bars,omitempty"`
	Cursor string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *ReadBarReply) Reset() {
	*x = ReadBarReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBarReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBarReply) ProtoMessage() {}

func (x *ReadBarReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBarReply.ProtoReflect.Descriptor instead.
func (*ReadBarReply) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{6}
}

func (x *ReadBarReply) GetBars() []*Bar {
	if x != nil {
		return x.Bars
	}
	return nil
}

func (x *ReadBarReply) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusReply_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=genericsapiv1.StatusReply_ServingStatus" json:"status,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{7}
}

func (x *StatusReply) GetStatus() StatusReply_ServingStatus {
	if x != nil {
		return x.Status
	}
	return StatusReply_UNKNOWN
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_genericsapi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_genericsapi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_protos_genericsapi_proto_rawDescGZIP(), []int{8}
}

var File_protos_genericsapi_proto protoreflect.FileDescriptor

var file_protos_genericsapi_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x73, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x03, 0x46,
	0x6f, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x6f, 0x6f, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x6f, 0x6f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xce, 0x01, 0x0a, 0x03, 0x42, 0x61,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62,
	0x61, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x62, 0x61, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x12, 0x30,
	0x0a, 0x14, 0x62, 0x61, 0x72, 0x5f, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x62, 0x61,
	0x72, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6f, 0x0a, 0x07, 0x53, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x6f,
	0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x87, 0x01, 0x0a, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07,
	0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x4e, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x46, 0x6f,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x4e, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x42, 0x61,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x62, 0x61, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x04, 0x62, 0x61, 0x72, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x2e, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x2a, 0xa1, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x51, 0x55, 0x41,
	0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x52, 0x54, 0x53, 0x57, 0x49, 0x54, 0x48, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x45, 0x4e, 0x44, 0x53, 0x57, 0x49, 0x54, 0x48, 0x10, 0x04, 0x12, 0x0f, 0x0a,
	0x0b, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x05, 0x12, 0x16,
	0x0a, 0x12, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x54, 0x48, 0x41, 0x4e, 0x4f, 0x52, 0x45,
	0x51, 0x55, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x45, 0x53, 0x53, 0x54, 0x48,
	0x41, 0x4e, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x45, 0x53, 0x53, 0x54, 0x48, 0x41, 0x4e,
	0x4f, 0x52, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x08, 0x32, 0x4d, 0x0a, 0x0a, 0x46, 0x6f, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x46, 0x6f, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x4d, 0x0a, 0x0a, 0x42, 0x61, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1a,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42,
	0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x53, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x1e, 0x5a, 0x1c,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x3b, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_genericsapi_proto_rawDescOnce sync.Once
	file_protos_genericsapi_proto_rawDescData = file_protos_genericsapi_proto_rawDesc
)

func file_protos_genericsapi_proto_rawDescGZIP() []byte {
	file_protos_genericsapi_proto_rawDescOnce.Do(func() {
		file_protos_genericsapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_genericsapi_proto_rawDescData)
	})
	return file_protos_genericsapi_proto_rawDescData
}

var file_protos_genericsapi_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protos_genericsapi_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_genericsapi_proto_goTypes = []interface{}{
	(SortDirection)(0),             // 0: genericsapiv1.SortDirection
	(FilterOperator)(0),            // 1: genericsapiv1.FilterOperator
	(StatusReply_ServingStatus)(0), // 2: genericsapiv1.StatusReply.ServingStatus
	(*Foo)(nil),                    // 3: genericsapiv1.Foo
	(*Bar)(nil),                    // 4: genericsapiv1.Bar
	(*Sorting)(nil),                // 5: genericsapiv1.Sorting
	(*Filter)(nil),                 // 6: genericsapiv1.Filter
	(*ReadRequest)(nil),            // 7: genericsapiv1.ReadRequest
	(*ReadFooReply)(nil),           // 8: genericsapiv1.ReadFooReply
	(*ReadBarReply)(nil),           // 9: genericsapiv1.ReadBarReply
	(*StatusReply)(nil),            // 10: genericsapiv1.StatusReply
	(*StatusRequest)(nil),          // 11: genericsapiv1.StatusRequest
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
}
var file_protos_genericsapi_proto_depIdxs = []int32{
	12, // 0: genericsapiv1.Foo.timestamp:type_name -> google.protobuf.Timestamp
	12, // 1: genericsapiv1.Bar.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 2: genericsapiv1.Sorting.sort_direction:type_name -> genericsapiv1.SortDirection
	1,  // 3: genericsapiv1.Filter.operator:type_name -> genericsapiv1.FilterOperator
	5,  // 4: genericsapiv1.ReadRequest.sorting:type_name -> genericsapiv1.Sorting
	6,  // 5: genericsapiv1.ReadRequest.filters:type_name -> genericsapiv1.Filter
	3,  // 6: genericsapiv1.ReadFooReply.foos:type_name -> genericsapiv1.Foo
	4,  // 7: genericsapiv1.ReadBarReply.bars:type_name -> genericsapiv1.Bar
	2,  // 8: genericsapiv1.StatusReply.status:type_name -> genericsapiv1.StatusReply.ServingStatus
	7,  // 9: genericsapiv1.FooService.Read:input_type -> genericsapiv1.ReadRequest
	7,  // 10: genericsapiv1.BarService.Read:input_type -> genericsapiv1.ReadRequest
	11, // 11: genericsapiv1.StatusService.Status:input_type -> genericsapiv1.StatusRequest
	8,  // 12: genericsapiv1.FooService.Read:output_type -> genericsapiv1.ReadFooReply
	9,  // 13: genericsapiv1.BarService.Read:output_type -> genericsapiv1.ReadBarReply
	10, // 14: genericsapiv1.StatusService.Status:output_type -> genericsapiv1.StatusReply
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_protos_genericsapi_proto_init() }
func file_protos_genericsapi_proto_init() {
	if File_protos_genericsapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_genericsapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_genericsapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_genericsapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sorting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_genericsapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_genericsapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_genericsapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadFooReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_genericsapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBarReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_genericsapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_genericsapi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_genericsapi_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_protos_genericsapi_proto_goTypes,
		DependencyIndexes: file_protos_genericsapi_proto_depIdxs,
		EnumInfos:         file_protos_genericsapi_proto_enumTypes,
		MessageInfos:      file_protos_genericsapi_proto_msgTypes,
	}.Build()
	File_protos_genericsapi_proto = out.File
	file_protos_genericsapi_proto_rawDesc = nil
	file_protos_genericsapi_proto_goTypes = nil
	file_protos_genericsapi_proto_depIdxs = nil
}
