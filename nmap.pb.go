// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nmap.proto

package grpcnmapservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ScanRequest struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	PortSpec             string   `protobuf:"bytes,2,opt,name=portSpec,proto3" json:"portSpec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanRequest) Reset()         { *m = ScanRequest{} }
func (m *ScanRequest) String() string { return proto.CompactTextString(m) }
func (*ScanRequest) ProtoMessage()    {}
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aaa3a40e8dcfc7a, []int{0}
}

func (m *ScanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRequest.Unmarshal(m, b)
}
func (m *ScanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRequest.Marshal(b, m, deterministic)
}
func (m *ScanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRequest.Merge(m, src)
}
func (m *ScanRequest) XXX_Size() int {
	return xxx_messageInfo_ScanRequest.Size(m)
}
func (m *ScanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRequest proto.InternalMessageInfo

func (m *ScanRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ScanRequest) GetPortSpec() string {
	if m != nil {
		return m.PortSpec
	}
	return ""
}

type ScanReply struct {
	State                string            `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Addresses            []string          `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Hostnames            []string          `protobuf:"bytes,3,rep,name=hostnames,proto3" json:"hostnames,omitempty"`
	Ports                []*ScanReply_Port `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ScanReply) Reset()         { *m = ScanReply{} }
func (m *ScanReply) String() string { return proto.CompactTextString(m) }
func (*ScanReply) ProtoMessage()    {}
func (*ScanReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aaa3a40e8dcfc7a, []int{1}
}

func (m *ScanReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanReply.Unmarshal(m, b)
}
func (m *ScanReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanReply.Marshal(b, m, deterministic)
}
func (m *ScanReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanReply.Merge(m, src)
}
func (m *ScanReply) XXX_Size() int {
	return xxx_messageInfo_ScanReply.Size(m)
}
func (m *ScanReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanReply.DiscardUnknown(m)
}

var xxx_messageInfo_ScanReply proto.InternalMessageInfo

func (m *ScanReply) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ScanReply) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ScanReply) GetHostnames() []string {
	if m != nil {
		return m.Hostnames
	}
	return nil
}

func (m *ScanReply) GetPorts() []*ScanReply_Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type ScanReply_Port struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanReply_Port) Reset()         { *m = ScanReply_Port{} }
func (m *ScanReply_Port) String() string { return proto.CompactTextString(m) }
func (*ScanReply_Port) ProtoMessage()    {}
func (*ScanReply_Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aaa3a40e8dcfc7a, []int{1, 0}
}

func (m *ScanReply_Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanReply_Port.Unmarshal(m, b)
}
func (m *ScanReply_Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanReply_Port.Marshal(b, m, deterministic)
}
func (m *ScanReply_Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanReply_Port.Merge(m, src)
}
func (m *ScanReply_Port) XXX_Size() int {
	return xxx_messageInfo_ScanReply_Port.Size(m)
}
func (m *ScanReply_Port) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanReply_Port.DiscardUnknown(m)
}

var xxx_messageInfo_ScanReply_Port proto.InternalMessageInfo

func (m *ScanReply_Port) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ScanReply_Port) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ScanReply_Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ScanRequest)(nil), "ch.guengel.nmapservice.ScanRequest")
	proto.RegisterType((*ScanReply)(nil), "ch.guengel.nmapservice.ScanReply")
	proto.RegisterType((*ScanReply_Port)(nil), "ch.guengel.nmapservice.ScanReply.Port")
}

func init() { proto.RegisterFile("nmap.proto", fileDescriptor_7aaa3a40e8dcfc7a) }

