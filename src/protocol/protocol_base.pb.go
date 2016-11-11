// Code generated by protoc-gen-go.
// source: protocol_base.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol_base.proto
	protocol_login.proto

It has these top-level messages:
	C2S_SystemTick
	S2C_GateState
	PlayerInfo
	Attribute
	Reward
	C2S_Login
	S2C_Login
	C2S_RandName
	S2C_RandName
	C2S_CreateRole
	S2C_CreateRole
	C2S_LoadRoleInfo
	S2C_LoadRoleInfo
	S2C_ServerTime
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtocolType int32

const (
	ProtocolType_S_2_C_KEEP_ALIVE_ACK ProtocolType = 4
	ProtocolType_S_2_C_GATE_STATE     ProtocolType = 100
	// login
	ProtocolType_C_2_S_LOGIN_BASE ProtocolType = 101
	ProtocolType_C_2_S_LOGIN_TOP  ProtocolType = 200
	// role
	ProtocolType_C_2_S_ROLE_BASE ProtocolType = 201
	ProtocolType_C_2_S_ROLE_TOP  ProtocolType = 250
	// gm or test
	ProtocolType_C_2_S_TEST_BASE ProtocolType = 251
	ProtocolType_C_2_S_TEST_TOP  ProtocolType = 300
	// hero
	ProtocolType_C_2_S_HERO_BASE ProtocolType = 301
	ProtocolType_C_2_S_HERO_TOP  ProtocolType = 400
	// item
	ProtocolType_C_2_S_ITEM_BASE ProtocolType = 401
	ProtocolType_C_2_S_ITEM_TOP  ProtocolType = 500
	// material
	ProtocolType_C_2_S_MATERIAL_BASE ProtocolType = 501
	ProtocolType_C_2_S_MATERIAL_TOP  ProtocolType = 600
	// formation
	ProtocolType_C_2_S_FORMATION_BASE ProtocolType = 601
	ProtocolType_C_2_S_FORMATION_TOP  ProtocolType = 700
	// tactic
	ProtocolType_C_2_S_TACTIC_BASE ProtocolType = 701
	ProtocolType_C_2_S_TACTIC_TOP  ProtocolType = 800
	// mail
	ProtocolType_C_2_S_MAIL_BASE ProtocolType = 801
	ProtocolType_C_2_S_MAIL_TOP  ProtocolType = 900
	// instance
	ProtocolType_C_2_S_INSTANCE_BASE ProtocolType = 901
	ProtocolType_C_2_S_INSTANCE_TOP  ProtocolType = 1000
	// rank
	ProtocolType_C_2_S_RANK_BASE ProtocolType = 1001
	ProtocolType_C_2_S_RANK_TOP  ProtocolType = 1100
	// tower
	ProtocolType_C_2_S_TOWER_BASE ProtocolType = 1101
	ProtocolType_C_2_S_TOWER_TOP  ProtocolType = 1200
	// shop
	ProtocolType_C_2_S_SHOP_BASE ProtocolType = 1201
	ProtocolType_C_2_S_SHOP_TOP  ProtocolType = 1300
	// task
	ProtocolType_C_2_S_TASK_BASE ProtocolType = 1301
	ProtocolType_C_2_S_TASK_TOP  ProtocolType = 1400
	// chat
	ProtocolType_C_2_S_CHAT_BASE ProtocolType = 1401
	ProtocolType_C_2_S_CHAT_TOP  ProtocolType = 1500
	// buddy
	ProtocolType_C_2_S_BUDDY_BASE ProtocolType = 1501
	ProtocolType_C_2_S_BUDDY_TOP  ProtocolType = 1600
	// clone
	ProtocolType_C_2_S_CLONE_BASE ProtocolType = 1601
	ProtocolType_C_2_S_CLONE_TOP  ProtocolType = 1700
	// standard
	ProtocolType_C_2_S_STANDARD_BASE ProtocolType = 1701
	ProtocolType_C_2_S_STANDARD_TOP  ProtocolType = 1800
	// arena
	ProtocolType_C_2_S_ARENA_BASE ProtocolType = 1801
	ProtocolType_C_2_S_ARENA_TOP  ProtocolType = 1900
	// crop
	ProtocolType_C_2_S_CROP_BASE ProtocolType = 1901
	ProtocolType_C_2_S_CROP_TOP  ProtocolType = 2000
	// champion race
	ProtocolType_C_2_S_CHAMPION_RACE_BASE ProtocolType = 2001
	ProtocolType_C_2_S_CHAMPION_RACE_TOP  ProtocolType = 2100
	// crop war
	ProtocolType_C_2_S_CROP_WAR_BASE ProtocolType = 2101
	ProtocolType_C_2_S_CROP_WAR_TOP  ProtocolType = 2200
	// main
	ProtocolType_C_2_S_MAIN_BASE ProtocolType = 2201
	ProtocolType_C_2_S_MAIN_TOP  ProtocolType = 2300
	// reward activity
	ProtocolType_C_2_S_REWARD_ACTIVITY_BASE ProtocolType = 2301
	ProtocolType_C_2_S_REWARD_ACTIVITY_TOP  ProtocolType = 2400
)

