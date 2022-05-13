// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/report/report.proto

package scammer_report

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's report.
type AddReportRequest struct {
	UserReport           *UserReport `protobuf:"bytes,1,opt,name=user_report,json=userReport,proto3" json:"user_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddReportRequest) Reset()         { *m = AddReportRequest{} }
func (m *AddReportRequest) String() string { return proto.CompactTextString(m) }
func (*AddReportRequest) ProtoMessage()    {}
func (*AddReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1926ef81aa8e283f, []int{0}
}

func (m *AddReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReportRequest.Unmarshal(m, b)
}
func (m *AddReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReportRequest.Marshal(b, m, deterministic)
}
func (m *AddReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReportRequest.Merge(m, src)
}
func (m *AddReportRequest) XXX_Size() int {
	return xxx_messageInfo_AddReportRequest.Size(m)
}
func (m *AddReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddReportRequest proto.InternalMessageInfo

func (m *AddReportRequest) GetUserReport() *UserReport {
	if m != nil {
		return m.UserReport
	}
	return nil
}

// The response message containing the greetings
type AddReportReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReportReply) Reset()         { *m = AddReportReply{} }
func (m *AddReportReply) String() string { return proto.CompactTextString(m) }
func (*AddReportReply) ProtoMessage()    {}
func (*AddReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1926ef81aa8e283f, []int{1}
}

func (m *AddReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReportReply.Unmarshal(m, b)
}
func (m *AddReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReportReply.Marshal(b, m, deterministic)
}
func (m *AddReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReportReply.Merge(m, src)
}
func (m *AddReportReply) XXX_Size() int {
	return xxx_messageInfo_AddReportReply.Size(m)
}
func (m *AddReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddReportReply proto.InternalMessageInfo

func (m *AddReportReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The request message containing the user's report.
type GetReportRequest struct {
	PublicKey            string   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReportRequest) Reset()         { *m = GetReportRequest{} }
func (m *GetReportRequest) String() string { return proto.CompactTextString(m) }
func (*GetReportRequest) ProtoMessage()    {}
func (*GetReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1926ef81aa8e283f, []int{2}
}

func (m *GetReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReportRequest.Unmarshal(m, b)
}
func (m *GetReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReportRequest.Marshal(b, m, deterministic)
}
func (m *GetReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportRequest.Merge(m, src)
}
func (m *GetReportRequest) XXX_Size() int {
	return xxx_messageInfo_GetReportRequest.Size(m)
}
func (m *GetReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportRequest proto.InternalMessageInfo

func (m *GetReportRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

// The response message containing the greetings
type GetReportReply struct {
	UserReport           *UserReport `protobuf:"bytes,1,opt,name=user_report,json=userReport,proto3" json:"user_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetReportReply) Reset()         { *m = GetReportReply{} }
func (m *GetReportReply) String() string { return proto.CompactTextString(m) }
func (*GetReportReply) ProtoMessage()    {}
func (*GetReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1926ef81aa8e283f, []int{3}
}

func (m *GetReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReportReply.Unmarshal(m, b)
}
func (m *GetReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReportReply.Marshal(b, m, deterministic)
}
func (m *GetReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReportReply.Merge(m, src)
}
func (m *GetReportReply) XXX_Size() int {
	return xxx_messageInfo_GetReportReply.Size(m)
}
func (m *GetReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReportReply proto.InternalMessageInfo

func (m *GetReportReply) GetUserReport() *UserReport {
	if m != nil {
		return m.UserReport
	}
	return nil
}

// The request message containing the user's report.
type ListReportRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReportRequest) Reset()         { *m = ListReportRequest{} }
func (m *ListReportRequest) String() string { return proto.CompactTextString(m) }
func (*ListReportRequest) ProtoMessage()    {}
func (*ListReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1926ef81aa8e283f, []int{4}
}

func (m *ListReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReportRequest.Unmarshal(m, b)
}
func (m *ListReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReportRequest.Marshal(b, m, deterministic)
}
func (m *ListReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReportRequest.Merge(m, src)
}
func (m *ListReportRequest) XXX_Size() int {
	return xxx_messageInfo_ListReportRequest.Size(m)
}
func (m *ListReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListReportRequest proto.InternalMessageInfo

// The response message containing the greetings
type ListReportReply struct {
	UserReport           *UserReport `protobuf:"bytes,1,opt,name=user_report,json=userReport,proto3" json:"user_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListReportReply) Reset()         { *m = ListReportReply{} }
func (m *ListReportReply) String() string { return proto.CompactTextString(m) }
func (*ListReportReply) ProtoMessage()    {}
func (*ListReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1926ef81aa8e283f, []int{5}
}

func (m *ListReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReportReply.Unmarshal(m, b)
}
func (m *ListReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReportReply.Marshal(b, m, deterministic)
}
func (m *ListReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReportReply.Merge(m, src)
}
func (m *ListReportReply) XXX_Size() int {
	return xxx_messageInfo_ListReportReply.Size(m)
}
func (m *ListReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListReportReply proto.InternalMessageInfo

func (m *ListReportReply) GetUserReport() *UserReport {
	if m != nil {
		return m.UserReport
	}
	return nil
}

//The user's report
type UserReport struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PublicKey            string               `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	HashTransaction      string               `protobuf:"bytes,3,opt,name=hash_transaction,json=hashTransaction,proto3" json:"hash_transaction,omitempty"`
	Send                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=send,proto3" json:"send,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserReport) Reset()         { *m = UserReport{} }
func (m *UserReport) String() string { return proto.CompactTextString(m) }
func (*UserReport) ProtoMessage()    {}
func (*UserReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_1926ef81aa8e283f, []int{6}
}

func (m *UserReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserReport.Unmarshal(m, b)
}
func (m *UserReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserReport.Marshal(b, m, deterministic)
}
func (m *UserReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserReport.Merge(m, src)
}
func (m *UserReport) XXX_Size() int {
	return xxx_messageInfo_UserReport.Size(m)
}
func (m *UserReport) XXX_DiscardUnknown() {
	xxx_messageInfo_UserReport.DiscardUnknown(m)
}

var xxx_messageInfo_UserReport proto.InternalMessageInfo

func (m *UserReport) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserReport) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *UserReport) GetHashTransaction() string {
	if m != nil {
		return m.HashTransaction
	}
	return ""
}

