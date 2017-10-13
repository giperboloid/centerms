// Code generated by protoc-gen-go. DO NOT EDIT.
// source: centerms.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	centerms.proto

It has these top-level messages:
	DevMeta
	SetDevInitConfigRequest
	SetDevInitConfigResponse
	SaveDevDataRequest
	SaveDevDataResponse
	PatchDevConfigRequest
	PatchDevConfigResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DevMeta struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Mac  string `protobuf:"bytes,3,opt,name=mac" json:"mac,omitempty"`
}

func (m *DevMeta) Reset()                    { *m = DevMeta{} }
func (m *DevMeta) String() string            { return proto.CompactTextString(m) }
func (*DevMeta) ProtoMessage()               {}
func (*DevMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DevMeta) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DevMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DevMeta) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type SetDevInitConfigRequest struct {
	Time int64    `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Meta *DevMeta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
}

func (m *SetDevInitConfigRequest) Reset()                    { *m = SetDevInitConfigRequest{} }
func (m *SetDevInitConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*SetDevInitConfigRequest) ProtoMessage()               {}
func (*SetDevInitConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SetDevInitConfigRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SetDevInitConfigRequest) GetMeta() *DevMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type SetDevInitConfigResponse struct {
	Config []byte `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *SetDevInitConfigResponse) Reset()                    { *m = SetDevInitConfigResponse{} }
func (m *SetDevInitConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*SetDevInitConfigResponse) ProtoMessage()               {}
func (*SetDevInitConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetDevInitConfigResponse) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

type SaveDevDataRequest struct {
	Time int64    `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Meta *DevMeta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Data []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SaveDevDataRequest) Reset()                    { *m = SaveDevDataRequest{} }
func (m *SaveDevDataRequest) String() string            { return proto.CompactTextString(m) }
func (*SaveDevDataRequest) ProtoMessage()               {}
func (*SaveDevDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SaveDevDataRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SaveDevDataRequest) GetMeta() *DevMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SaveDevDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type SaveDevDataResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *SaveDevDataResponse) Reset()                    { *m = SaveDevDataResponse{} }
func (m *SaveDevDataResponse) String() string            { return proto.CompactTextString(m) }
func (*SaveDevDataResponse) ProtoMessage()               {}
func (*SaveDevDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SaveDevDataResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type PatchDevConfigRequest struct {
	Config []byte `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *PatchDevConfigRequest) Reset()                    { *m = PatchDevConfigRequest{} }
func (m *PatchDevConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*PatchDevConfigRequest) ProtoMessage()               {}
func (*PatchDevConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PatchDevConfigRequest) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

type PatchDevConfigResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *PatchDevConfigResponse) Reset()                    { *m = PatchDevConfigResponse{} }
func (m *PatchDevConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*PatchDevConfigResponse) ProtoMessage()               {}
func (*PatchDevConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PatchDevConfigResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*DevMeta)(nil), "pb.DevMeta")
	proto.RegisterType((*SetDevInitConfigRequest)(nil), "pb.SetDevInitConfigRequest")
	proto.RegisterType((*SetDevInitConfigResponse)(nil), "pb.SetDevInitConfigResponse")
	proto.RegisterType((*SaveDevDataRequest)(nil), "pb.SaveDevDataRequest")
	proto.RegisterType((*SaveDevDataResponse)(nil), "pb.SaveDevDataResponse")
	proto.RegisterType((*PatchDevConfigRequest)(nil), "pb.PatchDevConfigRequest")
	proto.RegisterType((*PatchDevConfigResponse)(nil), "pb.PatchDevConfigResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CenterService service

type CenterServiceClient interface {
	SetDevInitConfig(ctx context.Context, in *SetDevInitConfigRequest, opts ...grpc.CallOption) (*SetDevInitConfigResponse, error)
	SaveDevData(ctx context.Context, in *SaveDevDataRequest, opts ...grpc.CallOption) (*SaveDevDataResponse, error)
}

type centerServiceClient struct {
	cc *grpc.ClientConn
}

func NewCenterServiceClient(cc *grpc.ClientConn) CenterServiceClient {
	return &centerServiceClient{cc}
}

func (c *centerServiceClient) SetDevInitConfig(ctx context.Context, in *SetDevInitConfigRequest, opts ...grpc.CallOption) (*SetDevInitConfigResponse, error) {
	out := new(SetDevInitConfigResponse)
	err := grpc.Invoke(ctx, "/pb.CenterService/SetDevInitConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centerServiceClient) SaveDevData(ctx context.Context, in *SaveDevDataRequest, opts ...grpc.CallOption) (*SaveDevDataResponse, error) {
	out := new(SaveDevDataResponse)
	err := grpc.Invoke(ctx, "/pb.CenterService/SaveDevData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CenterService service

type CenterServiceServer interface {
	SetDevInitConfig(context.Context, *SetDevInitConfigRequest) (*SetDevInitConfigResponse, error)
	SaveDevData(context.Context, *SaveDevDataRequest) (*SaveDevDataResponse, error)
}

func RegisterCenterServiceServer(s *grpc.Server, srv CenterServiceServer) {
	s.RegisterService(&_CenterService_serviceDesc, srv)
}

func _CenterService_SetDevInitConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDevInitConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).SetDevInitConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CenterService/SetDevInitConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).SetDevInitConfig(ctx, req.(*SetDevInitConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CenterService_SaveDevData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveDevDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CenterServiceServer).SaveDevData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CenterService/SaveDevData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CenterServiceServer).SaveDevData(ctx, req.(*SaveDevDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CenterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CenterService",
	HandlerType: (*CenterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDevInitConfig",
			Handler:    _CenterService_SetDevInitConfig_Handler,
		},
		{
			MethodName: "SaveDevData",
			Handler:    _CenterService_SaveDevData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "centerms.proto",
}

// Client API for DevService service

type DevServiceClient interface {
	PatchDevConfig(ctx context.Context, in *PatchDevConfigRequest, opts ...grpc.CallOption) (*PatchDevConfigResponse, error)
}

type devServiceClient struct {
	cc *grpc.ClientConn
}

func NewDevServiceClient(cc *grpc.ClientConn) DevServiceClient {
	return &devServiceClient{cc}
}

func (c *devServiceClient) PatchDevConfig(ctx context.Context, in *PatchDevConfigRequest, opts ...grpc.CallOption) (*PatchDevConfigResponse, error) {
	out := new(PatchDevConfigResponse)
	err := grpc.Invoke(ctx, "/pb.DevService/PatchDevConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DevService service

type DevServiceServer interface {
	PatchDevConfig(context.Context, *PatchDevConfigRequest) (*PatchDevConfigResponse, error)
}

func RegisterDevServiceServer(s *grpc.Server, srv DevServiceServer) {
	s.RegisterService(&_DevService_serviceDesc, srv)
}

func _DevService_PatchDevConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchDevConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevServiceServer).PatchDevConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DevService/PatchDevConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevServiceServer).PatchDevConfig(ctx, req.(*PatchDevConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DevService",
	HandlerType: (*DevServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PatchDevConfig",
			Handler:    _DevService_PatchDevConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "centerms.proto",
}

func init() { proto.RegisterFile("centerms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x3b, 0x4f, 0xf3, 0x40,
	0x10, 0xfc, 0xf2, 0x50, 0x3e, 0xb1, 0x09, 0x51, 0xb4, 0x88, 0xc4, 0x18, 0x24, 0xd0, 0x55, 0x34,
	0x04, 0x64, 0xfe, 0x00, 0x52, 0xdc, 0xa4, 0xe0, 0xa1, 0x4b, 0x41, 0x45, 0x71, 0x31, 0x0b, 0xb8,
	0xf0, 0x83, 0xdc, 0xe6, 0x24, 0xfe, 0x0e, 0xbf, 0x14, 0xdd, 0xc6, 0x11, 0xe4, 0x45, 0x43, 0x37,
	0x37, 0xeb, 0x9d, 0xdd, 0x99, 0x35, 0x74, 0x13, 0xca, 0x99, 0x66, 0x99, 0x1d, 0x96, 0xb3, 0x82,
	0x0b, 0xac, 0x97, 0x53, 0x35, 0x82, 0xff, 0x31, 0xb9, 0x5b, 0x62, 0x83, 0x08, 0x4d, 0xfe, 0x28,
	0x29, 0xa8, 0x9d, 0xd5, 0xce, 0xf7, 0xb4, 0x60, 0xcf, 0xe5, 0x26, 0xa3, 0xa0, 0xbe, 0xe0, 0x3c,
	0xc6, 0x1e, 0x34, 0x32, 0x93, 0x04, 0x0d, 0xa1, 0x3c, 0x54, 0x77, 0x30, 0x98, 0x10, 0xc7, 0xe4,
	0xc6, 0x79, 0xca, 0xa3, 0x22, 0x7f, 0x49, 0x5f, 0x35, 0xbd, 0xcf, 0xc9, 0xb2, 0x88, 0xa6, 0xd9,
	0x42, 0xb4, 0xa1, 0x05, 0xe3, 0x29, 0x34, 0x33, 0x62, 0x23, 0xa2, 0xed, 0xa8, 0x3d, 0x2c, 0xa7,
	0xc3, 0x6a, 0x07, 0x2d, 0x05, 0x15, 0x41, 0xb0, 0xa9, 0x67, 0xcb, 0x22, 0xb7, 0x84, 0x7d, 0x68,
	0x25, 0xc2, 0x88, 0x64, 0x47, 0x57, 0x2f, 0xf5, 0x04, 0x38, 0x31, 0x8e, 0x62, 0x72, 0xb1, 0x61,
	0xf3, 0x97, 0xf1, 0xbe, 0xe9, 0xd9, 0xb0, 0x11, 0x87, 0x1d, 0x2d, 0x58, 0x5d, 0xc0, 0xc1, 0x8a,
	0xfc, 0xf7, 0x36, 0x96, 0x0d, 0xcf, 0x6d, 0x95, 0x5a, 0xf5, 0x52, 0x97, 0x70, 0xf8, 0x60, 0x38,
	0x79, 0x8b, 0xc9, 0xad, 0xe6, 0xb1, 0x6b, 0xfd, 0x2b, 0xe8, 0xaf, 0x37, 0xfc, 0x3e, 0x22, 0xfa,
	0xac, 0xc1, 0xfe, 0x48, 0x0e, 0x3a, 0xa1, 0x99, 0x4b, 0x13, 0xc2, 0x7b, 0xe8, 0xad, 0xc7, 0x86,
	0xc7, 0xde, 0xde, 0x8e, 0xe3, 0x84, 0x27, 0xdb, 0x8b, 0x8b, 0xc1, 0xea, 0x1f, 0xde, 0x40, 0xfb,
	0x87, 0x69, 0xec, 0xcb, 0xe7, 0x1b, 0x21, 0x87, 0x83, 0x0d, 0x7e, 0xa9, 0x10, 0x3d, 0x02, 0xc4,
	0xe4, 0x96, 0x0b, 0x8e, 0xa1, 0xbb, 0x6a, 0x12, 0x8f, 0x7c, 0xeb, 0xd6, 0xa4, 0xc2, 0x70, 0x5b,
	0x69, 0x29, 0x3c, 0x6d, 0xc9, 0x2f, 0x7c, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xa6, 0xdb,
	0x5c, 0xd4, 0x02, 0x00, 0x00,
}
