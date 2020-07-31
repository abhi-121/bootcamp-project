// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: pkg/protos/customer.proto

package customer

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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Addr  string `protobuf:"bytes,4,opt,name=Addr,proto3" json:"Addr,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protos_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_pkg_protos_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Customer) GetAddr() string {
	if x != nil {
		return x.Addr
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
		mi := &file_pkg_protos_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_customer_proto_msgTypes[1]
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
	return file_pkg_protos_customer_proto_rawDescGZIP(), []int{1}
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
		mi := &file_pkg_protos_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParamRequest) ProtoMessage() {}

func (x *NoParamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protos_customer_proto_msgTypes[2]
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
	return file_pkg_protos_customer_proto_rawDescGZIP(), []int{2}
}

var File_pkg_protos_customer_proto protoreflect.FileDescriptor

var file_pkg_protos_customer_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x58, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41,
	0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x22,
	0x19, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xc4, 0x01, 0x0a,
	0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x12, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protos_customer_proto_rawDescOnce sync.Once
	file_pkg_protos_customer_proto_rawDescData = file_pkg_protos_customer_proto_rawDesc
)

func file_pkg_protos_customer_proto_rawDescGZIP() []byte {
	file_pkg_protos_customer_proto_rawDescOnce.Do(func() {
		file_pkg_protos_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protos_customer_proto_rawDescData)
	})
	return file_pkg_protos_customer_proto_rawDescData
}

var file_pkg_protos_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_protos_customer_proto_goTypes = []interface{}{
	(*Customer)(nil),       // 0: customer.Customer
	(*Request)(nil),        // 1: customer.Request
	(*NoParamRequest)(nil), // 2: customer.NoParamRequest
}
var file_pkg_protos_customer_proto_depIdxs = []int32{
	0, // 0: customer.CustomerService.AddCustomer:input_type -> customer.Customer
	2, // 1: customer.CustomerService.GetCustomers:input_type -> customer.NoParamRequest
	1, // 2: customer.CustomerService.GetCustomer:input_type -> customer.Request
	0, // 3: customer.CustomerService.AddCustomer:output_type -> customer.Customer
	0, // 4: customer.CustomerService.GetCustomers:output_type -> customer.Customer
	0, // 5: customer.CustomerService.GetCustomer:output_type -> customer.Customer
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_protos_customer_proto_init() }
func file_pkg_protos_customer_proto_init() {
	if File_pkg_protos_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protos_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_pkg_protos_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protos_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pkg_protos_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protos_customer_proto_goTypes,
		DependencyIndexes: file_pkg_protos_customer_proto_depIdxs,
		MessageInfos:      file_pkg_protos_customer_proto_msgTypes,
	}.Build()
	File_pkg_protos_customer_proto = out.File
	file_pkg_protos_customer_proto_rawDesc = nil
	file_pkg_protos_customer_proto_goTypes = nil
	file_pkg_protos_customer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerServiceClient interface {
	AddCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error)
	GetCustomers(ctx context.Context, in *NoParamRequest, opts ...grpc.CallOption) (CustomerService_GetCustomersClient, error)
	GetCustomer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Customer, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) AddCustomer(ctx context.Context, in *Customer, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/AddCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomers(ctx context.Context, in *NoParamRequest, opts ...grpc.CallOption) (CustomerService_GetCustomersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CustomerService_serviceDesc.Streams[0], "/customer.CustomerService/GetCustomers", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerServiceGetCustomersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CustomerService_GetCustomersClient interface {
	Recv() (*Customer, error)
	grpc.ClientStream
}

type customerServiceGetCustomersClient struct {
	grpc.ClientStream
}

func (x *customerServiceGetCustomersClient) Recv() (*Customer, error) {
	m := new(Customer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Customer, error) {
	out := new(Customer)
	err := c.cc.Invoke(ctx, "/customer.CustomerService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
type CustomerServiceServer interface {
	AddCustomer(context.Context, *Customer) (*Customer, error)
	GetCustomers(*NoParamRequest, CustomerService_GetCustomersServer) error
	GetCustomer(context.Context, *Request) (*Customer, error)
}

// UnimplementedCustomerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (*UnimplementedCustomerServiceServer) AddCustomer(context.Context, *Customer) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomer not implemented")
}
func (*UnimplementedCustomerServiceServer) GetCustomers(*NoParamRequest, CustomerService_GetCustomersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCustomers not implemented")
}
func (*UnimplementedCustomerServiceServer) GetCustomer(context.Context, *Request) (*Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}

func RegisterCustomerServiceServer(s *grpc.Server, srv CustomerServiceServer) {
	s.RegisterService(&_CustomerService_serviceDesc, srv)
}

func _CustomerService_AddCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Customer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).AddCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/AddCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).AddCustomer(ctx, req.(*Customer))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetCustomers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServiceServer).GetCustomers(m, &customerServiceGetCustomersServer{stream})
}

type CustomerService_GetCustomersServer interface {
	Send(*Customer) error
	grpc.ServerStream
}

type customerServiceGetCustomersServer struct {
	grpc.ServerStream
}

func (x *customerServiceGetCustomersServer) Send(m *Customer) error {
	return x.ServerStream.SendMsg(m)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.CustomerService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCustomer",
			Handler:    _CustomerService_AddCustomer_Handler,
		},
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCustomers",
			Handler:       _CustomerService_GetCustomers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/protos/customer.proto",
}
