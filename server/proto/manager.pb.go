// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/proto/manager.proto

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	server/proto/manager.proto

It has these top-level messages:
	ScheduleSavingRequest
	SaveVolumeRequest
	Error
*/
package server

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

type ScheduleSavingRequest struct {
	Schedule string             `protobuf:"bytes,1,opt,name=schedule" json:"schedule,omitempty"`
	Volume   *SaveVolumeRequest `protobuf:"bytes,2,opt,name=volume" json:"volume,omitempty"`
}

func (m *ScheduleSavingRequest) Reset()                    { *m = ScheduleSavingRequest{} }
func (m *ScheduleSavingRequest) String() string            { return proto.CompactTextString(m) }
func (*ScheduleSavingRequest) ProtoMessage()               {}
func (*ScheduleSavingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ScheduleSavingRequest) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *ScheduleSavingRequest) GetVolume() *SaveVolumeRequest {
	if m != nil {
		return m.Volume
	}
	return nil
}

type SaveVolumeRequest struct {
	VolumeName  string `protobuf:"bytes,1,opt,name=volumeName" json:"volumeName,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
}

func (m *SaveVolumeRequest) Reset()                    { *m = SaveVolumeRequest{} }
func (m *SaveVolumeRequest) String() string            { return proto.CompactTextString(m) }
func (*SaveVolumeRequest) ProtoMessage()               {}
func (*SaveVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SaveVolumeRequest) GetVolumeName() string {
	if m != nil {
		return m.VolumeName
	}
	return ""
}

func (m *SaveVolumeRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type Error struct {
	Code    int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Error) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ScheduleSavingRequest)(nil), "server.ScheduleSavingRequest")
	proto.RegisterType((*SaveVolumeRequest)(nil), "server.SaveVolumeRequest")
	proto.RegisterType((*Error)(nil), "server.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Manager service

type ManagerClient interface {
	SaveVolume(ctx context.Context, in *SaveVolumeRequest, opts ...grpc.CallOption) (*Error, error)
	ScheduleSaving(ctx context.Context, in *ScheduleSavingRequest, opts ...grpc.CallOption) (*Error, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) SaveVolume(ctx context.Context, in *SaveVolumeRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := grpc.Invoke(ctx, "/server.Manager/SaveVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ScheduleSaving(ctx context.Context, in *ScheduleSavingRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := grpc.Invoke(ctx, "/server.Manager/ScheduleSaving", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Manager service

type ManagerServer interface {
	SaveVolume(context.Context, *SaveVolumeRequest) (*Error, error)
	ScheduleSaving(context.Context, *ScheduleSavingRequest) (*Error, error)
}

func RegisterManagerServer(s *grpc.Server, srv ManagerServer) {
	s.RegisterService(&_Manager_serviceDesc, srv)
}

func _Manager_SaveVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).SaveVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Manager/SaveVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).SaveVolume(ctx, req.(*SaveVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ScheduleSaving_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleSavingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ScheduleSaving(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Manager/ScheduleSaving",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ScheduleSaving(ctx, req.(*ScheduleSavingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Manager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveVolume",
			Handler:    _Manager_SaveVolume_Handler,
		},
		{
			MethodName: "ScheduleSaving",
			Handler:    _Manager_ScheduleSaving_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/proto/manager.proto",
}

func init() { proto.RegisterFile("server/proto/manager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xad, 0x7f, 0xba, 0xee, 0x2c, 0x0a, 0x0e, 0x08, 0xb5, 0xa0, 0x94, 0x9c, 0xf6, 0xd4,
	0xc5, 0x15, 0xc1, 0xa3, 0x17, 0x8f, 0x7a, 0x48, 0xd1, 0x7b, 0xdc, 0x8e, 0xb5, 0xb0, 0x49, 0x34,
	0x49, 0xf3, 0x0d, 0xfc, 0xde, 0x42, 0x92, 0xae, 0x2b, 0xab, 0xb7, 0xbe, 0xf9, 0x75, 0xde, 0x4b,
	0x5e, 0xa0, 0xb4, 0x64, 0x3c, 0x99, 0xc5, 0x87, 0xd1, 0x4e, 0x2f, 0xa4, 0x50, 0xa2, 0x23, 0x53,
	0x07, 0x85, 0x79, 0x64, 0xec, 0x0d, 0xce, 0x9b, 0xd5, 0x3b, 0xb5, 0xc3, 0x9a, 0x1a, 0xe1, 0x7b,
	0xd5, 0x71, 0xfa, 0x1c, 0xc8, 0x3a, 0x2c, 0xe1, 0xd8, 0x26, 0x50, 0x64, 0x55, 0x36, 0x9f, 0xf2,
	0x8d, 0xc6, 0x6b, 0xc8, 0xbd, 0x5e, 0x0f, 0x92, 0x8a, 0xfd, 0x2a, 0x9b, 0xcf, 0x96, 0x17, 0x75,
	0x74, 0xab, 0x1b, 0xe1, 0xe9, 0x25, 0x90, 0x64, 0xc3, 0xd3, 0x8f, 0xec, 0x19, 0xce, 0x76, 0x20,
	0x5e, 0x01, 0x44, 0xfc, 0x24, 0xe4, 0x98, 0xb2, 0x35, 0xc1, 0x0a, 0x66, 0x2d, 0x59, 0xd7, 0x2b,
	0xe1, 0x7a, 0xad, 0x42, 0xd8, 0x94, 0x6f, 0x8f, 0xd8, 0x2d, 0x1c, 0x3d, 0x18, 0xa3, 0x0d, 0x22,
	0x1c, 0xae, 0x74, 0x1b, 0x4d, 0x0e, 0x78, 0xf8, 0xc6, 0x02, 0x26, 0x92, 0xac, 0x15, 0x1d, 0xa5,
	0xd5, 0x51, 0x2e, 0xbf, 0x32, 0x98, 0x3c, 0xc6, 0x3e, 0xf0, 0x0e, 0xe0, 0xe7, 0x64, 0xf8, 0xff,
	0x55, 0xca, 0x93, 0x11, 0x85, 0x44, 0xb6, 0x87, 0xf7, 0x70, 0xfa, 0xbb, 0x3b, 0xbc, 0xdc, 0x6c,
	0xff, 0xd5, 0xe9, 0x8e, 0xc3, 0x6b, 0x1e, 0x1e, 0xe3, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x92,
	0xe7, 0x2d, 0x48, 0xaa, 0x01, 0x00, 0x00,
}
