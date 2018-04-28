// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/proto/manager.proto

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	server/proto/manager.proto

It has these top-level messages:
	EntriesList
	Entry
	Empty
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

type EntriesList struct {
	Entries []*Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *EntriesList) Reset()                    { *m = EntriesList{} }
func (m *EntriesList) String() string            { return proto.CompactTextString(m) }
func (*EntriesList) ProtoMessage()               {}
func (*EntriesList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EntriesList) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Entry struct {
	Next string `protobuf:"bytes,1,opt,name=next" json:"next,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Entry) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ScheduleSavingRequest struct {
	Schedule string             `protobuf:"bytes,1,opt,name=schedule" json:"schedule,omitempty"`
	Volume   *SaveVolumeRequest `protobuf:"bytes,2,opt,name=volume" json:"volume,omitempty"`
}

func (m *ScheduleSavingRequest) Reset()                    { *m = ScheduleSavingRequest{} }
func (m *ScheduleSavingRequest) String() string            { return proto.CompactTextString(m) }
func (*ScheduleSavingRequest) ProtoMessage()               {}
func (*ScheduleSavingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*SaveVolumeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
	proto.RegisterType((*EntriesList)(nil), "server.EntriesList")
	proto.RegisterType((*Entry)(nil), "server.Entry")
	proto.RegisterType((*Empty)(nil), "server.Empty")
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
	ListEntries(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EntriesList, error)
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

func (c *managerClient) ListEntries(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EntriesList, error) {
	out := new(EntriesList)
	err := grpc.Invoke(ctx, "/server.Manager/ListEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Manager service

type ManagerServer interface {
	SaveVolume(context.Context, *SaveVolumeRequest) (*Error, error)
	ScheduleSaving(context.Context, *ScheduleSavingRequest) (*Error, error)
	ListEntries(context.Context, *Empty) (*EntriesList, error)
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

func _Manager_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Manager/ListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ListEntries(ctx, req.(*Empty))
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
		{
			MethodName: "ListEntries",
			Handler:    _Manager_ListEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/proto/manager.proto",
}

func init() { proto.RegisterFile("server/proto/manager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0xa5, 0x3f, 0x7e, 0x50, 0x99, 0x46, 0x13, 0xc7, 0x98, 0xd4, 0x1a, 0x4d, 0xb3, 0x17, 0x39,
	0x41, 0x84, 0x68, 0x3c, 0x7a, 0xe1, 0xa6, 0x1e, 0x4a, 0xf4, 0xbe, 0xc2, 0x88, 0x4d, 0xe8, 0x2e,
	0xee, 0x2e, 0x8d, 0x7c, 0x35, 0x3f, 0x9d, 0xe9, 0xee, 0x16, 0x6a, 0xd0, 0xdb, 0xbe, 0x79, 0xf3,
	0xf7, 0xbd, 0x85, 0x44, 0x93, 0x2a, 0x49, 0x0d, 0x57, 0x4a, 0x1a, 0x39, 0x2c, 0xb8, 0xe0, 0x0b,
	0x52, 0x03, 0x8b, 0xb0, 0xeb, 0x38, 0x76, 0x0b, 0xd1, 0x44, 0x18, 0x95, 0x93, 0x7e, 0xc8, 0xb5,
	0xc1, 0x2b, 0x08, 0xc9, 0xc1, 0x38, 0x48, 0xdb, 0xfd, 0x68, 0x74, 0x38, 0x70, 0x89, 0x83, 0x2a,
	0x6b, 0x93, 0xd5, 0x2c, 0x3b, 0x87, 0x8e, 0x8d, 0x20, 0xc2, 0x7f, 0x41, 0x9f, 0x26, 0x0e, 0xd2,
	0xa0, 0xdf, 0xcb, 0xec, 0x9b, 0x85, 0xd0, 0x99, 0x14, 0x2b, 0xb3, 0x61, 0x6f, 0x70, 0x3a, 0x9d,
	0xbd, 0xd3, 0x7c, 0xbd, 0xa4, 0x29, 0x2f, 0x73, 0xb1, 0xc8, 0xe8, 0x63, 0x4d, 0xda, 0x60, 0x02,
	0x07, 0xda, 0x13, 0xbe, 0x72, 0x8b, 0xf1, 0x1a, 0xba, 0xa5, 0x5c, 0xae, 0x0b, 0x8a, 0xff, 0xa5,
	0x41, 0x3f, 0x1a, 0x9d, 0xd5, 0x2b, 0x4c, 0x79, 0x49, 0x2f, 0x96, 0xf1, 0x6d, 0x32, 0x9f, 0xc8,
	0x9e, 0xe1, 0x78, 0x8f, 0xc4, 0x4b, 0x00, 0x47, 0x3f, 0xf1, 0xa2, 0x9e, 0xd2, 0x88, 0x60, 0x0a,
	0xd1, 0x9c, 0xb4, 0xc9, 0x05, 0x37, 0xb9, 0x14, 0x76, 0x58, 0x2f, 0x6b, 0x86, 0xd8, 0x0d, 0x74,
	0x26, 0x4a, 0x49, 0x55, 0x1d, 0x39, 0x93, 0x73, 0xd7, 0xa4, 0x9d, 0xd9, 0x37, 0xc6, 0x10, 0x16,
	0xa4, 0x35, 0x5f, 0x90, 0x2f, 0xad, 0xe1, 0xe8, 0x2b, 0x80, 0xf0, 0xd1, 0xa9, 0x8d, 0x77, 0x00,
	0xbb, 0xcd, 0xf0, 0xef, 0x53, 0x92, 0x9d, 0xd0, 0xd5, 0x44, 0xd6, 0xc2, 0x7b, 0x38, 0xfa, 0xa9,
	0x1d, 0x5e, 0x6c, 0xab, 0x7f, 0xd3, 0x74, 0xbf, 0xc3, 0x18, 0xa2, 0xca, 0x54, 0xef, 0x2f, 0xee,
	0xf8, 0xca, 0x9b, 0xe4, 0xa4, 0xe9, 0xac, 0xf7, 0x9f, 0xb5, 0x5e, 0xbb, 0xf6, 0x7f, 0x8c, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x76, 0x7a, 0x23, 0xff, 0x3d, 0x02, 0x00, 0x00,
}
