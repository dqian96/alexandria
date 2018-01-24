// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	Entry
	AppendEntriesRequest
	AppendEntriesReply
	RequestVoteRequest
	RequestVoteReply
*/
package director

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

type Entry struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AppendEntriesRequest struct {
	LeaderId    string   `protobuf:"bytes,1,opt,name=leaderId" json:"leaderId,omitempty"`
	CommitIndex uint64   `protobuf:"varint,2,opt,name=commitIndex" json:"commitIndex,omitempty"`
	Term        uint64   `protobuf:"varint,3,opt,name=term" json:"term,omitempty"`
	Entries     []*Entry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	LastEntry   *Entry   `protobuf:"bytes,5,opt,name=lastEntry" json:"lastEntry,omitempty"`
	LastIndex   uint64   `protobuf:"varint,6,opt,name=lastIndex" json:"lastIndex,omitempty"`
	LastTerm    uint64   `protobuf:"varint,7,opt,name=lastTerm" json:"lastTerm,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AppendEntriesRequest) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

func (m *AppendEntriesRequest) GetCommitIndex() uint64 {
	if m != nil {
		return m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AppendEntriesRequest) GetLastEntry() *Entry {
	if m != nil {
		return m.LastEntry
	}
	return nil
}

func (m *AppendEntriesRequest) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetLastTerm() uint64 {
	if m != nil {
		return m.LastTerm
	}
	return 0
}

type AppendEntriesReply struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Term    uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
}

func (m *AppendEntriesReply) Reset()                    { *m = AppendEntriesReply{} }
func (m *AppendEntriesReply) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesReply) ProtoMessage()               {}
func (*AppendEntriesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppendEntriesReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AppendEntriesReply) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type RequestVoteRequest struct {
	CandidateId string `protobuf:"bytes,1,opt,name=candidateId" json:"candidateId,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	LastTerm    uint64 `protobuf:"varint,3,opt,name=lastTerm" json:"lastTerm,omitempty"`
	LastIndex   uint64 `protobuf:"varint,4,opt,name=lastIndex" json:"lastIndex,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestVoteRequest) GetCandidateId() string {
	if m != nil {
		return m.CandidateId
	}
	return ""
}

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetLastTerm() uint64 {
	if m != nil {
		return m.LastTerm
	}
	return 0
}

func (m *RequestVoteRequest) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

type RequestVoteReply struct {
	VoteGranted bool   `protobuf:"varint,1,opt,name=voteGranted" json:"voteGranted,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
}

func (m *RequestVoteReply) Reset()                    { *m = RequestVoteReply{} }
func (m *RequestVoteReply) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteReply) ProtoMessage()               {}
func (*RequestVoteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestVoteReply) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