var ProtocolType_name = map[int32]string{
	4:    "S_2_C_KEEP_ALIVE_ACK",
	100:  "S_2_C_GATE_STATE",
	101:  "C_2_S_LOGIN_BASE",
	200:  "C_2_S_LOGIN_TOP",
	201:  "C_2_S_ROLE_BASE",
	250:  "C_2_S_ROLE_TOP",
	251:  "C_2_S_TEST_BASE",
	300:  "C_2_S_TEST_TOP",
	301:  "C_2_S_HERO_BASE",
	400:  "C_2_S_HERO_TOP",
	401:  "C_2_S_ITEM_BASE",
	500:  "C_2_S_ITEM_TOP",
	501:  "C_2_S_MATERIAL_BASE",
	600:  "C_2_S_MATERIAL_TOP",
	601:  "C_2_S_FORMATION_BASE",
	700:  "C_2_S_FORMATION_TOP",
	701:  "C_2_S_TACTIC_BASE",
	800:  "C_2_S_TACTIC_TOP",
	801:  "C_2_S_MAIL_BASE",
	900:  "C_2_S_MAIL_TOP",
	901:  "C_2_S_INSTANCE_BASE",
	1000: "C_2_S_INSTANCE_TOP",
	1001: "C_2_S_RANK_BASE",
	1100: "C_2_S_RANK_TOP",
	1101: "C_2_S_TOWER_BASE",
	1200: "C_2_S_TOWER_TOP",
	1201: "C_2_S_SHOP_BASE",
	1300: "C_2_S_SHOP_TOP",
	1301: "C_2_S_TASK_BASE",
	1400: "C_2_S_TASK_TOP",
	1401: "C_2_S_CHAT_BASE",
	1500: "C_2_S_CHAT_TOP",
	1501: "C_2_S_BUDDY_BASE",
	1600: "C_2_S_BUDDY_TOP",
	1601: "C_2_S_CLONE_BASE",
	1700: "C_2_S_CLONE_TOP",
	1701: "C_2_S_STANDARD_BASE",
	1800: "C_2_S_STANDARD_TOP",
	1801: "C_2_S_ARENA_BASE",
	1900: "C_2_S_ARENA_TOP",
	1901: "C_2_S_CROP_BASE",
	2000: "C_2_S_CROP_TOP",
	2001: "C_2_S_CHAMPION_RACE_BASE",
	2100: "C_2_S_CHAMPION_RACE_TOP",
	2101: "C_2_S_CROP_WAR_BASE",
	2200: "C_2_S_CROP_WAR_TOP",
	2201: "C_2_S_MAIN_BASE",
	2300: "C_2_S_MAIN_TOP",
	2301: "C_2_S_REWARD_ACTIVITY_BASE",
	2400: "C_2_S_REWARD_ACTIVITY_TOP",
}
var ProtocolType_value = map[string]int32{
	"S_2_C_KEEP_ALIVE_ACK":       4,
	"S_2_C_GATE_STATE":           100,
	"C_2_S_LOGIN_BASE":           101,
	"C_2_S_LOGIN_TOP":            200,
	"C_2_S_ROLE_BASE":            201,
	"C_2_S_ROLE_TOP":             250,
	"C_2_S_TEST_BASE":            251,
	"C_2_S_TEST_TOP":             300,
	"C_2_S_HERO_BASE":            301,
	"C_2_S_HERO_TOP":             400,
	"C_2_S_ITEM_BASE":            401,
	"C_2_S_ITEM_TOP":             500,
	"C_2_S_MATERIAL_BASE":        501,
	"C_2_S_MATERIAL_TOP":         600,
	"C_2_S_FORMATION_BASE":       601,
	"C_2_S_FORMATION_TOP":        700,
	"C_2_S_TACTIC_BASE":          701,
	"C_2_S_TACTIC_TOP":           800,
	"C_2_S_MAIL_BASE":            801,
	"C_2_S_MAIL_TOP":             900,
	"C_2_S_INSTANCE_BASE":        901,
	"C_2_S_INSTANCE_TOP":         1000,
	"C_2_S_RANK_BASE":            1001,
	"C_2_S_RANK_TOP":             1100,
	"C_2_S_TOWER_BASE":           1101,
	"C_2_S_TOWER_TOP":            1200,
	"C_2_S_SHOP_BASE":            1201,
	"C_2_S_SHOP_TOP":             1300,
	"C_2_S_TASK_BASE":            1301,
	"C_2_S_TASK_TOP":             1400,
	"C_2_S_CHAT_BASE":            1401,
	"C_2_S_CHAT_TOP":             1500,
	"C_2_S_BUDDY_BASE":           1501,
	"C_2_S_BUDDY_TOP":            1600,
	"C_2_S_CLONE_BASE":           1601,
	"C_2_S_CLONE_TOP":            1700,
	"C_2_S_STANDARD_BASE":        1701,
	"C_2_S_STANDARD_TOP":         1800,
	"C_2_S_ARENA_BASE":           1801,
	"C_2_S_ARENA_TOP":            1900,
	"C_2_S_CROP_BASE":            1901,
	"C_2_S_CROP_TOP":             2000,
	"C_2_S_CHAMPION_RACE_BASE":   2001,
	"C_2_S_CHAMPION_RACE_TOP":    2100,
	"C_2_S_CROP_WAR_BASE":        2101,
	"C_2_S_CROP_WAR_TOP":         2200,
	"C_2_S_MAIN_BASE":            2201,
	"C_2_S_MAIN_TOP":             2300,
	"C_2_S_REWARD_ACTIVITY_BASE": 2301,
	"C_2_S_REWARD_ACTIVITY_TOP":  2400,
}

