// Code generated by protoc-gen-go. DO NOT EDIT.
// source: patients.proto

package patientinfoapi

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetPatientOnUidRequest struct {
	PatientUid           string   `protobuf:"bytes,1,opt,name=patientUid,proto3" json:"patientUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPatientOnUidRequest) Reset()         { *m = GetPatientOnUidRequest{} }
func (m *GetPatientOnUidRequest) String() string { return proto.CompactTextString(m) }
func (*GetPatientOnUidRequest) ProtoMessage()    {}
func (*GetPatientOnUidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0a59435314f582, []int{0}
}

func (m *GetPatientOnUidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPatientOnUidRequest.Unmarshal(m, b)
}
func (m *GetPatientOnUidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPatientOnUidRequest.Marshal(b, m, deterministic)
}
func (m *GetPatientOnUidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPatientOnUidRequest.Merge(m, src)
}
func (m *GetPatientOnUidRequest) XXX_Size() int {
	return xxx_messageInfo_GetPatientOnUidRequest.Size(m)
}
func (m *GetPatientOnUidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPatientOnUidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPatientOnUidRequest proto.InternalMessageInfo

func (m *GetPatientOnUidRequest) GetPatientUid() string {
	if m != nil {
		return m.PatientUid
	}
	return ""
}

type GetPatientOnUidReply struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Patient              *Patient `protobuf:"bytes,2,opt,name=patient,proto3" json:"patient,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPatientOnUidReply) Reset()         { *m = GetPatientOnUidReply{} }
func (m *GetPatientOnUidReply) String() string { return proto.CompactTextString(m) }
func (*GetPatientOnUidReply) ProtoMessage()    {}
func (*GetPatientOnUidReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0a59435314f582, []int{1}
}

func (m *GetPatientOnUidReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPatientOnUidReply.Unmarshal(m, b)
}
func (m *GetPatientOnUidReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPatientOnUidReply.Marshal(b, m, deterministic)
}
func (m *GetPatientOnUidReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPatientOnUidReply.Merge(m, src)
}
func (m *GetPatientOnUidReply) XXX_Size() int {
	return xxx_messageInfo_GetPatientOnUidReply.Size(m)
}
func (m *GetPatientOnUidReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPatientOnUidReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetPatientOnUidReply proto.InternalMessageInfo

func (m *GetPatientOnUidReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GetPatientOnUidReply) GetPatient() *Patient {
	if m != nil {
		return m.Patient
	}
	return nil
}

type Patient struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	AddressLine          string   `protobuf:"bytes,3,opt,name=addressLine,proto3" json:"addressLine,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	EmailAddress         string   `protobuf:"bytes,5,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Patient) Reset()         { *m = Patient{} }
func (m *Patient) String() string { return proto.CompactTextString(m) }
func (*Patient) ProtoMessage()    {}
func (*Patient) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0a59435314f582, []int{2}
}

func (m *Patient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Patient.Unmarshal(m, b)
}
func (m *Patient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Patient.Marshal(b, m, deterministic)
}
func (m *Patient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Patient.Merge(m, src)
}
func (m *Patient) XXX_Size() int {
	return xxx_messageInfo_Patient.Size(m)
}
func (m *Patient) XXX_DiscardUnknown() {
	xxx_messageInfo_Patient.DiscardUnknown(m)
}

var xxx_messageInfo_Patient proto.InternalMessageInfo

func (m *Patient) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Patient) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Patient) GetAddressLine() string {
	if m != nil {
		return m.AddressLine
	}
	return ""
}

