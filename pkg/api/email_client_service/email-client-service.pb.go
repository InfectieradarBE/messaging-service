// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.11.4
// source: api/email-client-service.proto

package email_client_service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ServiceStatus_StatusValue int32

const (
	ServiceStatus_NORMAL  ServiceStatus_StatusValue = 0
	ServiceStatus_PROBLEM ServiceStatus_StatusValue = 1
)

// Enum value maps for ServiceStatus_StatusValue.
var (
	ServiceStatus_StatusValue_name = map[int32]string{
		0: "NORMAL",
		1: "PROBLEM",
	}
	ServiceStatus_StatusValue_value = map[string]int32{
		"NORMAL":  0,
		"PROBLEM": 1,
	}
)

func (x ServiceStatus_StatusValue) Enum() *ServiceStatus_StatusValue {
	p := new(ServiceStatus_StatusValue)
	*p = x
	return p
}

func (x ServiceStatus_StatusValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceStatus_StatusValue) Descriptor() protoreflect.EnumDescriptor {
	return file_api_email_client_service_proto_enumTypes[0].Descriptor()
}

func (ServiceStatus_StatusValue) Type() protoreflect.EnumType {
	return &file_api_email_client_service_proto_enumTypes[0]
}

func (x ServiceStatus_StatusValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceStatus_StatusValue.Descriptor instead.
func (ServiceStatus_StatusValue) EnumDescriptor() ([]byte, []int) {
	return file_api_email_client_service_proto_rawDescGZIP(), []int{0, 0}
}

//
// Status is typically used as a return value indicating if the method was
// performed normally, or the system has any internal error e.g. checking
// system status of a service
type ServiceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  ServiceStatus_StatusValue `protobuf:"varint,1,opt,name=status,proto3,enum=inf.email_client_service.ServiceStatus_StatusValue" json:"status,omitempty"`
	Msg     string                    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Version string                    `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ServiceStatus) Reset() {
	*x = ServiceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_email_client_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStatus) ProtoMessage() {}

func (x *ServiceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_email_client_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStatus.ProtoReflect.Descriptor instead.
func (*ServiceStatus) Descriptor() ([]byte, []int) {
	return file_api_email_client_service_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceStatus) GetStatus() ServiceStatus_StatusValue {
	if x != nil {
		return x.Status
	}
	return ServiceStatus_NORMAL
}

func (x *ServiceStatus) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ServiceStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type SendEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To      []string `protobuf:"bytes,1,rep,name=to,proto3" json:"to,omitempty"`
	Content string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendEmailReq) Reset() {
	*x = SendEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_email_client_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailReq) ProtoMessage() {}

func (x *SendEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_email_client_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailReq.ProtoReflect.Descriptor instead.
func (*SendEmailReq) Descriptor() ([]byte, []int) {
	return file_api_email_client_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailReq) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendEmailReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_api_email_client_service_proto protoreflect.FileDescriptor

var file_api_email_client_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x69, 0x6e, 0x66, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x69, 0x6e, 0x66, 0x2e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x01, 0x22, 0x38, 0x0a, 0x0c, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x32, 0xc0, 0x01, 0x0a, 0x15, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x12, 0x49,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5c, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x27,
	0x2e, 0x69, 0x6e, 0x66, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x7a, 0x61, 0x6e,
	0x65, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_email_client_service_proto_rawDescOnce sync.Once
	file_api_email_client_service_proto_rawDescData = file_api_email_client_service_proto_rawDesc
)

func file_api_email_client_service_proto_rawDescGZIP() []byte {
	file_api_email_client_service_proto_rawDescOnce.Do(func() {
		file_api_email_client_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_email_client_service_proto_rawDescData)
	})
	return file_api_email_client_service_proto_rawDescData
}

var file_api_email_client_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_email_client_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_email_client_service_proto_goTypes = []interface{}{
	(ServiceStatus_StatusValue)(0), // 0: inf.email_client_service.ServiceStatus.StatusValue
	(*ServiceStatus)(nil),          // 1: inf.email_client_service.ServiceStatus
	(*SendEmailReq)(nil),           // 2: inf.email_client_service.SendEmailReq
	(*empty.Empty)(nil),            // 3: google.protobuf.Empty
}
var file_api_email_client_service_proto_depIdxs = []int32{
	0, // 0: inf.email_client_service.ServiceStatus.status:type_name -> inf.email_client_service.ServiceStatus.StatusValue
	3, // 1: inf.email_client_service.EmailClientServiceApi.Status:input_type -> google.protobuf.Empty
	2, // 2: inf.email_client_service.EmailClientServiceApi.SendEmail:input_type -> inf.email_client_service.SendEmailReq
	1, // 3: inf.email_client_service.EmailClientServiceApi.Status:output_type -> inf.email_client_service.ServiceStatus
	1, // 4: inf.email_client_service.EmailClientServiceApi.SendEmail:output_type -> inf.email_client_service.ServiceStatus
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_email_client_service_proto_init() }
func file_api_email_client_service_proto_init() {
	if File_api_email_client_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_email_client_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStatus); i {
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
		file_api_email_client_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailReq); i {
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
			RawDescriptor: file_api_email_client_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_email_client_service_proto_goTypes,
		DependencyIndexes: file_api_email_client_service_proto_depIdxs,
		EnumInfos:         file_api_email_client_service_proto_enumTypes,
		MessageInfos:      file_api_email_client_service_proto_msgTypes,
	}.Build()
	File_api_email_client_service_proto = out.File
	file_api_email_client_service_proto_rawDesc = nil
	file_api_email_client_service_proto_goTypes = nil
	file_api_email_client_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmailClientServiceApiClient is the client API for EmailClientServiceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailClientServiceApiClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServiceStatus, error)
	SendEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*ServiceStatus, error)
}

type emailClientServiceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailClientServiceApiClient(cc grpc.ClientConnInterface) EmailClientServiceApiClient {
	return &emailClientServiceApiClient{cc}
}

func (c *emailClientServiceApiClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/inf.email_client_service.EmailClientServiceApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailClientServiceApiClient) SendEmail(ctx context.Context, in *SendEmailReq, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/inf.email_client_service.EmailClientServiceApi/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailClientServiceApiServer is the server API for EmailClientServiceApi service.
type EmailClientServiceApiServer interface {
	Status(context.Context, *empty.Empty) (*ServiceStatus, error)
	SendEmail(context.Context, *SendEmailReq) (*ServiceStatus, error)
}

// UnimplementedEmailClientServiceApiServer can be embedded to have forward compatible implementations.
type UnimplementedEmailClientServiceApiServer struct {
}

func (*UnimplementedEmailClientServiceApiServer) Status(context.Context, *empty.Empty) (*ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedEmailClientServiceApiServer) SendEmail(context.Context, *SendEmailReq) (*ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}

func RegisterEmailClientServiceApiServer(s *grpc.Server, srv EmailClientServiceApiServer) {
	s.RegisterService(&_EmailClientServiceApi_serviceDesc, srv)
}

func _EmailClientServiceApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailClientServiceApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.email_client_service.EmailClientServiceApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailClientServiceApiServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailClientServiceApi_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailClientServiceApiServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.email_client_service.EmailClientServiceApi/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailClientServiceApiServer).SendEmail(ctx, req.(*SendEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailClientServiceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inf.email_client_service.EmailClientServiceApi",
	HandlerType: (*EmailClientServiceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _EmailClientServiceApi_Status_Handler,
		},
		{
			MethodName: "SendEmail",
			Handler:    _EmailClientServiceApi_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/email-client-service.proto",
}