// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: pkg/protos/restaurant.proto

package restaurant

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_restaurant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_restaurant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pkg_protos_restaurant_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Online   string  `protobuf:"bytes,3,opt,name=Online,proto3" json:"Online,omitempty"`
	ItemLine []*Item `protobuf:"bytes,4,rep,name=ItemLine,proto3" json:"ItemLine,omitempty"`
	Rating   float32 `protobuf:"fixed32,5,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Category string  `protobuf:"bytes,6,opt,name=Category,proto3" json:"Category,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_restaurant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_restaurant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_pkg_protos_restaurant_proto_rawDescGZIP(), []int{1}
}

func (x *Restaurant) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetOnline() string {
	if x != nil {
		return x.Online
	}
	return ""
}

func (x *Restaurant) GetItemLine() []*Item {
	if x != nil {
		return x.ItemLine
	}
	return nil
}

func (x *Restaurant) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Restaurant) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_restaurant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_restaurant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_pkg_protos_restaurant_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type NoParamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParamRequest) Reset() {
	*x = NoParamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_restaurant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParamRequest) ProtoMessage() {}

func (x *NoParamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_restaurant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParamRequest.ProtoReflect.Descriptor instead.
func (*NoParamRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protos_restaurant_proto_rawDescGZIP(), []int{3}
}

var File_pkg_protos_restaurant_proto protoreflect.FileDescriptor

var file_pkg_protos_restaurant_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x04, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x0a,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x19, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xe0, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x73,
	0x12, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x4e, 0x6f,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x6b, 0x67, 0x2f,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_protos_restaurant_proto_rawDescOnce sync.Once
	file_pkg_protos_restaurant_proto_rawDescData = file_pkg_protos_restaurant_proto_rawDesc
)

func file_pkg_protos_restaurant_proto_rawDescGZIP() []byte {
	file_pkg_protos_restaurant_proto_rawDescOnce.Do(func() {
		file_pkg_protos_restaurant_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protos_restaurant_proto_rawDescData)
	})
	return file_pkg_protos_restaurant_proto_rawDescData
}

var file_pkg_protos_restaurant_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_protos_restaurant_proto_goTypes = []interface{}{
	(*Item)(nil),           // 0: restaurant.Item
	(*Restaurant)(nil),     // 1: restaurant.Restaurant
	(*Request)(nil),        // 2: restaurant.Request
	(*NoParamRequest)(nil), // 3: restaurant.NoParamRequest
}
var file_pkg_protos_restaurant_proto_depIdxs = []int32{
	0, // 0: restaurant.Restaurant.ItemLine:type_name -> restaurant.Item
	1, // 1: restaurant.RestaurantService.AddRestaurant:input_type -> restaurant.Restaurant
	3, // 2: restaurant.RestaurantService.GetRestaurants:input_type -> restaurant.NoParamRequest
	2, // 3: restaurant.RestaurantService.GetRestaurant:input_type -> restaurant.Request
	1, // 4: restaurant.RestaurantService.AddRestaurant:output_type -> restaurant.Restaurant
	1, // 5: restaurant.RestaurantService.GetRestaurants:output_type -> restaurant.Restaurant
	1, // 6: restaurant.RestaurantService.GetRestaurant:output_type -> restaurant.Restaurant
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_protos_restaurant_proto_init() }
func file_pkg_protos_restaurant_proto_init() {
	if File_pkg_protos_restaurant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protos_restaurant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_pkg_protos_restaurant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurant); i {
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
		file_pkg_protos_restaurant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_pkg_protos_restaurant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParamRequest); i {
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
			RawDescriptor: file_pkg_protos_restaurant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protos_restaurant_proto_goTypes,
		DependencyIndexes: file_pkg_protos_restaurant_proto_depIdxs,
		MessageInfos:      file_pkg_protos_restaurant_proto_msgTypes,
	}.Build()
	File_pkg_protos_restaurant_proto = out.File
	file_pkg_protos_restaurant_proto_rawDesc = nil
	file_pkg_protos_restaurant_proto_goTypes = nil
	file_pkg_protos_restaurant_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	AddRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*Restaurant, error)
	GetRestaurants(ctx context.Context, in *NoParamRequest, opts ...grpc.CallOption) (RestaurantService_GetRestaurantsClient, error)
	GetRestaurant(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Restaurant, error)
}

type restaurantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantServiceClient(cc grpc.ClientConnInterface) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) AddRestaurant(ctx context.Context, in *Restaurant, opts ...grpc.CallOption) (*Restaurant, error) {
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantService/AddRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetRestaurants(ctx context.Context, in *NoParamRequest, opts ...grpc.CallOption) (RestaurantService_GetRestaurantsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RestaurantService_serviceDesc.Streams[0], "/restaurant.RestaurantService/GetRestaurants", opts...)
	if err != nil {
		return nil, err
	}
	x := &restaurantServiceGetRestaurantsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RestaurantService_GetRestaurantsClient interface {
	Recv() (*Restaurant, error)
	grpc.ClientStream
}

type restaurantServiceGetRestaurantsClient struct {
	grpc.ClientStream
}

func (x *restaurantServiceGetRestaurantsClient) Recv() (*Restaurant, error) {
	m := new(Restaurant)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *restaurantServiceClient) GetRestaurant(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Restaurant, error) {
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, "/restaurant.RestaurantService/GetRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantServiceServer is the server API for RestaurantService service.
type RestaurantServiceServer interface {
	AddRestaurant(context.Context, *Restaurant) (*Restaurant, error)
	GetRestaurants(*NoParamRequest, RestaurantService_GetRestaurantsServer) error
	GetRestaurant(context.Context, *Request) (*Restaurant, error)
}

// UnimplementedRestaurantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRestaurantServiceServer struct {
}

func (*UnimplementedRestaurantServiceServer) AddRestaurant(context.Context, *Restaurant) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRestaurant not implemented")
}
func (*UnimplementedRestaurantServiceServer) GetRestaurants(*NoParamRequest, RestaurantService_GetRestaurantsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRestaurants not implemented")
}
func (*UnimplementedRestaurantServiceServer) GetRestaurant(context.Context, *Request) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurant not implemented")
}

func RegisterRestaurantServiceServer(s *grpc.Server, srv RestaurantServiceServer) {
	s.RegisterService(&_RestaurantService_serviceDesc, srv)
}

func _RestaurantService_AddRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Restaurant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).AddRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantService/AddRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).AddRestaurant(ctx, req.(*Restaurant))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetRestaurants_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RestaurantServiceServer).GetRestaurants(m, &restaurantServiceGetRestaurantsServer{stream})
}

type RestaurantService_GetRestaurantsServer interface {
	Send(*Restaurant) error
	grpc.ServerStream
}

type restaurantServiceGetRestaurantsServer struct {
	grpc.ServerStream
}

func (x *restaurantServiceGetRestaurantsServer) Send(m *Restaurant) error {
	return x.ServerStream.SendMsg(m)
}

func _RestaurantService_GetRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.RestaurantService/GetRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetRestaurant(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestaurantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant.RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRestaurant",
			Handler:    _RestaurantService_AddRestaurant_Handler,
		},
		{
			MethodName: "GetRestaurant",
			Handler:    _RestaurantService_GetRestaurant_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRestaurants",
			Handler:       _RestaurantService_GetRestaurants_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/protos/restaurant.proto",
}
