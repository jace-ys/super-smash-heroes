// Code generated by protoc-gen-go. DO NOT EDIT.
// source: battle.proto

package battle

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Player struct {
	Intelligence         int32    `protobuf:"varint,5,opt,name=intelligence,proto3" json:"intelligence,omitempty"`
	Strength             int32    `protobuf:"varint,6,opt,name=strength,proto3" json:"strength,omitempty"`
	Speed                int32    `protobuf:"varint,7,opt,name=speed,proto3" json:"speed,omitempty"`
	Durability           int32    `protobuf:"varint,8,opt,name=durability,proto3" json:"durability,omitempty"`
	Power                int32    `protobuf:"varint,9,opt,name=power,proto3" json:"power,omitempty"`
	Combat               int32    `protobuf:"varint,10,opt,name=combat,proto3" json:"combat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetIntelligence() int32 {
	if m != nil {
		return m.Intelligence
	}
	return 0
}

func (m *Player) GetStrength() int32 {
	if m != nil {
		return m.Strength
	}
	return 0
}

func (m *Player) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *Player) GetDurability() int32 {
	if m != nil {
		return m.Durability
	}
	return 0
}

func (m *Player) GetPower() int32 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *Player) GetCombat() int32 {
	if m != nil {
		return m.Combat
	}
	return 0
}

type GetResultRequest struct {
	PlayerOne            *Player  `protobuf:"bytes,1,opt,name=player_one,json=playerOne,proto3" json:"player_one,omitempty"`
	PlayerTwo            *Player  `protobuf:"bytes,2,opt,name=player_two,json=playerTwo,proto3" json:"player_two,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResultRequest) Reset()         { *m = GetResultRequest{} }
func (m *GetResultRequest) String() string { return proto.CompactTextString(m) }
func (*GetResultRequest) ProtoMessage()    {}
func (*GetResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{1}
}

func (m *GetResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResultRequest.Unmarshal(m, b)
}
func (m *GetResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResultRequest.Marshal(b, m, deterministic)
}
func (m *GetResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResultRequest.Merge(m, src)
}
func (m *GetResultRequest) XXX_Size() int {
	return xxx_messageInfo_GetResultRequest.Size(m)
}
func (m *GetResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResultRequest proto.InternalMessageInfo

func (m *GetResultRequest) GetPlayerOne() *Player {
	if m != nil {
		return m.PlayerOne
	}
	return nil
}

func (m *GetResultRequest) GetPlayerTwo() *Player {
	if m != nil {
		return m.PlayerTwo
	}
	return nil
}

type GetResultResponse struct {
	Winner               int32    `protobuf:"varint,1,opt,name=winner,proto3" json:"winner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResultResponse) Reset()         { *m = GetResultResponse{} }
func (m *GetResultResponse) String() string { return proto.CompactTextString(m) }
func (*GetResultResponse) ProtoMessage()    {}
func (*GetResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{2}
}

func (m *GetResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResultResponse.Unmarshal(m, b)
}
func (m *GetResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResultResponse.Marshal(b, m, deterministic)
}
func (m *GetResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResultResponse.Merge(m, src)
}
func (m *GetResultResponse) XXX_Size() int {
	return xxx_messageInfo_GetResultResponse.Size(m)
}
func (m *GetResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResultResponse proto.InternalMessageInfo

func (m *GetResultResponse) GetWinner() int32 {
	if m != nil {
		return m.Winner
	}
	return 0
}

func init() {
	proto.RegisterType((*Player)(nil), "battle.Player")
	proto.RegisterType((*GetResultRequest)(nil), "battle.GetResultRequest")
	proto.RegisterType((*GetResultResponse)(nil), "battle.GetResultResponse")
}

func init() {
	proto.RegisterFile("battle.proto", fileDescriptor_88d4c72b5869c382)
}

var fileDescriptor_88d4c72b5869c382 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x86, 0x99, 0xc2, 0xcc, 0x6d, 0xcf, 0xed, 0xbd, 0x6a, 0x50, 0x89, 0x45, 0xa4, 0xcc, 0xaa,
	0x28, 0x76, 0xb0, 0xee, 0x5c, 0xba, 0x71, 0xa9, 0x8c, 0x82, 0x4b, 0xc9, 0xb4, 0x87, 0x69, 0x20,
	0x26, 0x31, 0x39, 0xed, 0xd0, 0xad, 0xaf, 0xe0, 0x6b, 0xf8, 0x36, 0xbe, 0x82, 0x0f, 0x22, 0x4d,
	0x6a, 0xa9, 0x22, 0x2e, 0xbf, 0xf3, 0xff, 0x27, 0x7f, 0xfe, 0x04, 0xba, 0x95, 0x20, 0x52, 0x38,
	0xb4, 0xce, 0x90, 0x61, 0x59, 0xa4, 0xde, 0x61, 0x6d, 0x4c, 0xad, 0xb0, 0x10, 0x56, 0x16, 0x42,
	0x6b, 0x43, 0x82, 0xa4, 0xd1, 0x3e, 0xba, 0xf2, 0xd7, 0x04, 0xb2, 0x1b, 0x25, 0x16, 0xe8, 0x58,
	0x0e, 0x5d, 0xa9, 0x09, 0x95, 0x92, 0x35, 0xea, 0x31, 0xf2, 0xb4, 0x9f, 0x0c, 0xd2, 0xf2, 0xcb,
	0x8c, 0xf5, 0xa0, 0xed, 0xc9, 0xa1, 0xae, 0x69, 0xca, 0xb3, 0xa0, 0xaf, 0x99, 0xed, 0x42, 0xea,
	0x2d, 0xe2, 0x84, 0xff, 0x09, 0x42, 0x04, 0x76, 0x04, 0x30, 0x99, 0x39, 0x51, 0x49, 0x25, 0x69,
	0xc1, 0xdb, 0x41, 0xda, 0x98, 0x2c, 0xb7, 0xac, 0x69, 0xd0, 0xf1, 0x4e, 0xdc, 0x0a, 0xc0, 0xf6,
	0x21, 0x1b, 0x9b, 0xc7, 0x4a, 0x10, 0x87, 0x30, 0x5e, 0x51, 0x6e, 0x61, 0xfb, 0x0a, 0xa9, 0x44,
	0x3f, 0x53, 0x54, 0xe2, 0xd3, 0x0c, 0x3d, 0xb1, 0x53, 0x00, 0x1b, 0x1a, 0x3c, 0x18, 0x8d, 0x3c,
	0xe9, 0x27, 0x83, 0xbf, 0xa3, 0xff, 0xc3, 0xd5, 0x5b, 0xc4, 0x6e, 0x65, 0x27, 0x3a, 0xae, 0x35,
	0x6e, 0xd8, 0xa9, 0x31, 0xbc, 0xf5, 0x9b, 0xfd, 0xae, 0x31, 0xf9, 0x09, 0xec, 0x6c, 0x24, 0x7a,
	0x6b, 0xb4, 0xc7, 0xe5, 0xf5, 0x1a, 0xa9, 0x35, 0xba, 0x10, 0x97, 0x96, 0x2b, 0x1a, 0x4d, 0xe1,
	0xdf, 0x65, 0x38, 0xe8, 0x16, 0xdd, 0x5c, 0x8e, 0x91, 0xdd, 0x43, 0x67, 0xbd, 0xcd, 0xf8, 0x67,
	0xca, 0xf7, 0x0a, 0xbd, 0x83, 0x1f, 0x94, 0x18, 0x95, 0xef, 0x3d, 0xbf, 0xbd, 0xbf, 0xb4, 0xb6,
	0x72, 0x28, 0xe6, 0x67, 0x45, 0x74, 0x5d, 0x24, 0xc7, 0x55, 0x16, 0xbe, 0xef, 0xfc, 0x23, 0x00,
	0x00, 0xff, 0xff, 0x03, 0x4e, 0x56, 0xb7, 0xf4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BattleServiceClient is the client API for BattleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BattleServiceClient interface {
	GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error)
}

type battleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBattleServiceClient(cc grpc.ClientConnInterface) BattleServiceClient {
	return &battleServiceClient{cc}
}

func (c *battleServiceClient) GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*GetResultResponse, error) {
	out := new(GetResultResponse)
	err := c.cc.Invoke(ctx, "/battle.BattleService/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BattleServiceServer is the server API for BattleService service.
type BattleServiceServer interface {
	GetResult(context.Context, *GetResultRequest) (*GetResultResponse, error)
}

// UnimplementedBattleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBattleServiceServer struct {
}

func (*UnimplementedBattleServiceServer) GetResult(ctx context.Context, req *GetResultRequest) (*GetResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}

func RegisterBattleServiceServer(s *grpc.Server, srv BattleServiceServer) {
	s.RegisterService(&_BattleService_serviceDesc, srv)
}

func _BattleService_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleServiceServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/battle.BattleService/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleServiceServer).GetResult(ctx, req.(*GetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BattleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "battle.BattleService",
	HandlerType: (*BattleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResult",
			Handler:    _BattleService_GetResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "battle.proto",
}