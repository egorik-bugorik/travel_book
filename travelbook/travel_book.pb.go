// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: travel_book.proto

package travelbook

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Long int32 `protobuf:"varint,1,opt,name=long,proto3" json:"long,omitempty"`
	Lat  int32 `protobuf:"varint,2,opt,name=lat,proto3" json:"lat,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_travel_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_travel_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_travel_book_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLong() int32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *Point) GetLat() int32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location *Point `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_travel_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_travel_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_travel_book_proto_rawDescGZIP(), []int{1}
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

type Rectangle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lo *Point `protobuf:"bytes,1,opt,name=lo,proto3" json:"lo,omitempty"`
	Hi *Point `protobuf:"bytes,2,opt,name=hi,proto3" json:"hi,omitempty"`
}

func (x *Rectangle) Reset() {
	*x = Rectangle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_travel_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rectangle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rectangle) ProtoMessage() {}

func (x *Rectangle) ProtoReflect() protoreflect.Message {
	mi := &file_travel_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rectangle.ProtoReflect.Descriptor instead.
func (*Rectangle) Descriptor() ([]byte, []int) {
	return file_travel_book_proto_rawDescGZIP(), []int{2}
}

func (x *Rectangle) GetLo() *Point {
	if x != nil {
		return x.Lo
	}
	return nil
}

func (x *Rectangle) GetHi() *Point {
	if x != nil {
		return x.Hi
	}
	return nil
}

type RouteNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Point `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RouteNote) Reset() {
	*x = RouteNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_travel_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteNote) ProtoMessage() {}

func (x *RouteNote) ProtoReflect() protoreflect.Message {
	mi := &file_travel_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteNote.ProtoReflect.Descriptor instead.
func (*RouteNote) Descriptor() ([]byte, []int) {
	return file_travel_book_proto_rawDescGZIP(), []int{3}
}

func (x *RouteNote) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *RouteNote) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RouteSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointCount   int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount,proto3" json:"point_count,omitempty"`
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount,proto3" json:"feature_count,omitempty"`
	ElapsedTime  int32 `protobuf:"varint,3,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	Distance     int32 `protobuf:"varint,4,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *RouteSummary) Reset() {
	*x = RouteSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_travel_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteSummary) ProtoMessage() {}

func (x *RouteSummary) ProtoReflect() protoreflect.Message {
	mi := &file_travel_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteSummary.ProtoReflect.Descriptor instead.
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return file_travel_book_proto_rawDescGZIP(), []int{4}
}

func (x *RouteSummary) GetPointCount() int32 {
	if x != nil {
		return x.PointCount
	}
	return 0
}

func (x *RouteSummary) GetFeatureCount() int32 {
	if x != nil {
		return x.FeatureCount
	}
	return 0
}

func (x *RouteSummary) GetElapsedTime() int32 {
	if x != nil {
		return x.ElapsedTime
	}
	return 0
}

func (x *RouteSummary) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

var File_travel_book_proto protoreflect.FileDescriptor

var file_travel_book_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x22,
	0x2d, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x22, 0x4c,
	0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x09,
	0x52, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x02, 0x6c, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x6c, 0x6f, 0x12, 0x21, 0x0a, 0x02,
	0x68, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x68, 0x69, 0x22,
	0x54, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x84, 0x02, 0x0a, 0x0a,
	0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x36, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x72,
	0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x52,
	0x65, 0x63, 0x74, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x1a, 0x13, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3e, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x3f, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x15,
	0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x17, 0x5a, 0x15, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b,
	0x2f, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_travel_book_proto_rawDescOnce sync.Once
	file_travel_book_proto_rawDescData = file_travel_book_proto_rawDesc
)

func file_travel_book_proto_rawDescGZIP() []byte {
	file_travel_book_proto_rawDescOnce.Do(func() {
		file_travel_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_travel_book_proto_rawDescData)
	})
	return file_travel_book_proto_rawDescData
}

var file_travel_book_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_travel_book_proto_goTypes = []interface{}{
	(*Point)(nil),        // 0: travelbook.Point
	(*Feature)(nil),      // 1: travelbook.Feature
	(*Rectangle)(nil),    // 2: travelbook.Rectangle
	(*RouteNote)(nil),    // 3: travelbook.RouteNote
	(*RouteSummary)(nil), // 4: travelbook.RouteSummary
}
var file_travel_book_proto_depIdxs = []int32{
	0, // 0: travelbook.Feature.location:type_name -> travelbook.Point
	0, // 1: travelbook.Rectangle.lo:type_name -> travelbook.Point
	0, // 2: travelbook.Rectangle.hi:type_name -> travelbook.Point
	0, // 3: travelbook.RouteNote.location:type_name -> travelbook.Point
	0, // 4: travelbook.TravelBook.GetFeature:input_type -> travelbook.Point
	2, // 5: travelbook.TravelBook.ListFeature:input_type -> travelbook.Rectangle
	0, // 6: travelbook.TravelBook.RecordRoute:input_type -> travelbook.Point
	3, // 7: travelbook.TravelBook.RouteChat:input_type -> travelbook.RouteNote
	1, // 8: travelbook.TravelBook.GetFeature:output_type -> travelbook.Feature
	1, // 9: travelbook.TravelBook.ListFeature:output_type -> travelbook.Feature
	4, // 10: travelbook.TravelBook.RecordRoute:output_type -> travelbook.RouteSummary
	3, // 11: travelbook.TravelBook.RouteChat:output_type -> travelbook.RouteNote
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_travel_book_proto_init() }
func file_travel_book_proto_init() {
	if File_travel_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_travel_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_travel_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_travel_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rectangle); i {
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
		file_travel_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteNote); i {
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
		file_travel_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteSummary); i {
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
			RawDescriptor: file_travel_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_travel_book_proto_goTypes,
		DependencyIndexes: file_travel_book_proto_depIdxs,
		MessageInfos:      file_travel_book_proto_msgTypes,
	}.Build()
	File_travel_book_proto = out.File
	file_travel_book_proto_rawDesc = nil
	file_travel_book_proto_goTypes = nil
	file_travel_book_proto_depIdxs = nil
}