// Code generated by protoc-gen-go. DO NOT EDIT.
// source: centerms.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	centerms.proto

It has these top-level messages:
	DevMeta
	SaveDevDataRequest
	SaveDevDataResponse
	DevMetaResponse
	DevDataResponse
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

type SaveDevDataRequest struct {
	Action string   `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Time   int64    `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Meta   *DevMeta `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
	Data   []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SaveDevDataRequest) Reset()                    { *m = SaveDevDataRequest{} }
func (m *SaveDevDataRequest) String() string            { return proto.CompactTextString(m) }
func (*SaveDevDataRequest) ProtoMessage()               {}
func (*SaveDevDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SaveDevDataRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

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
func (*SaveDevDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SaveDevDataResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DevMetaResponse struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Mac  string `protobuf:"bytes,3,opt,name=mac" json:"mac,omitempty"`
}

func (m *DevMetaResponse) Reset()                    { *m = DevMetaResponse{} }
func (m *DevMetaResponse) String() string            { return proto.CompactTextString(m) }
func (*DevMetaResponse) ProtoMessage()               {}
func (*DevMetaResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DevMetaResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DevMetaResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DevMetaResponse) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

type DevDataResponse struct {
	Site string            `protobuf:"bytes,1,opt,name=site" json:"site,omitempty"`
	Meta *DevMetaResponse  `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Data map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DevDataResponse) Reset()                    { *m = DevDataResponse{} }
func (m *DevDataResponse) String() string            { return proto.CompactTextString(m) }
func (*DevDataResponse) ProtoMessage()               {}
func (*DevDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DevDataResponse) GetSite() string {
	if m != nil {
		return m.Site
	}
	return ""
}

