// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control_plane_ca.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	control_plane_ca.proto

It has these top-level messages:
	SignCsrRequest
	SignCsrResponse
	GenerateCsrRequest
	GenerateCsrResponse
	FetchCertificateRequest
	FetchCertificateResponse
	LoadCertificateRequest
	LoadCertificateResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type SignCsrRequest struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *SignCsrRequest) Reset()                    { *m = SignCsrRequest{} }
func (m *SignCsrRequest) String() string            { return proto1.CompactTextString(m) }
func (*SignCsrRequest) ProtoMessage()               {}
func (*SignCsrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignCsrRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type SignCsrResponse struct {
	SignedCertificate []byte `protobuf:"bytes,1,opt,name=signedCertificate,proto3" json:"signedCertificate,omitempty"`
}

func (m *SignCsrResponse) Reset()                    { *m = SignCsrResponse{} }
func (m *SignCsrResponse) String() string            { return proto1.CompactTextString(m) }
func (*SignCsrResponse) ProtoMessage()               {}
func (*SignCsrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignCsrResponse) GetSignedCertificate() []byte {
	if m != nil {
		return m.SignedCertificate
	}
	return nil
}

type GenerateCsrRequest struct {
}

func (m *GenerateCsrRequest) Reset()                    { *m = GenerateCsrRequest{} }
func (m *GenerateCsrRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenerateCsrRequest) ProtoMessage()               {}
func (*GenerateCsrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GenerateCsrResponse struct {
	Csr []byte `protobuf:"bytes,1,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (m *GenerateCsrResponse) Reset()                    { *m = GenerateCsrResponse{} }
func (m *GenerateCsrResponse) String() string            { return proto1.CompactTextString(m) }
func (*GenerateCsrResponse) ProtoMessage()               {}
func (*GenerateCsrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GenerateCsrResponse) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

type FetchCertificateRequest struct {
}

func (m *FetchCertificateRequest) Reset()                    { *m = FetchCertificateRequest{} }
func (m *FetchCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchCertificateRequest) ProtoMessage()               {}
func (*FetchCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type FetchCertificateResponse struct {
	StoredIntermediateCert []byte `protobuf:"bytes,1,opt,name=storedIntermediateCert,proto3" json:"storedIntermediateCert,omitempty"`
}

func (m *FetchCertificateResponse) Reset()                    { *m = FetchCertificateResponse{} }
func (m *FetchCertificateResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchCertificateResponse) ProtoMessage()               {}
func (*FetchCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FetchCertificateResponse) GetStoredIntermediateCert() []byte {
	if m != nil {
		return m.StoredIntermediateCert
	}
	return nil
}

type LoadCertificateRequest struct {
	SignedIntermediateCert []byte `protobuf:"bytes,1,opt,name=signedIntermediateCert,proto3" json:"signedIntermediateCert,omitempty"`
}

func (m *LoadCertificateRequest) Reset()                    { *m = LoadCertificateRequest{} }
func (m *LoadCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*LoadCertificateRequest) ProtoMessage()               {}
func (*LoadCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoadCertificateRequest) GetSignedIntermediateCert() []byte {
	if m != nil {
		return m.SignedIntermediateCert
	}
	return nil
}

type LoadCertificateResponse struct {
}

func (m *LoadCertificateResponse) Reset()                    { *m = LoadCertificateResponse{} }
func (m *LoadCertificateResponse) String() string            { return proto1.CompactTextString(m) }
func (*LoadCertificateResponse) ProtoMessage()               {}
func (*LoadCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto1.RegisterType((*SignCsrRequest)(nil), "proto.SignCsrRequest")
	proto1.RegisterType((*SignCsrResponse)(nil), "proto.SignCsrResponse")
	proto1.RegisterType((*GenerateCsrRequest)(nil), "proto.GenerateCsrRequest")
	proto1.RegisterType((*GenerateCsrResponse)(nil), "proto.GenerateCsrResponse")
	proto1.RegisterType((*FetchCertificateRequest)(nil), "proto.FetchCertificateRequest")
	proto1.RegisterType((*FetchCertificateResponse)(nil), "proto.FetchCertificateResponse")
	proto1.RegisterType((*LoadCertificateRequest)(nil), "proto.LoadCertificateRequest")
	proto1.RegisterType((*LoadCertificateResponse)(nil), "proto.LoadCertificateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ControlPlaneCA service

type ControlPlaneCAClient interface {
	SignCsr(ctx context.Context, in *SignCsrRequest, opts ...grpc.CallOption) (*SignCsrResponse, error)
	GenerateCsr(ctx context.Context, in *GenerateCsrRequest, opts ...grpc.CallOption) (*GenerateCsrResponse, error)
	FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error)
	LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error)
}

type controlPlaneCAClient struct {
	cc *grpc.ClientConn
}

func NewControlPlaneCAClient(cc *grpc.ClientConn) ControlPlaneCAClient {
	return &controlPlaneCAClient{cc}
}

func (c *controlPlaneCAClient) SignCsr(ctx context.Context, in *SignCsrRequest, opts ...grpc.CallOption) (*SignCsrResponse, error) {
	out := new(SignCsrResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/SignCsr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) GenerateCsr(ctx context.Context, in *GenerateCsrRequest, opts ...grpc.CallOption) (*GenerateCsrResponse, error) {
	out := new(GenerateCsrResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/GenerateCsr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) FetchCertificate(ctx context.Context, in *FetchCertificateRequest, opts ...grpc.CallOption) (*FetchCertificateResponse, error) {
	out := new(FetchCertificateResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/FetchCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneCAClient) LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error) {
	out := new(LoadCertificateResponse)
	err := grpc.Invoke(ctx, "/proto.ControlPlaneCA/LoadCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ControlPlaneCA service

type ControlPlaneCAServer interface {
	SignCsr(context.Context, *SignCsrRequest) (*SignCsrResponse, error)
	GenerateCsr(context.Context, *GenerateCsrRequest) (*GenerateCsrResponse, error)
	FetchCertificate(context.Context, *FetchCertificateRequest) (*FetchCertificateResponse, error)
	LoadCertificate(context.Context, *LoadCertificateRequest) (*LoadCertificateResponse, error)
}

func RegisterControlPlaneCAServer(s *grpc.Server, srv ControlPlaneCAServer) {
	s.RegisterService(&_ControlPlaneCA_serviceDesc, srv)
}

func _ControlPlaneCA_SignCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).SignCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/SignCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).SignCsr(ctx, req.(*SignCsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_GenerateCsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).GenerateCsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/GenerateCsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).GenerateCsr(ctx, req.(*GenerateCsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_FetchCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).FetchCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/FetchCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).FetchCertificate(ctx, req.(*FetchCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlaneCA_LoadCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneCAServer).LoadCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ControlPlaneCA/LoadCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneCAServer).LoadCertificate(ctx, req.(*LoadCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlPlaneCA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ControlPlaneCA",
	HandlerType: (*ControlPlaneCAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignCsr",
			Handler:    _ControlPlaneCA_SignCsr_Handler,
		},
		{
			MethodName: "GenerateCsr",
			Handler:    _ControlPlaneCA_GenerateCsr_Handler,
		},
		{
			MethodName: "FetchCertificate",
			Handler:    _ControlPlaneCA_FetchCertificate_Handler,
		},
		{
			MethodName: "LoadCertificate",
			Handler:    _ControlPlaneCA_LoadCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control_plane_ca.proto",
}

func init() { proto1.RegisterFile("control_plane_ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xa9, 0xa2, 0xc2, 0x28, 0x6d, 0x5d, 0x35, 0xb6, 0x01, 0xab, 0xec, 0x45, 0x0f, 0xd2,
	0x83, 0x82, 0x78, 0x13, 0x89, 0x28, 0x82, 0x87, 0x90, 0xfe, 0x80, 0x12, 0x93, 0xb1, 0x06, 0xea,
	0x6e, 0xdc, 0x1d, 0x7f, 0x9b, 0x7f, 0x4f, 0x92, 0x8c, 0xda, 0x66, 0xb3, 0xa7, 0x84, 0x7d, 0xb3,
	0xdf, 0xdb, 0x79, 0x0f, 0x82, 0x4c, 0x2b, 0x32, 0x7a, 0x39, 0x2f, 0x97, 0xa9, 0xc2, 0x79, 0x96,
	0x4e, 0x4b, 0xa3, 0x49, 0x8b, 0xad, 0xfa, 0x23, 0x25, 0xf4, 0x67, 0xc5, 0x42, 0x45, 0xd6, 0x24,
	0xf8, 0xf9, 0x85, 0x96, 0xc4, 0x10, 0x36, 0x33, 0x6b, 0x46, 0xbd, 0xb3, 0xde, 0xc5, 0x5e, 0x52,
	0xfd, 0xca, 0x3b, 0x18, 0xfc, 0xcd, 0xd8, 0x52, 0x2b, 0x8b, 0xe2, 0x12, 0xf6, 0x6d, 0xb1, 0x50,
	0x98, 0x47, 0x68, 0xa8, 0x78, 0x2b, 0xb2, 0x94, 0x90, 0xaf, 0xb8, 0x82, 0x3c, 0x04, 0xf1, 0x84,
	0x0a, 0x4d, 0x4a, 0xf8, 0x6f, 0x24, 0xcf, 0xe1, 0x60, 0xed, 0x94, 0xd1, 0xae, 0xff, 0x18, 0x8e,
	0x1f, 0x91, 0xb2, 0xf7, 0x15, 0xe4, 0x2f, 0x23, 0x81, 0x91, 0x2b, 0x31, 0xe8, 0x06, 0x02, 0x4b,
	0xda, 0x60, 0xfe, 0xac, 0x08, 0xcd, 0x07, 0xe6, 0x45, 0xe5, 0x84, 0x86, 0x98, 0xed, 0x51, 0x65,
	0x0c, 0xc1, 0x8b, 0x4e, 0x73, 0xd7, 0xad, 0x26, 0xd6, 0xcb, 0x79, 0x89, 0x9d, 0x6a, 0xb5, 0x80,
	0x43, 0x6c, 0x1e, 0x79, 0xf5, 0xbd, 0x01, 0xfd, 0xa8, 0x69, 0x28, 0xae, 0x0a, 0x8a, 0xee, 0xc5,
	0x2d, 0xec, 0x70, 0xdc, 0xe2, 0xa8, 0x29, 0x6b, 0xba, 0x5e, 0x51, 0x18, 0xb4, 0x8f, 0x79, 0xe3,
	0x07, 0xd8, 0x5d, 0x49, 0x54, 0x8c, 0x79, 0xcc, 0xcd, 0x3e, 0x0c, 0xbb, 0x24, 0xa6, 0xcc, 0x60,
	0xd8, 0xce, 0x54, 0x4c, 0x78, 0xde, 0xd3, 0x43, 0x78, 0xea, 0xd5, 0x19, 0x1a, 0xc3, 0xa0, 0x15,
	0x81, 0x38, 0xe1, 0x3b, 0xdd, 0x61, 0x87, 0x13, 0x9f, 0xdc, 0x10, 0x5f, 0xb7, 0x6b, 0xf9, 0xfa,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa4, 0x1b, 0x25, 0xe1, 0x02, 0x00, 0x00,
}