var fileDescriptor_7aaa3a40e8dcfc7a = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x6b, 0xfb, 0x30,
	0x10, 0xc5, 0xbf, 0xb6, 0xe5, 0xf0, 0xf5, 0x65, 0x28, 0x1c, 0x25, 0x08, 0xd3, 0xc1, 0x75, 0xa1,
	0xb8, 0x8b, 0x86, 0x74, 0x6d, 0x97, 0x4e, 0x9d, 0x4a, 0x70, 0x96, 0xd2, 0x4d, 0x71, 0x0e, 0xa7,
	0x10, 0xdb, 0xaa, 0x24, 0x17, 0xfa, 0x07, 0xf7, 0xff, 0x28, 0x92, 0x82, 0x93, 0xa1, 0x3f, 0x36,
	0xbd, 0x77, 0x3c, 0xde, 0x47, 0x77, 0x00, 0x7d, 0x27, 0x95, 0x50, 0x7a, 0xb0, 0x03, 0x2e, 0x9a,
	0x9d, 0x68, 0x47, 0xea, 0x5b, 0xda, 0x0b, 0x67, 0x1b, 0xd2, 0xef, 0xaf, 0x0d, 0x95, 0xf7, 0x30,
	0x5f, 0x37, 0xb2, 0xaf, 0xe9, 0x6d, 0x24, 0x63, 0x11, 0x81, 0xed, 0x06, 0x63, 0x79, 0x54, 0x44,
	0x55, 0x56, 0xfb, 0x37, 0xe6, 0xf0, 0x5f, 0x0d, 0xda, 0xae, 0x15, 0x35, 0x3c, 0xf6, 0xfe, 0xa4,
	0xcb, 0xcf, 0x08, 0xb2, 0x90, 0x57, 0xfb, 0x0f, 0x3c, 0x87, 0xd4, 0x58, 0x69, 0xe9, 0x10, 0x0f,
	0x02, 0x2f, 0x20, 0x93, 0xdb, 0xad, 0x26, 0x63, 0xc8, 0xf0, 0xb8, 0x48, 0xaa, 0xac, 0x3e, 0x1a,
	0x6e, 0xea, 0x5a, 0x7a, 0xd9, 0x91, 0xe1, 0x49, 0x98, 0x4e, 0x06, 0xde, 0x41, 0xea, 0xba, 0x0c,
	0x67, 0x45, 0x52, 0xcd, 0x97, 0xd7, 0xe2, 0xfb, 0x6f, 0x88, 0x89, 0x41, 0xac, 0x06, 0x6d, 0xeb,
	0x10, 0xca, 0x1f, 0x81, 0x39, 0x89, 0x0b, 0x98, 0xf5, 0x63, 0xb7, 0x21, 0xed, 0xc1, 0xd2, 0xfa,
	0xa0, 0x8e, 0xbc, 0xf1, 0x29, 0x2f, 0x02, 0x73, 0xe5, 0x3c, 0x09, 0x3b, 0x70, 0xef, 0xe5, 0x33,
	0xb0, 0xa7, 0x4e, 0x2a, 0x5c, 0x01, 0x73, 0x55, 0x78, 0xf5, 0x3b, 0x88, 0x5f, 0x66, 0x7e, 0xf9,
	0x27, 0x6d, 0xf9, 0xef, 0xe1, 0x06, 0x7e, 0x38, 0xcd, 0xcb, 0x59, 0xab, 0x55, 0x73, 0x62, 0x6c,
	0x66, 0xfe, 0x94, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x35, 0xb8, 0x29, 0xd8, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NmapClient is the client API for Nmap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NmapClient interface {
	Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanReply, error)
}

type nmapClient struct {
	cc *grpc.ClientConn
}

func NewNmapClient(cc *grpc.ClientConn) NmapClient {
	return &nmapClient{cc}
}

func (c *nmapClient) Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanReply, error) {
	out := new(ScanReply)
	err := c.cc.Invoke(ctx, "/ch.guengel.nmapservice.Nmap/Scan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NmapServer is the server API for Nmap service.
type NmapServer interface {
	Scan(context.Context, *ScanRequest) (*ScanReply, error)
}

// UnimplementedNmapServer can be embedded to have forward compatible implementations.
type UnimplementedNmapServer struct {
}

func (*UnimplementedNmapServer) Scan(ctx context.Context, req *ScanRequest) (*ScanReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}

func RegisterNmapServer(s *grpc.Server, srv NmapServer) {
	s.RegisterService(&_Nmap_serviceDesc, srv)
}

func _Nmap_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NmapServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ch.guengel.nmapservice.Nmap/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NmapServer).Scan(ctx, req.(*ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nmap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ch.guengel.nmapservice.Nmap",
	HandlerType: (*NmapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Scan",
			Handler:    _Nmap_Scan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nmap.proto",
}