func (m *UserReport) GetSend() *timestamp.Timestamp {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *UserReport) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*AddReportRequest)(nil), "report.AddReportRequest")
	proto.RegisterType((*AddReportReply)(nil), "report.AddReportReply")
	proto.RegisterType((*GetReportRequest)(nil), "report.GetReportRequest")
	proto.RegisterType((*GetReportReply)(nil), "report.GetReportReply")
	proto.RegisterType((*ListReportRequest)(nil), "report.ListReportRequest")
	proto.RegisterType((*ListReportReply)(nil), "report.ListReportReply")
	proto.RegisterType((*UserReport)(nil), "report.UserReport")
}

func init() { proto.RegisterFile("pkg/report/report.proto", fileDescriptor_1926ef81aa8e283f) }

var fileDescriptor_1926ef81aa8e283f = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x31, 0xa5, 0x54, 0x8c, 0x25, 0xa0, 0x5b, 0xa9, 0xb8, 0x96, 0xaa, 0xa2, 0x55, 0x0f,
	0x6d, 0x25, 0xec, 0x06, 0x1e, 0x20, 0x22, 0x11, 0x41, 0x51, 0x38, 0x59, 0xe4, 0x92, 0x0b, 0xf2,
	0x9f, 0x8d, 0x59, 0x61, 0xe3, 0xcd, 0xee, 0xfa, 0xe0, 0x37, 0xcb, 0x6b, 0xe4, 0x8d, 0x22, 0xd6,
	0x36, 0x36, 0x0e, 0x27, 0x4e, 0xf6, 0x7e, 0xfb, 0xcd, 0x6f, 0x66, 0xfc, 0x19, 0x46, 0x6c, 0x17,
	0xda, 0x9c, 0xb0, 0x84, 0xcb, 0xe2, 0x61, 0x31, 0x9e, 0xc8, 0x04, 0x75, 0xf3, 0x93, 0xf9, 0x2b,
	0x4c, 0x92, 0x30, 0x22, 0xb6, 0x52, 0xbd, 0xf4, 0xd9, 0x96, 0x34, 0x26, 0x42, 0xba, 0x31, 0xcb,
	0x8d, 0x78, 0x09, 0xc3, 0x79, 0x10, 0x38, 0xca, 0xed, 0x90, 0x97, 0x94, 0x08, 0x89, 0x66, 0xa0,
	0xa7, 0x82, 0xf0, 0x4d, 0xce, 0x30, 0xb4, 0xb1, 0xf6, 0x47, 0x9f, 0x22, 0xab, 0x68, 0xf0, 0x28,
	0x08, 0x2f, 0xfc, 0x90, 0x1e, 0xdf, 0xf1, 0x3f, 0xe8, 0xd7, 0x40, 0x2c, 0xca, 0x90, 0x01, 0x5f,
	0x62, 0x22, 0x84, 0x1b, 0x12, 0x85, 0xe8, 0x39, 0xe5, 0x11, 0x5f, 0xc1, 0x70, 0x49, 0xe4, 0x69,
	0xd3, 0x9f, 0x00, 0x2c, 0xf5, 0x22, 0xea, 0x6f, 0x76, 0x24, 0x2b, 0x0a, 0x7a, 0xb9, 0xf2, 0x40,
	0x32, 0xbc, 0x80, 0x7e, 0xad, 0xe4, 0x80, 0xbf, 0x68, 0xca, 0x6f, 0xf0, 0x75, 0x45, 0xc5, 0x69,
	0x6b, 0x7c, 0x07, 0x83, 0xba, 0x78, 0x31, 0xfc, 0x55, 0x03, 0xa8, 0xae, 0x50, 0x1f, 0xda, 0x34,
	0x28, 0x36, 0x69, 0xd3, 0xa0, 0xb1, 0x61, 0xbb, 0xb1, 0x21, 0xfa, 0x0b, 0xc3, 0xad, 0x2b, 0xb6,
	0x1b, 0xc9, 0xdd, 0xbd, 0x70, 0x7d, 0x49, 0x93, 0xbd, 0xf1, 0x49, 0x99, 0x06, 0x07, 0x7d, 0x5d,
	0xc9, 0xc8, 0x82, 0x8e, 0x20, 0xfb, 0xc0, 0xe8, 0xa8, 0xb1, 0x4c, 0x2b, 0x0f, 0xd9, 0x2a, 0x43,
	0xb6, 0xd6, 0x65, 0xc8, 0x8e, 0xf2, 0xa1, 0x31, 0xe8, 0x01, 0x11, 0x3e, 0xa7, 0x4c, 0x51, 0x3f,
	0x2b, 0x6a, 0x5d, 0x9a, 0xbe, 0x69, 0xd0, 0x2d, 0xc6, 0xbe, 0x86, 0xde, 0x31, 0x48, 0x64, 0x94,
	0x2b, 0x37, 0x7f, 0x12, 0xf3, 0xfb, 0x99, 0x1b, 0x16, 0x65, 0xb8, 0x75, 0x00, 0x1c, 0xa3, 0xaa,
	0x00, 0xcd, 0xc0, 0x2b, 0xc0, 0x69, 0xae, 0xb8, 0x85, 0x16, 0xa0, 0x57, 0x79, 0x08, 0xf4, 0xa3,
	0x34, 0x7e, 0x48, 0xce, 0x1c, 0x9d, 0xbb, 0x52, 0x90, 0xff, 0xda, 0xcd, 0xef, 0x27, 0x1c, 0x52,
	0xb9, 0x4d, 0x3d, 0xcb, 0x4f, 0x62, 0xfb, 0x76, 0x75, 0x3f, 0x9f, 0xac, 0x5c, 0xcf, 0x16, 0xbe,
	0x1b, 0xc7, 0x84, 0x4f, 0xf2, 0x4a, 0xaf, 0xab, 0xbe, 0xda, 0xec, 0x3d, 0x00, 0x00, 0xff, 0xff,
	0x55, 0xf7, 0xef, 0x08, 0x4b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReportClient is the client API for Report service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportClient interface {
	// Get Report
	AddReport(ctx context.Context, in *AddReportRequest, opts ...grpc.CallOption) (*AddReportReply, error)
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportReply, error)
	ListReports(ctx context.Context, in *ListReportRequest, opts ...grpc.CallOption) (Report_ListReportsClient, error)
}