func (m *RequestVoteReply) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func init() {
	proto.RegisterType((*Entry)(nil), "server.Entry")
	proto.RegisterType((*AppendEntriesRequest)(nil), "server.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesReply)(nil), "server.AppendEntriesReply")
	proto.RegisterType((*RequestVoteRequest)(nil), "server.RequestVoteRequest")
	proto.RegisterType((*RequestVoteReply)(nil), "server.RequestVoteReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Director service

type DirectorClient interface {
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesReply, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteReply, error)
}

type directorClient struct {
	cc *grpc.ClientConn
}

func NewDirectorClient(cc *grpc.ClientConn) DirectorClient {
	return &directorClient{cc}
}

func (c *directorClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesReply, error) {
	out := new(AppendEntriesReply)
	err := grpc.Invoke(ctx, "/server.Director/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteReply, error) {
	out := new(RequestVoteReply)
	err := grpc.Invoke(ctx, "/server.Director/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Director service

type DirectorServer interface {
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesReply, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteReply, error)
}

func RegisterDirectorServer(s *grpc.Server, srv DirectorServer) {
	s.RegisterService(&_Director_serviceDesc, srv)
}

func _Director_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Director/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Director/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Director_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Director",
	HandlerType: (*DirectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendEntries",
			Handler:    _Director_AppendEntries_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _Director_RequestVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcb, 0x8e, 0xd3, 0x40,
	0x10, 0xc4, 0x6b, 0xe7, 0xb1, 0x6d, 0x56, 0x5a, 0x35, 0x7b, 0xb0, 0xac, 0x3d, 0x58, 0xbe, 0xb0,
	0x12, 0xc2, 0x48, 0x8b, 0x84, 0xc4, 0x91, 0x88, 0x00, 0x11, 0x97, 0xc8, 0x20, 0xee, 0x13, 0x4f,
	0x0b, 0x2c, 0xfc, 0xca, 0xcc, 0x38, 0x8a, 0xef, 0x7c, 0x07, 0xdf, 0xc9, 0x11, 0x79, 0xfc, 0x88,
	0x9d, 0x78, 0x6f, 0x53, 0x55, 0x9d, 0x9e, 0xaa, 0x8a, 0x07, 0x9e, 0x4b, 0x12, 0x07, 0x12, 0x41,
	0x21, 0x72, 0x95, 0xe3, 0xbc, 0x41, 0xfe, 0x1b, 0x98, 0xad, 0x33, 0x25, 0x2a, 0xbc, 0x05, 0xf3,
	0x37, 0x55, 0x8e, 0xe1, 0x19, 0x0f, 0xd7, 0x61, 0x7d, 0xc4, 0x3b, 0x98, 0x1d, 0x58, 0x52, 0x92,
	0x73, 0xa5, 0xb9, 0x06, 0xf8, 0xff, 0x0c, 0xb8, 0xfb, 0x50, 0x14, 0x94, 0xf1, 0xfa, 0x77, 0x31,
	0xc9, 0x90, 0xf6, 0x25, 0x49, 0x85, 0x2e, 0x2c, 0x13, 0x62, 0x9c, 0xc4, 0x86, 0xb7, 0x5b, 0x7a,
	0x8c, 0x1e, 0xd8, 0x51, 0x9e, 0xa6, 0xb1, 0xda, 0x64, 0x9c, 0x8e, 0x7a, 0xa1, 0x15, 0x0e, 0x29,
	0x44, 0xb0, 0x14, 0x89, 0xd4, 0x31, 0xb5, 0xa4, 0xcf, 0xf8, 0x12, 0x16, 0xd4, 0xdc, 0xe1, 0x58,
	0x9e, 0xf9, 0x60, 0x3f, 0xde, 0x04, 0x6d, 0x06, 0x6d, 0x39, 0xec, 0x54, 0x7c, 0x05, 0xd7, 0x09,
	0x93, 0x4a, 0xb3, 0xce, 0xcc, 0x33, 0x2e, 0x47, 0x4f, 0x3a, 0xde, 0x37, 0xc3, 0x8d, 0x93, 0xb9,
	0xbe, 0xee, 0x44, 0xe8, 0x14, 0x4c, 0xaa, 0xef, 0xb5, 0x97, 0x85, 0x16, 0x7b, 0xec, 0xaf, 0x00,
	0xcf, 0x92, 0x17, 0x49, 0x85, 0x0e, 0x2c, 0x64, 0x19, 0x45, 0x24, 0xa5, 0x8e, 0xbd, 0x0c, 0x3b,
	0xd8, 0x67, 0xba, 0x3a, 0x65, 0xf2, 0xff, 0x18, 0x80, 0x6d, 0x63, 0x3f, 0x72, 0x45, 0x5d, 0x79,
	0x75, 0x41, 0x2c, 0xe3, 0x31, 0x67, 0x8a, 0xfa, 0xfe, 0x86, 0xd4, 0xd4, 0xb2, 0x91, 0x59, 0x73,
	0x6c, 0x76, 0x1c, 0xd3, 0x3a, 0x8b, 0xe9, 0x7f, 0x81, 0xdb, 0x91, 0x8b, 0x3a, 0x88, 0x07, 0xf6,
	0x21, 0x57, 0xf4, 0x59, 0xb0, 0x4c, 0x11, 0x6f, 0xc3, 0x0c, 0xa9, 0x29, 0x0f, 0x8f, 0x7f, 0x0d,
	0x58, 0x7e, 0x8c, 0x05, 0x45, 0x2a, 0x17, 0xf8, 0x15, 0x6e, 0x46, 0x0d, 0xe1, 0x7d, 0xf7, 0x37,
	0x4c, 0x7d, 0x32, 0xae, 0xfb, 0x84, 0x5a, 0x24, 0x95, 0xff, 0x0c, 0xd7, 0x60, 0x0f, 0x3c, 0x62,
	0x3f, 0x7c, 0x59, 0x9f, 0xeb, 0x4c, 0x6a, 0x7a, 0xcd, 0xea, 0x13, 0xbc, 0x8e, 0xf2, 0x34, 0xf8,
	0x19, 0xab, 0x5f, 0xe5, 0x2e, 0xe0, 0xfb, 0x98, 0x65, 0xef, 0xdf, 0x05, 0x2c, 0xa1, 0x23, 0xcb,
	0xb8, 0x88, 0x59, 0xc0, 0x5b, 0xf7, 0xed, 0x8e, 0xd5, 0x8b, 0x2e, 0xce, 0x37, 0x8d, 0xb7, 0xf5,
	0x7b, 0xd9, 0x1a, 0xbb, 0xb9, 0x7e, 0x38, 0x6f, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x52,
	0x74, 0xaf, 0x48, 0x03, 0x00, 0x00,
}
