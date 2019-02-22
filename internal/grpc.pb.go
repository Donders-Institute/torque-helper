// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// JobInfoRequest defines required input for retrieving job-related information.
type JobInfoRequest struct {
	Jid                  string   `protobuf:"bytes,1,opt,name=jid,proto3" json:"jid,omitempty"`
	Xml                  bool     `protobuf:"varint,2,opt,name=xml,proto3" json:"xml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobInfoRequest) Reset()         { *m = JobInfoRequest{} }
func (m *JobInfoRequest) String() string { return proto.CompactTextString(m) }
func (*JobInfoRequest) ProtoMessage()    {}
func (*JobInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *JobInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobInfoRequest.Unmarshal(m, b)
}
func (m *JobInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobInfoRequest.Marshal(b, m, deterministic)
}
func (m *JobInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobInfoRequest.Merge(m, src)
}
func (m *JobInfoRequest) XXX_Size() int {
	return xxx_messageInfo_JobInfoRequest.Size(m)
}
func (m *JobInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobInfoRequest proto.InternalMessageInfo

func (m *JobInfoRequest) GetJid() string {
	if m != nil {
		return m.Jid
	}
	return ""
}

func (m *JobInfoRequest) GetXml() bool {
	if m != nil {
		return m.Xml
	}
	return false
}

// UserInfoRequest defines required input for retrieving user-related information.
type UserInfoRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Xml                  bool     `protobuf:"varint,2,opt,name=xml,proto3" json:"xml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRequest) Reset()         { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()    {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *UserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRequest.Unmarshal(m, b)
}
func (m *UserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRequest.Marshal(b, m, deterministic)
}
func (m *UserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRequest.Merge(m, src)
}
func (m *UserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoRequest.Size(m)
}
func (m *UserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRequest proto.InternalMessageInfo

func (m *UserInfoRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserInfoRequest) GetXml() bool {
	if m != nil {
		return m.Xml
	}
	return false
}

// GeneralResponse is a very simple and naive output message.
type GeneralResponse struct {
	ResponseData         string   `protobuf:"bytes,1,opt,name=responseData,proto3" json:"responseData,omitempty"`
	ExitCode             int32    `protobuf:"varint,2,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralResponse) Reset()         { *m = GeneralResponse{} }
func (m *GeneralResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralResponse) ProtoMessage()    {}
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{2}
}

func (m *GeneralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralResponse.Unmarshal(m, b)
}
func (m *GeneralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralResponse.Marshal(b, m, deterministic)
}
func (m *GeneralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralResponse.Merge(m, src)
}
func (m *GeneralResponse) XXX_Size() int {
	return xxx_messageInfo_GeneralResponse.Size(m)
}
func (m *GeneralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralResponse proto.InternalMessageInfo

func (m *GeneralResponse) GetResponseData() string {
	if m != nil {
		return m.ResponseData
	}
	return ""
}

func (m *GeneralResponse) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *GeneralResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*JobInfoRequest)(nil), "grpc.JobInfoRequest")
	proto.RegisterType((*UserInfoRequest)(nil), "grpc.UserInfoRequest")
	proto.RegisterType((*GeneralResponse)(nil), "grpc.GeneralResponse")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x5f, 0x6b, 0xe2, 0x40,
	0x14, 0xc5, 0x71, 0x5d, 0xc5, 0xbd, 0xc8, 0xba, 0x3b, 0xb8, 0x22, 0xee, 0x8b, 0xf8, 0xe4, 0x53,
	0x04, 0xf7, 0x8f, 0x2c, 0x4b, 0xa1, 0x54, 0x8b, 0x6d, 0x20, 0x94, 0x46, 0xfb, 0x01, 0x66, 0xe2,
	0x4d, 0x48, 0x9b, 0xe4, 0xc6, 0x99, 0x89, 0xd8, 0x2f, 0xd8, 0xcf, 0x55, 0x26, 0xd1, 0xb6, 0x4a,
	0x7d, 0x48, 0xdf, 0x4e, 0x0e, 0xf3, 0xbb, 0x39, 0x9c, 0x03, 0x10, 0xc8, 0xd4, 0xb3, 0x52, 0x49,
	0x9a, 0xd8, 0x67, 0xa3, 0x7b, 0x3f, 0x03, 0xa2, 0x20, 0xc2, 0x51, 0xee, 0x89, 0xcc, 0x1f, 0x61,
	0x9c, 0xea, 0xc7, 0xe2, 0xc9, 0xe0, 0x37, 0x7c, 0xb5, 0x49, 0x5c, 0x27, 0x3e, 0xb9, 0xb8, 0xce,
	0x50, 0x69, 0xf6, 0x0d, 0xaa, 0xf7, 0xe1, 0xaa, 0x5b, 0xe9, 0x57, 0x86, 0x5f, 0x5c, 0x23, 0x8d,
	0xb3, 0x8d, 0xa3, 0xee, 0xa7, 0x7e, 0x65, 0xd8, 0x70, 0x8d, 0x1c, 0xfc, 0x81, 0xd6, 0x9d, 0x42,
	0x79, 0x84, 0x65, 0xaf, 0x58, 0xf6, 0x2e, 0x96, 0x41, 0x6b, 0x8e, 0x09, 0x4a, 0x1e, 0xb9, 0xa8,
	0x52, 0x4a, 0x14, 0xb2, 0x01, 0x34, 0xe5, 0x4e, 0xcf, 0xb8, 0xe6, 0x3b, 0xfe, 0xc0, 0x63, 0x3d,
	0x68, 0xe0, 0x36, 0xd4, 0x53, 0x5a, 0x61, 0x7e, 0xad, 0xe6, 0xbe, 0x7c, 0x1b, 0x1e, 0xa5, 0x24,
	0xe9, 0xa0, 0x52, 0x3c, 0xc0, 0x6e, 0xb5, 0xe0, 0xdf, 0x7a, 0xe3, 0xa7, 0x2a, 0x74, 0x96, 0x24,
	0xd7, 0x19, 0x5e, 0x61, 0x94, 0xa2, 0x5c, 0xc8, 0xcd, 0x02, 0xe5, 0x26, 0xf4, 0x90, 0x4d, 0xa0,
	0xb1, 0x94, 0xdc, 0x43, 0x9b, 0x04, 0x6b, 0x5b, 0x79, 0x75, 0x87, 0x75, 0xf4, 0x7e, 0x14, 0xee,
	0x71, 0xee, 0x33, 0x68, 0x16, 0x27, 0xa7, 0x94, 0xf8, 0x61, 0xc0, 0x3a, 0x56, 0xd1, 0xb2, 0xb5,
	0x6f, 0xd9, 0xba, 0x34, 0x2d, 0x9f, 0xc2, 0xff, 0x03, 0x38, 0xc4, 0xc5, 0xc7, 0xe0, 0x73, 0xf8,
	0x3e, 0x47, 0x6d, 0x93, 0xb8, 0x88, 0xc8, 0x7b, 0x70, 0x91, 0x2b, 0x4a, 0xca, 0xa5, 0x9f, 0x41,
	0x7b, 0x8e, 0x3a, 0xc7, 0x71, 0x65, 0x93, 0x50, 0x37, 0xbe, 0xd9, 0x93, 0xed, 0x9e, 0x1f, 0x6d,
	0x7b, 0xea, 0xca, 0x5f, 0xa8, 0xdd, 0x2a, 0xcd, 0x75, 0xd9, 0xfc, 0x13, 0xa8, 0xe7, 0xdc, 0xb6,
	0x24, 0x38, 0x5e, 0x1c, 0xee, 0xe8, 0x50, 0xbc, 0xdf, 0xf1, 0x1f, 0x80, 0x4d, 0xc2, 0xc1, 0xd8,
	0xc4, 0x2e, 0xd5, 0x85, 0xa8, 0xe7, 0xff, 0xfe, 0xf5, 0x1c, 0x00, 0x00, 0xff, 0xff, 0xba, 0x41,
	0x63, 0x35, 0x39, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TorqueHelperSrvServiceClient is the client API for TorqueHelperSrvService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TorqueHelperSrvServiceClient interface {
	TraceJob(ctx context.Context, in *JobInfoRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	TorqueConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GeneralResponse, error)
	MoabConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GeneralResponse, error)
	GetJobBlockReason(ctx context.Context, in *JobInfoRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	GetBlockedJobsOfUser(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	Qstat(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GeneralResponse, error)
	Qstatx(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type torqueHelperSrvServiceClient struct {
	cc *grpc.ClientConn
}

func NewTorqueHelperSrvServiceClient(cc *grpc.ClientConn) TorqueHelperSrvServiceClient {
	return &torqueHelperSrvServiceClient{cc}
}

func (c *torqueHelperSrvServiceClient) TraceJob(ctx context.Context, in *JobInfoRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/grpc.TorqueHelperSrvService/TraceJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torqueHelperSrvServiceClient) TorqueConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/grpc.TorqueHelperSrvService/TorqueConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torqueHelperSrvServiceClient) MoabConfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/grpc.TorqueHelperSrvService/MoabConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torqueHelperSrvServiceClient) GetJobBlockReason(ctx context.Context, in *JobInfoRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/grpc.TorqueHelperSrvService/GetJobBlockReason", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torqueHelperSrvServiceClient) GetBlockedJobsOfUser(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/grpc.TorqueHelperSrvService/GetBlockedJobsOfUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torqueHelperSrvServiceClient) Qstat(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/grpc.TorqueHelperSrvService/Qstat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torqueHelperSrvServiceClient) Qstatx(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/grpc.TorqueHelperSrvService/Qstatx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TorqueHelperSrvServiceServer is the server API for TorqueHelperSrvService service.
type TorqueHelperSrvServiceServer interface {
	TraceJob(context.Context, *JobInfoRequest) (*GeneralResponse, error)
	TorqueConfig(context.Context, *empty.Empty) (*GeneralResponse, error)
	MoabConfig(context.Context, *empty.Empty) (*GeneralResponse, error)
	GetJobBlockReason(context.Context, *JobInfoRequest) (*GeneralResponse, error)
	GetBlockedJobsOfUser(context.Context, *UserInfoRequest) (*GeneralResponse, error)
	Qstat(context.Context, *empty.Empty) (*GeneralResponse, error)
	Qstatx(context.Context, *empty.Empty) (*GeneralResponse, error)
}

func RegisterTorqueHelperSrvServiceServer(s *grpc.Server, srv TorqueHelperSrvServiceServer) {
	s.RegisterService(&_TorqueHelperSrvService_serviceDesc, srv)
}

func _TorqueHelperSrvService_TraceJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqueHelperSrvServiceServer).TraceJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TorqueHelperSrvService/TraceJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqueHelperSrvServiceServer).TraceJob(ctx, req.(*JobInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TorqueHelperSrvService_TorqueConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqueHelperSrvServiceServer).TorqueConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TorqueHelperSrvService/TorqueConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqueHelperSrvServiceServer).TorqueConfig(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TorqueHelperSrvService_MoabConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqueHelperSrvServiceServer).MoabConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TorqueHelperSrvService/MoabConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqueHelperSrvServiceServer).MoabConfig(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TorqueHelperSrvService_GetJobBlockReason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqueHelperSrvServiceServer).GetJobBlockReason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TorqueHelperSrvService/GetJobBlockReason",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqueHelperSrvServiceServer).GetJobBlockReason(ctx, req.(*JobInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TorqueHelperSrvService_GetBlockedJobsOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqueHelperSrvServiceServer).GetBlockedJobsOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TorqueHelperSrvService/GetBlockedJobsOfUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqueHelperSrvServiceServer).GetBlockedJobsOfUser(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TorqueHelperSrvService_Qstat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqueHelperSrvServiceServer).Qstat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TorqueHelperSrvService/Qstat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqueHelperSrvServiceServer).Qstat(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TorqueHelperSrvService_Qstatx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqueHelperSrvServiceServer).Qstatx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TorqueHelperSrvService/Qstatx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqueHelperSrvServiceServer).Qstatx(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TorqueHelperSrvService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.TorqueHelperSrvService",
	HandlerType: (*TorqueHelperSrvServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TraceJob",
			Handler:    _TorqueHelperSrvService_TraceJob_Handler,
		},
		{
			MethodName: "TorqueConfig",
			Handler:    _TorqueHelperSrvService_TorqueConfig_Handler,
		},
		{
			MethodName: "MoabConfig",
			Handler:    _TorqueHelperSrvService_MoabConfig_Handler,
		},
		{
			MethodName: "GetJobBlockReason",
			Handler:    _TorqueHelperSrvService_GetJobBlockReason_Handler,
		},
		{
			MethodName: "GetBlockedJobsOfUser",
			Handler:    _TorqueHelperSrvService_GetBlockedJobsOfUser_Handler,
		},
		{
			MethodName: "Qstat",
			Handler:    _TorqueHelperSrvService_Qstat_Handler,
		},
		{
			MethodName: "Qstatx",
			Handler:    _TorqueHelperSrvService_Qstatx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

// TorqueHelperMomServiceClient is the client API for TorqueHelperMomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TorqueHelperMomServiceClient interface {
	JobMemInfo(ctx context.Context, in *JobInfoRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
}

type torqueHelperMomServiceClient struct {
	cc *grpc.ClientConn
}

func NewTorqueHelperMomServiceClient(cc *grpc.ClientConn) TorqueHelperMomServiceClient {
	return &torqueHelperMomServiceClient{cc}
}

func (c *torqueHelperMomServiceClient) JobMemInfo(ctx context.Context, in *JobInfoRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/grpc.TorqueHelperMomService/JobMemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TorqueHelperMomServiceServer is the server API for TorqueHelperMomService service.
type TorqueHelperMomServiceServer interface {
	JobMemInfo(context.Context, *JobInfoRequest) (*GeneralResponse, error)
}

func RegisterTorqueHelperMomServiceServer(s *grpc.Server, srv TorqueHelperMomServiceServer) {
	s.RegisterService(&_TorqueHelperMomService_serviceDesc, srv)
}

func _TorqueHelperMomService_JobMemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqueHelperMomServiceServer).JobMemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.TorqueHelperMomService/JobMemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqueHelperMomServiceServer).JobMemInfo(ctx, req.(*JobInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TorqueHelperMomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.TorqueHelperMomService",
	HandlerType: (*TorqueHelperMomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JobMemInfo",
			Handler:    _TorqueHelperMomService_JobMemInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
