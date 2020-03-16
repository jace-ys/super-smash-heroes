// Code generated by protoc-gen-go. DO NOT EDIT.
// source: battle.proto

package battle

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
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xad, 0xd0, 0xba, 0x5e, 0xa7, 0x68, 0x10, 0x89, 0x7b, 0x90, 0x91, 0xa7, 0x81, 0xb8,
	0x87, 0xf9, 0x0f, 0xf6, 0xe2, 0xa3, 0x92, 0xf9, 0x2e, 0x6d, 0x77, 0x99, 0x81, 0x98, 0xc4, 0xe4,
	0xd6, 0xb2, 0xff, 0xe4, 0x8f, 0x94, 0x25, 0x75, 0x54, 0x91, 0x3d, 0x7e, 0xf7, 0x9c, 0xd3, 0xd3,
	0x43, 0x60, 0x5c, 0x57, 0x44, 0x1a, 0xe7, 0xce, 0x5b, 0xb2, 0xac, 0x48, 0x24, 0xbe, 0x32, 0x28,
	0x9e, 0x75, 0xb5, 0x45, 0xcf, 0x04, 0x8c, 0x95, 0x21, 0xd4, 0x5a, 0x6d, 0xd0, 0x34, 0xc8, 0xf3,
	0x69, 0x36, 0xcb, 0xe5, 0xaf, 0x1b, 0x9b, 0xc0, 0x28, 0x90, 0x47, 0xb3, 0xa1, 0x37, 0x5e, 0x44,
	0x7d, 0xcf, 0xec, 0x0a, 0xf2, 0xe0, 0x10, 0xd7, 0xfc, 0x24, 0x0a, 0x09, 0xd8, 0x2d, 0xc0, 0xba,
	0xf5, 0x55, 0xad, 0xb4, 0xa2, 0x2d, 0x1f, 0x45, 0x69, 0x70, 0xd9, 0xa5, 0x9c, 0xed, 0xd0, 0xf3,
	0x32, 0xa5, 0x22, 0xb0, 0x6b, 0x28, 0x1a, 0xfb, 0x5e, 0x57, 0xc4, 0x21, 0x9e, 0x7b, 0x12, 0x0e,
	0x2e, 0x1e, 0x91, 0x24, 0x86, 0x56, 0x93, 0xc4, 0x8f, 0x16, 0x03, 0xb1, 0x7b, 0x00, 0x17, 0x17,
	0xbc, 0x5a, 0x83, 0x3c, 0x9b, 0x66, 0xb3, 0xd3, 0xc5, 0xf9, 0xbc, 0x5f, 0x9b, 0xb6, 0xc9, 0x32,
	0x39, 0x9e, 0x0c, 0x0e, 0xec, 0xd4, 0x59, 0x7e, 0x7c, 0xc8, 0xfe, 0xd2, 0x59, 0x71, 0x07, 0x97,
	0x83, 0xc6, 0xe0, 0xac, 0x09, 0xb8, 0xfb, 0xbd, 0x4e, 0x19, 0x83, 0x3e, 0xd6, 0xe5, 0xb2, 0xa7,
	0xc5, 0x0a, 0xce, 0x96, 0xf1, 0x43, 0x2b, 0xf4, 0x9f, 0xaa, 0x41, 0xb6, 0x84, 0x72, 0x9f, 0x66,
	0xfc, 0xa7, 0xe5, 0xef, 0x84, 0xc9, 0xcd, 0x3f, 0x4a, 0xaa, 0x12, 0x47, 0x75, 0x11, 0x5f, 0xec,
	0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x2e, 0x92, 0xc1, 0xc1, 0x01, 0x00, 0x00,
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