func (x ProtocolType) Enum() *ProtocolType {
	p := new(ProtocolType)
	*p = x
	return p
}
func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}
func (x *ProtocolType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProtocolType_value, data, "ProtocolType")
	if err != nil {
		return err
	}
	*x = ProtocolType(value)
	return nil
}
func (ProtocolType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type S2C_GateState_StateType int32

const (
	S2C_GateState_state_ok             S2C_GateState_StateType = 0
	S2C_GateState_state_gs_offline     S2C_GateState_StateType = 1
	S2C_GateState_state_server_is_full S2C_GateState_StateType = 2
)

var S2C_GateState_StateType_name = map[int32]string{
	0: "state_ok",
	1: "state_gs_offline",
	2: "state_server_is_full",
}
var S2C_GateState_StateType_value = map[string]int32{
	"state_ok":             0,
	"state_gs_offline":     1,
	"state_server_is_full": 2,
}

func (x S2C_GateState_StateType) Enum() *S2C_GateState_StateType {
	p := new(S2C_GateState_StateType)
	*p = x
	return p
}
func (x S2C_GateState_StateType) String() string {
	return proto.EnumName(S2C_GateState_StateType_name, int32(x))
}
func (x *S2C_GateState_StateType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(S2C_GateState_StateType_value, data, "S2C_GateState_StateType")
	if err != nil {
		return err
	}
	*x = S2C_GateState_StateType(value)
	return nil
}
func (S2C_GateState_StateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type C2S_SystemTick struct {
	Type             *ProtocolType `protobuf:"varint,1,opt,name=type,enum=protocol.ProtocolType,def=4" json:"type,omitempty"`
	KeepAliveAck     *int32        `protobuf:"varint,2,req,name=keep_alive_ack" json:"keep_alive_ack,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *C2S_SystemTick) Reset()                    { *m = C2S_SystemTick{} }
func (m *C2S_SystemTick) String() string            { return proto.CompactTextString(m) }
func (*C2S_SystemTick) ProtoMessage()               {}
func (*C2S_SystemTick) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_C2S_SystemTick_Type ProtocolType = ProtocolType_S_2_C_KEEP_ALIVE_ACK

func (m *C2S_SystemTick) GetType() ProtocolType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_C2S_SystemTick_Type
}

func (m *C2S_SystemTick) GetKeepAliveAck() int32 {
	if m != nil && m.KeepAliveAck != nil {
		return *m.KeepAliveAck
	}
	return 0
}

type S2C_GateState struct {
	Type             *ProtocolType `protobuf:"varint,1,opt,name=type,enum=protocol.ProtocolType,def=100" json:"type,omitempty"`
	State            *int32        `protobuf:"varint,2,req,name=state" json:"state,omitempty"`
	Key              *int64        `protobuf:"varint,3,req,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *S2C_GateState) Reset()                    { *m = S2C_GateState{} }
func (m *S2C_GateState) String() string            { return proto.CompactTextString(m) }
func (*S2C_GateState) ProtoMessage()               {}
func (*S2C_GateState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_S2C_GateState_Type ProtocolType = ProtocolType_S_2_C_GATE_STATE

func (m *S2C_GateState) GetType() ProtocolType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_S2C_GateState_Type
}

func (m *S2C_GateState) GetState() int32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *S2C_GateState) GetKey() int64 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

type PlayerInfo struct {
	Id               *int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Head             *int32  `protobuf:"varint,3,opt,name=head" json:"head,omitempty"`
	Level            *int32  `protobuf:"varint,4,opt,name=level" json:"level,omitempty"`
	VipLvl           *int32  `protobuf:"varint,5,opt,name=vip_lvl" json:"vip_lvl,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PlayerInfo) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PlayerInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PlayerInfo) GetHead() int32 {
	if m != nil && m.Head != nil {
		return *m.Head
	}
	return 0
}

func (m *PlayerInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *PlayerInfo) GetVipLvl() int32 {
	if m != nil && m.VipLvl != nil {
		return *m.VipLvl
	}
	return 0
}

type Attribute struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Val              *float64 `protobuf:"fixed64,2,opt,name=val" json:"val,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Attribute) Reset()                    { *m = Attribute{} }
func (m *Attribute) String() string            { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()               {}
func (*Attribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Attribute) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Attribute) GetVal() float64 {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return 0
}

type Reward struct {
	Type             *int32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Value            *int32 `protobuf:"varint,2,req,name=value" json:"value,omitempty"`
	Nums             *int32 `protobuf:"varint,3,req,name=nums" json:"nums,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Reward) Reset()                    { *m = Reward{} }
func (m *Reward) String() string            { return proto.CompactTextString(m) }
func (*Reward) ProtoMessage()               {}
func (*Reward) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Reward) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Reward) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Reward) GetNums() int32 {
	if m != nil && m.Nums != nil {
		return *m.Nums
	}
	return 0
}

func init() {
	proto.RegisterType((*C2S_SystemTick)(nil), "protocol.C2S_SystemTick")
	proto.RegisterType((*S2C_GateState)(nil), "protocol.S2C_GateState")
	proto.RegisterType((*PlayerInfo)(nil), "protocol.PlayerInfo")
	proto.RegisterType((*Attribute)(nil), "protocol.Attribute")
	proto.RegisterType((*Reward)(nil), "protocol.Reward")
	proto.RegisterEnum("protocol.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterEnum("protocol.S2C_GateState_StateType", S2C_GateState_StateType_name, S2C_GateState_StateType_value)
}

func init() { proto.RegisterFile("protocol_base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 838 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x68, 0x3b, 0x45,
	0x14, 0xc7, 0xdd, 0xdd, 0xa4, 0x49, 0xa6, 0xf9, 0xb3, 0x99, 0xc6, 0x36, 0x15, 0xff, 0x94, 0xe0,
	0xa1, 0x78, 0xe8, 0x21, 0x82, 0x87, 0xde, 0xa6, 0x9b, 0xb5, 0x59, 0x9a, 0x64, 0xc3, 0xee, 0xda,
	0xe2, 0x69, 0xd8, 0x36, 0x13, 0x0d, 0xd9, 0x36, 0x21, 0xff, 0x24, 0x77, 0x05, 0xbd, 0x29, 0x28,
	0xe8, 0x4d, 0xc1, 0xde, 0x54, 0x14, 0xf4, 0xa6, 0xa0, 0x37, 0x05, 0x05, 0xbd, 0x29, 0x28, 0x78,
	0x54, 0xd0, 0x9b, 0x82, 0x7f, 0xe1, 0x37, 0x6f, 0x66, 0x93, 0xd9, 0x96, 0x1e, 0x7e, 0x97, 0x90,
	0xfd, 0xbe, 0xcf, 0x7b, 0xf3, 0xde, 0xf7, 0x3d, 0xb4, 0x35, 0x9e, 0x8c, 0x66, 0xa3, 0x8b, 0x51,
	0x44, 0xcf, 0xc3, 0x29, 0x3b, 0x10, 0x5f, 0x38, 0xbb, 0x12, 0x6b, 0x3d, 0x54, 0xb4, 0xea, 0x3e,
	0xf5, 0x97, 0xd3, 0x19, 0xbb, 0x0c, 0x06, 0x17, 0x43, 0x7c, 0x88, 0x52, 0xb3, 0xe5, 0x98, 0x55,
	0xb5, 0x3d, 0x6d, 0xbf, 0x58, 0xdf, 0x3e, 0x58, 0xa1, 0x07, 0xdd, 0xf8, 0x4f, 0xc0, 0xa3, 0x87,
	0x15, 0x9f, 0xd6, 0xa9, 0x45, 0x4f, 0x6c, 0xbb, 0x4b, 0x49, 0xcb, 0x39, 0xb5, 0x29, 0xb1, 0x4e,
	0xf0, 0x36, 0x2a, 0x0e, 0x19, 0x1b, 0xd3, 0x30, 0x1a, 0x2c, 0x18, 0x0d, 0x2f, 0x86, 0x55, 0x7d,
	0x4f, 0xdf, 0x4f, 0xd7, 0x3e, 0xd0, 0x50, 0xc1, 0xaf, 0x5b, 0xf4, 0x38, 0x9c, 0x31, 0x7f, 0xc6,
	0x7f, 0xf0, 0x53, 0xf7, 0xf5, 0x8a, 0x29, 0x5f, 0x39, 0x26, 0x81, 0x4d, 0xfd, 0x80, 0xff, 0xe2,
	0x02, 0x4a, 0x4f, 0xa1, 0x80, 0x2c, 0x8c, 0x37, 0x91, 0x31, 0x64, 0xcb, 0xaa, 0xc1, 0x3f, 0x8c,
	0x9a, 0x83, 0x72, 0xa2, 0x38, 0xa4, 0xe2, 0x3c, 0xca, 0x0a, 0x90, 0x8e, 0x86, 0xe6, 0x03, 0xb8,
	0x82, 0x4c, 0xf9, 0xf5, 0xdc, 0x94, 0x8e, 0xfa, 0xfd, 0x68, 0x70, 0xc5, 0x4c, 0x0d, 0x57, 0x51,
	0x45, 0xaa, 0x53, 0x36, 0x59, 0xb0, 0x09, 0x1d, 0x4c, 0x69, 0x7f, 0x1e, 0x45, 0xa6, 0x5e, 0x0b,
	0x10, 0xea, 0x46, 0xe1, 0x92, 0x4d, 0x9c, 0xab, 0xfe, 0x08, 0x23, 0xa4, 0x0f, 0x7a, 0xa2, 0x55,
	0x83, 0xd7, 0x4d, 0x5d, 0x85, 0x97, 0xf0, 0xbe, 0xb6, 0x9f, 0x83, 0xaf, 0xe7, 0x59, 0xd8, 0xe3,
	0x0d, 0x68, 0xbc, 0x1b, 0xde, 0x5c, 0xc4, 0x16, 0x2c, 0xaa, 0xa6, 0xc4, 0x67, 0x09, 0x65, 0x16,
	0x83, 0x31, 0x8d, 0x16, 0x51, 0x35, 0x0d, 0x42, 0xed, 0x71, 0x94, 0x23, 0xb3, 0xd9, 0x64, 0x70,
	0x3e, 0xe7, 0x0e, 0xa8, 0xa2, 0x62, 0x8c, 0x45, 0x18, 0x89, 0x9a, 0x5a, 0xed, 0x49, 0xb4, 0xe1,
	0xb1, 0x17, 0xc2, 0x49, 0x0f, 0xaa, 0xc7, 0x26, 0xe9, 0xb2, 0x3a, 0x87, 0xe6, 0xab, 0xd1, 0xa1,
	0x91, 0xf9, 0xe5, 0x54, 0xcc, 0x9e, 0x7e, 0xe2, 0x3a, 0x87, 0xf2, 0x49, 0xeb, 0x60, 0xb6, 0xbb,
	0x56, 0x64, 0xa6, 0xc0, 0x8b, 0xdb, 0xb6, 0x9a, 0x3d, 0x50, 0x2d, 0xae, 0xfa, 0xb4, 0xe5, 0x1e,
	0x3b, 0x1d, 0x7a, 0x44, 0x7c, 0xdb, 0x64, 0x5c, 0x2d, 0x25, 0xd5, 0xc0, 0xed, 0x9a, 0x5f, 0x6a,
	0x4a, 0xf5, 0xdc, 0x96, 0x2d, 0xd1, 0xaf, 0x34, 0xbc, 0xc5, 0x4f, 0x49, 0xa9, 0x80, 0xfe, 0x93,
	0x40, 0x03, 0xdb, 0x0f, 0x24, 0xfa, 0x6f, 0x02, 0x15, 0x2a, 0xa0, 0xef, 0xe9, 0x0a, 0x6d, 0xda,
	0x9e, 0x2b, 0xd1, 0xf7, 0x75, 0x85, 0x0a, 0x15, 0xd0, 0x57, 0x0d, 0x85, 0x3a, 0x81, 0xdd, 0x96,
	0xe8, 0x6b, 0x86, 0x42, 0x85, 0x0a, 0xe8, 0x1f, 0x06, 0xf7, 0x61, 0x4b, 0x8a, 0x6d, 0x3e, 0xa7,
	0xe7, 0x90, 0x96, 0xc4, 0xff, 0x34, 0xf0, 0x0e, 0xc2, 0xb7, 0x22, 0x90, 0xf2, 0x7d, 0x0a, 0xef,
	0xa2, 0x8a, 0x0c, 0x3c, 0xed, 0x7a, 0x3c, 0xe6, 0xb8, 0xb1, 0x1d, 0x3f, 0xa4, 0x54, 0x35, 0x15,
	0x82, 0xa4, 0x4f, 0xd3, 0xfc, 0xf4, 0xcb, 0xf1, 0x48, 0xc4, 0x0a, 0x1c, 0x4b, 0x66, 0x7c, 0x96,
	0xc6, 0x0f, 0xae, 0x7c, 0x8d, 0x75, 0xc0, 0xdf, 0xde, 0x50, 0x13, 0xb4, 0x89, 0x13, 0xb7, 0xf4,
	0xce, 0x86, 0x9a, 0x40, 0xa8, 0x80, 0xbe, 0x98, 0x51, 0x6f, 0x3a, 0x1d, 0xbe, 0xac, 0x8e, 0x15,
	0x3b, 0xfe, 0x52, 0x46, 0x4d, 0xb0, 0x8e, 0x40, 0xca, 0x2f, 0x99, 0xc4, 0x82, 0x48, 0xe7, 0x44,
	0xe2, 0xbf, 0x66, 0x12, 0x0b, 0x02, 0x15, 0xd0, 0xaf, 0xb3, 0x89, 0xfe, 0xdc, 0x33, 0xdb, 0x93,
	0xec, 0x37, 0xd9, 0xc4, 0xde, 0x84, 0x0c, 0xf0, 0x87, 0x39, 0xa5, 0xfa, 0x4d, 0xb7, 0x2b, 0xd9,
	0x8f, 0x72, 0xaa, 0xae, 0x50, 0x01, 0x7d, 0x1d, 0x25, 0x0a, 0x10, 0x3f, 0x6e, 0xe1, 0x0d, 0x94,
	0x58, 0x3c, 0xa8, 0x80, 0xfe, 0x95, 0x40, 0xad, 0x26, 0x89, 0x6f, 0xe4, 0xef, 0x04, 0x2a, 0x54,
	0x40, 0x7f, 0xdc, 0x54, 0xdd, 0x1e, 0x3d, 0xd3, 0x68, 0x3c, 0x2b, 0xd9, 0x9f, 0x36, 0x55, 0x05,
	0x29, 0x03, 0xfc, 0x79, 0x5e, 0xc1, 0x56, 0xcb, 0xed, 0xc4, 0xae, 0x7d, 0x91, 0x4f, 0x3c, 0x27,
	0x64, 0x80, 0xdf, 0x2d, 0x28, 0x97, 0xc1, 0xc9, 0x06, 0xf1, 0x1a, 0x92, 0xbf, 0x2e, 0x28, 0x97,
	0xd7, 0x11, 0x48, 0x79, 0xb9, 0xa8, 0xea, 0x13, 0xcf, 0xee, 0x10, 0xc9, 0xbf, 0x52, 0x54, 0xf5,
	0xa5, 0x0c, 0xf0, 0x6f, 0x09, 0xd5, 0xf2, 0x56, 0xd6, 0xfd, 0x5e, 0x4c, 0x0c, 0xe9, 0xc5, 0xd6,
	0x7d, 0x5b, 0xc2, 0x8f, 0xa0, 0xea, 0x7a, 0xf2, 0x76, 0x17, 0x6e, 0xcc, 0x23, 0xab, 0xad, 0x7f,
	0x57, 0xc2, 0x0f, 0xa3, 0x9d, 0xbb, 0xc2, 0x90, 0xfc, 0xb1, 0xa9, 0xe6, 0x10, 0x15, 0xcf, 0x48,
	0xbc, 0xd2, 0x4f, 0x4c, 0x35, 0xc7, 0x3a, 0x02, 0x29, 0x6f, 0x96, 0x6f, 0xdc, 0x62, 0x7c, 0xea,
	0x6f, 0x95, 0x6f, 0xdc, 0xa2, 0xbc, 0xf2, 0xff, 0xca, 0xf8, 0x31, 0xf4, 0x50, 0x7c, 0x42, 0xf6,
	0x19, 0x38, 0x01, 0x37, 0x7d, 0xea, 0x04, 0xf1, 0x26, 0xfe, 0x2f, 0xe3, 0x47, 0xd1, 0xee, 0xdd,
	0x00, 0x14, 0xf8, 0x19, 0x1f, 0xe9, 0x4d, 0xed, 0x5e, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xda,
	0x7e, 0xbf, 0x93, 0x06, 0x00, 0x00,
}