func (m *Patient) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Patient) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details              string   `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_df0a59435314f582, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
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

func (m *Error) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPatientOnUidRequest)(nil), "patientinfoapi.GetPatientOnUidRequest")
	proto.RegisterType((*GetPatientOnUidReply)(nil), "patientinfoapi.GetPatientOnUidReply")
	proto.RegisterType((*Patient)(nil), "patientinfoapi.Patient")
	proto.RegisterType((*Error)(nil), "patientinfoapi.Error")
}

func init() { proto.RegisterFile("patients.proto", fileDescriptor_df0a59435314f582) }

var fileDescriptor_df0a59435314f582 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x4a, 0x2b, 0x31,
	0x14, 0xc6, 0x99, 0xb6, 0x73, 0x7b, 0x7b, 0x7a, 0xb9, 0x4a, 0xf0, 0xcf, 0x58, 0x45, 0xea, 0x20,
	0x22, 0x08, 0x2d, 0xd6, 0x17, 0x5f, 0x7d, 0x10, 0x11, 0xa4, 0x95, 0x81, 0x2e, 0x20, 0x75, 0x4e,
	0x6b, 0x20, 0x93, 0xc4, 0x49, 0x46, 0x28, 0xe2, 0x4b, 0x97, 0xa0, 0x0b, 0x70, 0x51, 0x6e, 0xc1,
	0x85, 0xc8, 0x64, 0x32, 0xb5, 0xad, 0x82, 0x6f, 0x39, 0xdf, 0xf9, 0x7d, 0xc9, 0x97, 0xe4, 0xc0,
	0x7f, 0x45, 0x0d, 0x43, 0x61, 0x74, 0x47, 0xa5, 0xd2, 0x48, 0x52, 0xd6, 0x4c, 0x8c, 0x25, 0x55,
	0xac, 0xb5, 0x37, 0x91, 0x72, 0xc2, 0xb1, 0x4b, 0x15, 0xeb, 0x52, 0x21, 0xa4, 0xa1, 0x86, 0x49,
	0xe1, 0xe8, 0xf0, 0x1c, 0xb6, 0xae, 0xd0, 0xdc, 0x16, 0x96, 0x81, 0x18, 0xb2, 0x38, 0xc2, 0x87,
	0x0c, 0xb5, 0x21, 0xfb, 0x00, 0x6e, 0xa7, 0x21, 0x8b, 0x03, 0xaf, 0xed, 0x1d, 0x37, 0xa2, 0x05,
	0x25, 0x7c, 0x84, 0x8d, 0x6f, 0x4e, 0xc5, 0xa7, 0xe4, 0x04, 0x7c, 0x4c, 0x53, 0x99, 0x5a, 0x4b,
	0xb3, 0xb7, 0xd9, 0x59, 0xce, 0xd3, 0xb9, 0xcc, 0x9b, 0x51, 0xc1, 0x90, 0x53, 0xa8, 0xbb, 0x76,
	0x50, 0xb1, 0xf8, 0xf6, 0x2a, 0xee, 0x0e, 0x88, 0x4a, 0x2e, 0x7c, 0xf3, 0xa0, 0xee, 0x44, 0xb2,
	0x0e, 0xd5, 0x6c, 0x1e, 0x2e, 0x5f, 0x92, 0x16, 0xfc, 0x1d, 0x67, 0x9c, 0xf7, 0x69, 0x82, 0x76,
	0xc7, 0x46, 0x34, 0xaf, 0x49, 0x1b, 0x9a, 0x34, 0x8e, 0x53, 0xd4, 0xfa, 0x86, 0x09, 0x0c, 0xaa,
	0xb6, 0xbd, 0x28, 0xe5, 0x84, 0xba, 0x97, 0x02, 0xfb, 0x59, 0x32, 0xc2, 0x34, 0xa8, 0x15, 0xc4,
	0x82, 0x44, 0x42, 0xf8, 0x87, 0x09, 0x65, 0xfc, 0xa2, 0x70, 0x05, 0xbe, 0x45, 0x96, 0xb4, 0x70,
	0x00, 0xbe, 0xbd, 0x24, 0x21, 0x50, 0xbb, 0x93, 0x31, 0xda, 0x7c, 0x7e, 0x64, 0xd7, 0x24, 0x80,
	0x7a, 0x82, 0x5a, 0xd3, 0x49, 0x99, 0xaf, 0x2c, 0xf3, 0x4e, 0x8c, 0x86, 0x32, 0xae, 0x5d, 0xb4,
	0xb2, 0xec, 0xbd, 0x78, 0xd0, 0x74, 0x57, 0xbe, 0x16, 0x63, 0x49, 0x66, 0x1e, 0xac, 0xad, 0xbc,
	0x3d, 0x39, 0x5a, 0x7d, 0xb8, 0x9f, 0xbf, 0xb5, 0x75, 0xf8, 0x2b, 0xa7, 0xf8, 0x34, 0x3c, 0x98,
	0xbd, 0x7f, 0xbc, 0x56, 0x76, 0xc3, 0x1d, 0x3b, 0x36, 0xce, 0xd1, 0x7d, 0xfa, 0xfa, 0xfe, 0xe7,
	0xd1, 0x1f, 0x3b, 0x40, 0x67, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x71, 0x63, 0x18, 0x80,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PatientInfoClient is the client API for PatientInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PatientInfoClient interface {
	GetPatientOnUid(ctx context.Context, in *GetPatientOnUidRequest, opts ...grpc.CallOption) (*GetPatientOnUidReply, error)
}

type patientInfoClient struct {
	cc *grpc.ClientConn
}

func NewPatientInfoClient(cc *grpc.ClientConn) PatientInfoClient {
	return &patientInfoClient{cc}
}

func (c *patientInfoClient) GetPatientOnUid(ctx context.Context, in *GetPatientOnUidRequest, opts ...grpc.CallOption) (*GetPatientOnUidReply, error) {
	out := new(GetPatientOnUidReply)
	err := c.cc.Invoke(ctx, "/patientinfoapi.PatientInfo/GetPatientOnUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientInfoServer is the server API for PatientInfo service.
type PatientInfoServer interface {
	GetPatientOnUid(context.Context, *GetPatientOnUidRequest) (*GetPatientOnUidReply, error)
}

// UnimplementedPatientInfoServer can be embedded to have forward compatible implementations.
type UnimplementedPatientInfoServer struct {
}

func (*UnimplementedPatientInfoServer) GetPatientOnUid(ctx context.Context, req *GetPatientOnUidRequest) (*GetPatientOnUidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatientOnUid not implemented")
}

func RegisterPatientInfoServer(s *grpc.Server, srv PatientInfoServer) {
	s.RegisterService(&_PatientInfo_serviceDesc, srv)
}

func _PatientInfo_GetPatientOnUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatientOnUidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientInfoServer).GetPatientOnUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/patientinfoapi.PatientInfo/GetPatientOnUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientInfoServer).GetPatientOnUid(ctx, req.(*GetPatientOnUidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PatientInfo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "patientinfoapi.PatientInfo",
	HandlerType: (*PatientInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPatientOnUid",
			Handler:    _PatientInfo_GetPatientOnUid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "patients.proto",
}
