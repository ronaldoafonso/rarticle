// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rarticlepb/rarticle.proto

package rarticlepb

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

type RArticleRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RArticleRequest) Reset()         { *m = RArticleRequest{} }
func (m *RArticleRequest) String() string { return proto.CompactTextString(m) }
func (*RArticleRequest) ProtoMessage()    {}
func (*RArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff80b058ea99c105, []int{0}
}

func (m *RArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RArticleRequest.Unmarshal(m, b)
}
func (m *RArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RArticleRequest.Marshal(b, m, deterministic)
}
func (m *RArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RArticleRequest.Merge(m, src)
}
func (m *RArticleRequest) XXX_Size() int {
	return xxx_messageInfo_RArticleRequest.Size(m)
}
func (m *RArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RArticleRequest proto.InternalMessageInfo

func (m *RArticleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RArticleRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type RArticleResponse struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RArticleResponse) Reset()         { *m = RArticleResponse{} }
func (m *RArticleResponse) String() string { return proto.CompactTextString(m) }
func (*RArticleResponse) ProtoMessage()    {}
func (*RArticleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff80b058ea99c105, []int{1}
}

func (m *RArticleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RArticleResponse.Unmarshal(m, b)
}
func (m *RArticleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RArticleResponse.Marshal(b, m, deterministic)
}
func (m *RArticleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RArticleResponse.Merge(m, src)
}
func (m *RArticleResponse) XXX_Size() int {
	return xxx_messageInfo_RArticleResponse.Size(m)
}
func (m *RArticleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RArticleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RArticleResponse proto.InternalMessageInfo

func (m *RArticleResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SearchRequest struct {
	Word                 string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff80b058ea99c105, []int{2}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *SearchRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type SearchResponse struct {
	Link                 string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff80b058ea99c105, []int{3}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func init() {
	proto.RegisterType((*RArticleRequest)(nil), "rarticle.RArticleRequest")
	proto.RegisterType((*RArticleResponse)(nil), "rarticle.RArticleResponse")
	proto.RegisterType((*SearchRequest)(nil), "rarticle.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "rarticle.SearchResponse")
}

func init() {
	proto.RegisterFile("rarticlepb/rarticle.proto", fileDescriptor_ff80b058ea99c105)
}

var fileDescriptor_ff80b058ea99c105 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x4a, 0x2c, 0x2a,
	0xc9, 0x4c, 0xce, 0x49, 0x2d, 0x48, 0xd2, 0x87, 0x31, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x38, 0x60, 0x7c, 0x25, 0x4b, 0x2e, 0xfe, 0x20, 0x47, 0x08, 0x3b, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xcc, 0x06, 0x89, 0xe5, 0x24, 0xe6, 0xa5, 0x4b, 0x30, 0x41, 0xc4, 0x40, 0x6c, 0x25, 0x1d,
	0x2e, 0x01, 0x84, 0xd6, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xfc,
	0xbc, 0x92, 0xd4, 0xbc, 0x12, 0xa8, 0x76, 0x18, 0x57, 0xc9, 0x9c, 0x8b, 0x37, 0x38, 0x35, 0xb1,
	0x28, 0x39, 0x03, 0xc9, 0x9a, 0xf2, 0xfc, 0xa2, 0x14, 0x98, 0x35, 0x20, 0x36, 0x56, 0x6b, 0x54,
	0xb8, 0xf8, 0x60, 0x1a, 0xa1, 0x96, 0x80, 0x54, 0x65, 0xe6, 0x65, 0xc3, 0x74, 0x82, 0xd8, 0x46,
	0xb3, 0x19, 0xb9, 0x38, 0x60, 0xae, 0x11, 0x72, 0xe3, 0xe2, 0x4e, 0x4f, 0x2d, 0x81, 0x73, 0x25,
	0xf5, 0xe0, 0xde, 0x47, 0xf3, 0xab, 0x94, 0x14, 0x36, 0x29, 0x88, 0x35, 0x4a, 0x0c, 0x42, 0xee,
	0x5c, 0x7c, 0xc5, 0x10, 0xab, 0x61, 0x46, 0x89, 0x23, 0xd4, 0xa3, 0xf8, 0x46, 0x4a, 0x02, 0x53,
	0x02, 0x66, 0x8c, 0x01, 0xa3, 0x13, 0x4f, 0x14, 0x17, 0x22, 0x32, 0x92, 0xd8, 0xc0, 0x91, 0x60,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x94, 0x12, 0xc1, 0x52, 0xa1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RArticleClient is the client API for RArticle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RArticleClient interface {
	GetRArticle(ctx context.Context, in *RArticleRequest, opts ...grpc.CallOption) (*RArticleResponse, error)
	SearchRArticle(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (RArticle_SearchRArticleClient, error)
}

type rArticleClient struct {
	cc grpc.ClientConnInterface
}

func NewRArticleClient(cc grpc.ClientConnInterface) RArticleClient {
	return &rArticleClient{cc}
}

func (c *rArticleClient) GetRArticle(ctx context.Context, in *RArticleRequest, opts ...grpc.CallOption) (*RArticleResponse, error) {
	out := new(RArticleResponse)
	err := c.cc.Invoke(ctx, "/rarticle.RArticle/getRArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rArticleClient) SearchRArticle(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (RArticle_SearchRArticleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RArticle_serviceDesc.Streams[0], "/rarticle.RArticle/searchRArticle", opts...)
	if err != nil {
		return nil, err
	}
	x := &rArticleSearchRArticleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RArticle_SearchRArticleClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type rArticleSearchRArticleClient struct {
	grpc.ClientStream
}

func (x *rArticleSearchRArticleClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RArticleServer is the server API for RArticle service.
type RArticleServer interface {
	GetRArticle(context.Context, *RArticleRequest) (*RArticleResponse, error)
	SearchRArticle(*SearchRequest, RArticle_SearchRArticleServer) error
}

// UnimplementedRArticleServer can be embedded to have forward compatible implementations.
type UnimplementedRArticleServer struct {
}

func (*UnimplementedRArticleServer) GetRArticle(ctx context.Context, req *RArticleRequest) (*RArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRArticle not implemented")
}
func (*UnimplementedRArticleServer) SearchRArticle(req *SearchRequest, srv RArticle_SearchRArticleServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchRArticle not implemented")
}

func RegisterRArticleServer(s *grpc.Server, srv RArticleServer) {
	s.RegisterService(&_RArticle_serviceDesc, srv)
}

func _RArticle_GetRArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RArticleServer).GetRArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rarticle.RArticle/GetRArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RArticleServer).GetRArticle(ctx, req.(*RArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RArticle_SearchRArticle_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RArticleServer).SearchRArticle(m, &rArticleSearchRArticleServer{stream})
}

type RArticle_SearchRArticleServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type rArticleSearchRArticleServer struct {
	grpc.ServerStream
}

func (x *rArticleSearchRArticleServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RArticle_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rarticle.RArticle",
	HandlerType: (*RArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getRArticle",
			Handler:    _RArticle_GetRArticle_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchRArticle",
			Handler:       _RArticle_SearchRArticle_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rarticlepb/rarticle.proto",
}
