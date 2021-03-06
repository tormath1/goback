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
	Volume string `protobuf:"bytes,1,opt,name=volume" json:"volume,omitempty"`
	Cron   string `protobuf:"bytes,2,opt,name=cron" json:"cron,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Entry) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *Entry) GetCron() string {
	if m != nil {
		return m.Cron
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

func init() {
	proto.RegisterType((*EntriesList)(nil), "server.EntriesList")
	proto.RegisterType((*Entry)(nil), "server.Entry")
	proto.RegisterType((*Empty)(nil), "server.Empty")
	proto.RegisterType((*ScheduleSavingRequest)(nil), "server.ScheduleSavingRequest")
	proto.RegisterType((*SaveVolumeRequest)(nil), "server.SaveVolumeRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Manager service

type ManagerClient interface {
	SaveVolume(ctx context.Context, in *SaveVolumeRequest, opts ...grpc.CallOption) (*Empty, error)
	ScheduleSaving(ctx context.Context, in *ScheduleSavingRequest, opts ...grpc.CallOption) (*Empty, error)
	ListEntries(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EntriesList, error)
}

type managerClient struct {
	cc *grpc.ClientConn
}

func NewManagerClient(cc *grpc.ClientConn) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) SaveVolume(ctx context.Context, in *SaveVolumeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/server.Manager/SaveVolume", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ScheduleSaving(ctx context.Context, in *ScheduleSavingRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
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
	SaveVolume(context.Context, *SaveVolumeRequest) (*Empty, error)
	ScheduleSaving(context.Context, *ScheduleSavingRequest) (*Empty, error)
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
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x4e, 0x83, 0x40,
	0x14, 0x2c, 0x6a, 0x41, 0x1f, 0xd1, 0xc4, 0x67, 0x34, 0x48, 0xa2, 0x21, 0x7b, 0xb1, 0x27, 0x1a,
	0x21, 0x31, 0x1e, 0xbd, 0xf4, 0xa6, 0x1e, 0x20, 0x7a, 0x5f, 0xdb, 0x67, 0x25, 0x29, 0x50, 0x77,
	0x17, 0x92, 0xfe, 0x9a, 0x5f, 0x67, 0x58, 0x16, 0x4a, 0x53, 0xbd, 0x31, 0x33, 0xcb, 0xec, 0x30,
	0x03, 0xf8, 0x92, 0x44, 0x4d, 0x62, 0xba, 0x16, 0xa5, 0x2a, 0xa7, 0x39, 0x2f, 0xf8, 0x92, 0x44,
	0xa8, 0x11, 0xda, 0xad, 0xc6, 0x1e, 0xc0, 0x9d, 0x15, 0x4a, 0x64, 0x24, 0x9f, 0x33, 0xa9, 0xf0,
	0x0e, 0x1c, 0x6a, 0xa1, 0x67, 0x05, 0x87, 0x13, 0x37, 0x3a, 0x0d, 0xdb, 0x83, 0x61, 0x73, 0x6a,
	0x93, 0x74, 0x2a, 0x8b, 0x61, 0xac, 0x19, 0xbc, 0x02, 0xbb, 0x2e, 0x57, 0x55, 0x4e, 0x9e, 0x15,
	0x58, 0x93, 0x93, 0xc4, 0x20, 0x44, 0x38, 0x9a, 0x8b, 0xb2, 0xf0, 0x0e, 0x34, 0xab, 0x9f, 0x99,
	0x03, 0xe3, 0x59, 0xbe, 0x56, 0x1b, 0xf6, 0x09, 0x97, 0xe9, 0xfc, 0x8b, 0x16, 0xd5, 0x8a, 0x52,
	0x5e, 0x67, 0xc5, 0x32, 0xa1, 0xef, 0x8a, 0xa4, 0x42, 0x1f, 0x8e, 0xa5, 0x11, 0x8c, 0x5f, 0x8f,
	0xf1, 0xbe, 0xbf, 0xa9, 0xf1, 0x74, 0xa3, 0xeb, 0x2e, 0x5a, 0xca, 0x6b, 0x7a, 0xd7, 0x8a, 0xb1,
	0xe9, 0x42, 0xb0, 0x37, 0x38, 0xdf, 0x13, 0xf1, 0x16, 0xa0, 0x95, 0x5f, 0x79, 0x9f, 0x7a, 0xc0,
	0x60, 0x00, 0xee, 0x82, 0xa4, 0xca, 0x0a, 0xae, 0xb2, 0xfe, 0x03, 0x86, 0x54, 0xf4, 0x63, 0x81,
	0xf3, 0xd2, 0xd6, 0x89, 0x8f, 0x00, 0xdb, 0x2b, 0xf0, 0xff, 0x4c, 0xfe, 0xb6, 0x49, 0x5d, 0xc1,
	0x08, 0x9f, 0xe0, 0x6c, 0xb7, 0x04, 0xbc, 0xe9, 0xdf, 0xfe, 0xab, 0x9c, 0x7d, 0x87, 0x18, 0xdc,
	0x66, 0x35, 0x33, 0x20, 0xee, 0xea, 0xfe, 0xc5, 0x70, 0x3a, 0x33, 0x30, 0x1b, 0x7d, 0xd8, 0xfa,
	0x07, 0x88, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x2e, 0xe0, 0x4d, 0x1e, 0x02, 0x00, 0x00,
}