func (m *DevDataResponse) GetMeta() *DevMetaResponse {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DevDataResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*DevMeta)(nil), "pb.DevMeta")
	proto.RegisterType((*SaveDevDataRequest)(nil), "pb.SaveDevDataRequest")
	proto.RegisterType((*SaveDevDataResponse)(nil), "pb.SaveDevDataResponse")
	proto.RegisterType((*DevMetaResponse)(nil), "pb.DevMetaResponse")
	proto.RegisterType((*DevDataResponse)(nil), "pb.DevDataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DevService service

type DevServiceClient interface {
	SaveDevData(ctx context.Context, in *SaveDevDataRequest, opts ...grpc.CallOption) (*SaveDevDataResponse, error)
}

type devServiceClient struct {
	cc *grpc.ClientConn
}

func NewDevServiceClient(cc *grpc.ClientConn) DevServiceClient {
	return &devServiceClient{cc}
}

func (c *devServiceClient) SaveDevData(ctx context.Context, in *SaveDevDataRequest, opts ...grpc.CallOption) (*SaveDevDataResponse, error) {
	out := new(SaveDevDataResponse)
	err := grpc.Invoke(ctx, "/pb.DevService/SaveDevData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DevService service

type DevServiceServer interface {
	SaveDevData(context.Context, *SaveDevDataRequest) (*SaveDevDataResponse, error)
}

func RegisterDevServiceServer(s *grpc.Server, srv DevServiceServer) {
	s.RegisterService(&_DevService_serviceDesc, srv)
}

func _DevService_SaveDevData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveDevDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevServiceServer).SaveDevData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DevService/SaveDevData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevServiceServer).SaveDevData(ctx, req.(*SaveDevDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DevService",
	HandlerType: (*DevServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveDevData",
			Handler:    _DevService_SaveDevData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "centerms.proto",
}

func init() { proto.RegisterFile("centerms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4a, 0xfb, 0x50,
	0x10, 0xc5, 0xff, 0xf9, 0xf8, 0x57, 0x3a, 0x11, 0x95, 0xa9, 0xd4, 0x50, 0x10, 0x4b, 0x36, 0x76,
	0x63, 0xc0, 0xba, 0x50, 0x5c, 0x09, 0xc6, 0x95, 0xe8, 0x22, 0x7d, 0x82, 0x69, 0x9c, 0x45, 0xd0,
	0x7c, 0x98, 0x3b, 0x09, 0xf4, 0xd5, 0x7c, 0x3a, 0xb9, 0x37, 0x37, 0xa5, 0xa5, 0x3b, 0x77, 0x67,
	0x4e, 0xee, 0x3d, 0x67, 0x7e, 0x49, 0xe0, 0x24, 0xe3, 0x52, 0xb8, 0x29, 0x54, 0x5c, 0x37, 0x95,
	0x54, 0xe8, 0xd6, 0xeb, 0xe8, 0x19, 0x8e, 0x12, 0xee, 0xde, 0x58, 0x08, 0x11, 0x7c, 0xd9, 0xd4,
	0x1c, 0x3a, 0x73, 0x67, 0x31, 0x4e, 0x8d, 0xd6, 0x5e, 0x49, 0x05, 0x87, 0x6e, 0xef, 0x69, 0x8d,
	0x67, 0xe0, 0x15, 0x94, 0x85, 0x9e, 0xb1, 0xb4, 0x8c, 0x5a, 0xc0, 0x15, 0x75, 0x9c, 0x70, 0x97,
	0x90, 0x50, 0xca, 0xdf, 0x2d, 0x2b, 0xc1, 0x29, 0x8c, 0x28, 0x93, 0xbc, 0x2a, 0x6d, 0xa2, 0x9d,
	0x4c, 0x4f, 0x6e, 0x33, 0xbd, 0xd4, 0x68, 0xbc, 0x02, 0xbf, 0x60, 0x21, 0x13, 0x1a, 0x2c, 0x83,
	0xb8, 0x5e, 0xc7, 0x76, 0xad, 0xd4, 0x3c, 0xd0, 0x97, 0x3e, 0x48, 0x28, 0xf4, 0xe7, 0xce, 0xe2,
	0x38, 0x35, 0x3a, 0xba, 0x81, 0xc9, 0x5e, 0xad, 0xaa, 0xab, 0x52, 0xb1, 0xee, 0x55, 0x42, 0xd2,
	0xaa, 0xa1, 0xb7, 0x9f, 0xa2, 0x57, 0x38, 0x1d, 0x32, 0x87, 0xa3, 0x7f, 0x47, 0xfe, 0x71, 0x4c,
	0xda, 0x5e, 0x31, 0x82, 0xaf, 0x72, 0xd9, 0xa6, 0x69, 0x8d, 0xd7, 0x16, 0xcc, 0x35, 0x60, 0x93,
	0x5d, 0x30, 0x7b, 0xcd, 0x02, 0xde, 0x5a, 0x40, 0x6f, 0xee, 0x2d, 0x82, 0xe5, 0xa5, 0x3d, 0xb8,
	0x9b, 0x1f, 0xeb, 0xe1, 0xa5, 0x94, 0x66, 0xd3, 0xf3, 0xcf, 0xee, 0x61, 0xbc, 0xb5, 0xf4, 0x8a,
	0x9f, 0xbc, 0xb1, 0xdd, 0x5a, 0xe2, 0x39, 0xfc, 0xef, 0xe8, 0xab, 0x1d, 0x48, 0xfa, 0xe1, 0xd1,
	0x7d, 0x70, 0x96, 0xef, 0x00, 0x09, 0x77, 0x2b, 0x6e, 0xba, 0x3c, 0x63, 0x7c, 0x82, 0x60, 0xe7,
	0x35, 0xe2, 0x54, 0x57, 0x1f, 0x7e, 0xce, 0xd9, 0xc5, 0x81, 0xdf, 0xaf, 0x15, 0xfd, 0x5b, 0x8f,
	0xcc, 0xff, 0x74, 0xf7, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xd2, 0x6d, 0x6a, 0x61, 0x02, 0x00,
	0x00,
}