type reportClient struct {
	cc *grpc.ClientConn
}

func NewReportClient(cc *grpc.ClientConn) ReportClient {
	return &reportClient{cc}
}

func (c *reportClient) AddReport(ctx context.Context, in *AddReportRequest, opts ...grpc.CallOption) (*AddReportReply, error) {
	out := new(AddReportReply)
	err := c.cc.Invoke(ctx, "/report.Report/AddReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*GetReportReply, error) {
	out := new(GetReportReply)
	err := c.cc.Invoke(ctx, "/report.Report/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportClient) ListReports(ctx context.Context, in *ListReportRequest, opts ...grpc.CallOption) (Report_ListReportsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Report_serviceDesc.Streams[0], "/report.Report/ListReports", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportListReportsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Report_ListReportsClient interface {
	Recv() (*ListReportReply, error)
	grpc.ClientStream
}

type reportListReportsClient struct {
	grpc.ClientStream
}

func (x *reportListReportsClient) Recv() (*ListReportReply, error) {
	m := new(ListReportReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReportServer is the server API for Report service.
type ReportServer interface {
	// Get Report
	AddReport(context.Context, *AddReportRequest) (*AddReportReply, error)
	GetReport(context.Context, *GetReportRequest) (*GetReportReply, error)
	ListReports(*ListReportRequest, Report_ListReportsServer) error
}

// UnimplementedReportServer can be embedded to have forward compatible implementations.
type UnimplementedReportServer struct {
}

func (*UnimplementedReportServer) AddReport(ctx context.Context, req *AddReportRequest) (*AddReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReport not implemented")
}
func (*UnimplementedReportServer) GetReport(ctx context.Context, req *GetReportRequest) (*GetReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReport not implemented")
}
func (*UnimplementedReportServer) ListReports(req *ListReportRequest, srv Report_ListReportsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListReports not implemented")
}

func RegisterReportServer(s *grpc.Server, srv ReportServer) {
	s.RegisterService(&_Report_serviceDesc, srv)
}

func _Report_AddReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).AddReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.Report/AddReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).AddReport(ctx, req.(*AddReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Report_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.Report/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Report_ListReports_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListReportRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReportServer).ListReports(m, &reportListReportsServer{stream})
}

type Report_ListReportsServer interface {
	Send(*ListReportReply) error
	grpc.ServerStream
}

type reportListReportsServer struct {
	grpc.ServerStream
}

func (x *reportListReportsServer) Send(m *ListReportReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Report_serviceDesc = grpc.ServiceDesc{
	ServiceName: "report.Report",
	HandlerType: (*ReportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddReport",
			Handler:    _Report_AddReport_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _Report_GetReport_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListReports",
			Handler:       _Report_ListReports_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/report/report.proto",
}
