// Code generated by protoc-gen-go.
// source: protocol_login.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LoginProtoType int32

const (
	LoginProtoType_C_2_S_LOGIN          LoginProtoType = 101
	LoginProtoType_S_2_C_LOGIN          LoginProtoType = 102
	LoginProtoType_C_2_S_RAND_NAME      LoginProtoType = 103
	LoginProtoType_S_2_C_RAND_NAME      LoginProtoType = 104
	LoginProtoType_C_2_S_CREATE_ROLE    LoginProtoType = 105
	LoginProtoType_S_2_C_CREATE_ROLE    LoginProtoType = 106
	LoginProtoType_C_2_S_LOAD_ROLE_INFO LoginProtoType = 107
	LoginProtoType_S_2_C_LOAD_ROLE_INFO LoginProtoType = 108
	LoginProtoType_S_2_C_SERVER_TIME    LoginProtoType = 110
)

var LoginProtoType_name = map[int32]string{
	101: "C_2_S_LOGIN",
	102: "S_2_C_LOGIN",
	103: "C_2_S_RAND_NAME",
	104: "S_2_C_RAND_NAME",
	105: "C_2_S_CREATE_ROLE",
	106: "S_2_C_CREATE_ROLE",
	107: "C_2_S_LOAD_ROLE_INFO",
	108: "S_2_C_LOAD_ROLE_INFO",
	110: "S_2_C_SERVER_TIME",
}
var LoginProtoType_value = map[string]int32{
	"C_2_S_LOGIN":          101,
	"S_2_C_LOGIN":          102,
	"C_2_S_RAND_NAME":      103,
	"S_2_C_RAND_NAME":      104,
	"C_2_S_CREATE_ROLE":    105,
	"S_2_C_CREATE_ROLE":    106,
	"C_2_S_LOAD_ROLE_INFO": 107,
	"S_2_C_LOAD_ROLE_INFO": 108,
	"S_2_C_SERVER_TIME":    110,
}

func (x LoginProtoType) Enum() *LoginProtoType {
	p := new(LoginProtoType)
	*p = x
	return p
}
func (x LoginProtoType) String() string {
	return proto.EnumName(LoginProtoType_name, int32(x))
}
func (x *LoginProtoType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LoginProtoType_value, data, "LoginProtoType")
	if err != nil {
		return err
	}
	*x = LoginProtoType(value)
	return nil
}
func (LoginProtoType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type S2C_CreateRole_CreateRoleRet int32

const (
	S2C_CreateRole_st_ok                   S2C_CreateRole_CreateRoleRet = 0
	S2C_CreateRole_st_role_name_duplicated S2C_CreateRole_CreateRoleRet = 1
	S2C_CreateRole_st_fail_not_enough_info S2C_CreateRole_CreateRoleRet = 2
)

var S2C_CreateRole_CreateRoleRet_name = map[int32]string{
	0: "st_ok",
	1: "st_role_name_duplicated",
	2: "st_fail_not_enough_info",
}
var S2C_CreateRole_CreateRoleRet_value = map[string]int32{
	"st_ok":                   0,
	"st_role_name_duplicated": 1,
	"st_fail_not_enough_info": 2,
}

func (x S2C_CreateRole_CreateRoleRet) Enum() *S2C_CreateRole_CreateRoleRet {
	p := new(S2C_CreateRole_CreateRoleRet)
	*p = x
	return p
}
func (x S2C_CreateRole_CreateRoleRet) String() string {
	return proto.EnumName(S2C_CreateRole_CreateRoleRet_name, int32(x))
}
func (x *S2C_CreateRole_CreateRoleRet) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(S2C_CreateRole_CreateRoleRet_value, data, "S2C_CreateRole_CreateRoleRet")
	if err != nil {
		return err
	}
	*x = S2C_CreateRole_CreateRoleRet(value)
	return nil
}
func (S2C_CreateRole_CreateRoleRet) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{5, 0}
}

type C2S_Login struct {
	Type             *LoginProtoType `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=101" json:"type,omitempty"`
	LoginName        *string         `protobuf:"bytes,2,req,name=login_name" json:"login_name,omitempty"`
	PlatformId       *string         `protobuf:"bytes,3,req,name=platform_id" json:"platform_id,omitempty"`
	PartitionId      *string         `protobuf:"bytes,4,req,name=partition_id" json:"partition_id,omitempty"`
	Source           *string         `protobuf:"bytes,5,req,name=source" json:"source,omitempty"`
	Timestamp        *int32          `protobuf:"varint,6,req,name=timestamp" json:"timestamp,omitempty"`
	Fcm              *int32          `protobuf:"varint,7,req,name=fcm" json:"fcm,omitempty"`
	Ticket           *string         `protobuf:"bytes,8,req,name=ticket" json:"ticket,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *C2S_Login) Reset()                    { *m = C2S_Login{} }
func (m *C2S_Login) String() string            { return proto.CompactTextString(m) }
func (*C2S_Login) ProtoMessage()               {}
func (*C2S_Login) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

const Default_C2S_Login_Type LoginProtoType = LoginProtoType_C_2_S_LOGIN

func (m *C2S_Login) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_C2S_Login_Type
}

func (m *C2S_Login) GetLoginName() string {
	if m != nil && m.LoginName != nil {
		return *m.LoginName
	}
	return ""
}

func (m *C2S_Login) GetPlatformId() string {
	if m != nil && m.PlatformId != nil {
		return *m.PlatformId
	}
	return ""
}

func (m *C2S_Login) GetPartitionId() string {
	if m != nil && m.PartitionId != nil {
		return *m.PartitionId
	}
	return ""
}

func (m *C2S_Login) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *C2S_Login) GetTimestamp() int32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *C2S_Login) GetFcm() int32 {
	if m != nil && m.Fcm != nil {
		return *m.Fcm
	}
	return 0
}

func (m *C2S_Login) GetTicket() string {
	if m != nil && m.Ticket != nil {
		return *m.Ticket
	}
	return ""
}

type S2C_Login struct {
	Type             *LoginProtoType   `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=102" json:"type,omitempty"`
	Ret              *int32            `protobuf:"varint,2,req,name=ret" json:"ret,omitempty"`
	Roles            []*S2C_Login_Role `protobuf:"bytes,3,rep,name=roles" json:"roles,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *S2C_Login) Reset()                    { *m = S2C_Login{} }
func (m *S2C_Login) String() string            { return proto.CompactTextString(m) }
func (*S2C_Login) ProtoMessage()               {}
func (*S2C_Login) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

const Default_S2C_Login_Type LoginProtoType = LoginProtoType_S_2_C_LOGIN

func (m *S2C_Login) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_S2C_Login_Type
}

func (m *S2C_Login) GetRet() int32 {
	if m != nil && m.Ret != nil {
		return *m.Ret
	}
	return 0
}

func (m *S2C_Login) GetRoles() []*S2C_Login_Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type S2C_Login_Role struct {
	RoleId           *int64  `protobuf:"varint,1,req,name=role_id" json:"role_id,omitempty"`
	Level            *int32  `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Name             *string `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	Head             *int32  `protobuf:"varint,4,req,name=head" json:"head,omitempty"`
	ForceVal         *int32  `protobuf:"varint,5,opt,name=force_val" json:"force_val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *S2C_Login_Role) Reset()                    { *m = S2C_Login_Role{} }
func (m *S2C_Login_Role) String() string            { return proto.CompactTextString(m) }
func (*S2C_Login_Role) ProtoMessage()               {}
func (*S2C_Login_Role) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *S2C_Login_Role) GetRoleId() int64 {
	if m != nil && m.RoleId != nil {
		return *m.RoleId
	}
	return 0
}

func (m *S2C_Login_Role) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *S2C_Login_Role) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *S2C_Login_Role) GetHead() int32 {
	if m != nil && m.Head != nil {
		return *m.Head
	}
	return 0
}

func (m *S2C_Login_Role) GetForceVal() int32 {
	if m != nil && m.ForceVal != nil {
		return *m.ForceVal
	}
	return 0
}

type C2S_RandName struct {
	Type             *LoginProtoType `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=103" json:"type,omitempty"`
	Sex              *int32          `protobuf:"varint,2,req,name=sex" json:"sex,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *C2S_RandName) Reset()                    { *m = C2S_RandName{} }
func (m *C2S_RandName) String() string            { return proto.CompactTextString(m) }
func (*C2S_RandName) ProtoMessage()               {}
func (*C2S_RandName) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

const Default_C2S_RandName_Type LoginProtoType = LoginProtoType_C_2_S_RAND_NAME

func (m *C2S_RandName) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_C2S_RandName_Type
}

func (m *C2S_RandName) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

type S2C_RandName struct {
	Type             *LoginProtoType `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=104" json:"type,omitempty"`
	Ret              *int32          `protobuf:"varint,2,req,name=ret" json:"ret,omitempty"`
	Name             *string         `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *S2C_RandName) Reset()                    { *m = S2C_RandName{} }
func (m *S2C_RandName) String() string            { return proto.CompactTextString(m) }
func (*S2C_RandName) ProtoMessage()               {}
func (*S2C_RandName) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

const Default_S2C_RandName_Type LoginProtoType = LoginProtoType_S_2_C_RAND_NAME

func (m *S2C_RandName) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_S2C_RandName_Type
}

func (m *S2C_RandName) GetRet() int32 {
	if m != nil && m.Ret != nil {
		return *m.Ret
	}
	return 0
}

func (m *S2C_RandName) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type C2S_CreateRole struct {
	Type             *LoginProtoType `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=105" json:"type,omitempty"`
	Name             *string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Head             *int32          `protobuf:"varint,3,opt,name=head" json:"head,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *C2S_CreateRole) Reset()                    { *m = C2S_CreateRole{} }
func (m *C2S_CreateRole) String() string            { return proto.CompactTextString(m) }
func (*C2S_CreateRole) ProtoMessage()               {}
func (*C2S_CreateRole) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

const Default_C2S_CreateRole_Type LoginProtoType = LoginProtoType_C_2_S_CREATE_ROLE

func (m *C2S_CreateRole) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_C2S_CreateRole_Type
}

func (m *C2S_CreateRole) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *C2S_CreateRole) GetHead() int32 {
	if m != nil && m.Head != nil {
		return *m.Head
	}
	return 0
}

type S2C_CreateRole struct {
	Type             *LoginProtoType `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=106" json:"type,omitempty"`
	Ret              *int32          `protobuf:"varint,2,req,name=ret" json:"ret,omitempty"`
	RoleId           *int64          `protobuf:"varint,3,req,name=role_id" json:"role_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *S2C_CreateRole) Reset()                    { *m = S2C_CreateRole{} }
func (m *S2C_CreateRole) String() string            { return proto.CompactTextString(m) }
func (*S2C_CreateRole) ProtoMessage()               {}
func (*S2C_CreateRole) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

const Default_S2C_CreateRole_Type LoginProtoType = LoginProtoType_S_2_C_CREATE_ROLE

func (m *S2C_CreateRole) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_S2C_CreateRole_Type
}

func (m *S2C_CreateRole) GetRet() int32 {
	if m != nil && m.Ret != nil {
		return *m.Ret
	}
	return 0
}

func (m *S2C_CreateRole) GetRoleId() int64 {
	if m != nil && m.RoleId != nil {
		return *m.RoleId
	}
	return 0
}

type C2S_LoadRoleInfo struct {
	Type             *LoginProtoType `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=107" json:"type,omitempty"`
	RoleId           *int64          `protobuf:"varint,2,req,name=role_id" json:"role_id,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *C2S_LoadRoleInfo) Reset()                    { *m = C2S_LoadRoleInfo{} }
func (m *C2S_LoadRoleInfo) String() string            { return proto.CompactTextString(m) }
func (*C2S_LoadRoleInfo) ProtoMessage()               {}
func (*C2S_LoadRoleInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

const Default_C2S_LoadRoleInfo_Type LoginProtoType = LoginProtoType_C_2_S_LOAD_ROLE_INFO

func (m *C2S_LoadRoleInfo) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_C2S_LoadRoleInfo_Type
}

func (m *C2S_LoadRoleInfo) GetRoleId() int64 {
	if m != nil && m.RoleId != nil {
		return *m.RoleId
	}
	return 0
}

type S2C_LoadRoleInfo struct {
	Type             *LoginProtoType `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=108" json:"type,omitempty"`
	Ret              *int32          `protobuf:"varint,2,req,name=ret" json:"ret,omitempty"`
	RoleId           *int64          `protobuf:"varint,3,req,name=role_id" json:"role_id,omitempty"`
	RoleName         *string         `protobuf:"bytes,4,req,name=role_name" json:"role_name,omitempty"`
	Level            *int32          `protobuf:"varint,5,req,name=level" json:"level,omitempty"`
	Head             *int32          `protobuf:"varint,6,req,name=head" json:"head,omitempty"`
	Exp              *int32          `protobuf:"varint,7,req,name=exp" json:"exp,omitempty"`
	VipLvl           *int32          `protobuf:"varint,8,req,name=vip_lvl" json:"vip_lvl,omitempty"`
	VipExp           *int32          `protobuf:"varint,9,req,name=vip_exp" json:"vip_exp,omitempty"`
	ChargeCash       *int32          `protobuf:"varint,10,req,name=charge_cash" json:"charge_cash,omitempty"`
	ChargeDiamond    *int32          `protobuf:"varint,11,req,name=charge_diamond" json:"charge_diamond,omitempty"`
	Diamond          *int32          `protobuf:"varint,12,req,name=diamond" json:"diamond,omitempty"`
	Gold             *int32          `protobuf:"varint,13,req,name=gold" json:"gold,omitempty"`
	Strength         *int32          `protobuf:"varint,14,req,name=strength" json:"strength,omitempty"`
	SkillPoint       *int32          `protobuf:"varint,15,req,name=skill_point" json:"skill_point,omitempty"`
	TrainPoint       *int32          `protobuf:"varint,16,req,name=train_point" json:"train_point,omitempty"`
	FirstLog         *int64          `protobuf:"varint,17,req,name=first_log" json:"first_log,omitempty"`
	VipGiftFirstLog  *int64          `protobuf:"varint,18,req,name=vip_gift_first_log" json:"vip_gift_first_log,omitempty"`
	TrialVipLevel    *int32          `protobuf:"varint,19,req,name=trial_vip_level" json:"trial_vip_level,omitempty"`
	TrialVipTime     *int32          `protobuf:"varint,20,req,name=trial_vip_time" json:"trial_vip_time,omitempty"`
	Power            *int32          `protobuf:"varint,21,req,name=power" json:"power,omitempty"`
	Morale           *int32          `protobuf:"varint,22,req,name=morale" json:"morale,omitempty"`
	ArenaCoin        *int32          `protobuf:"varint,23,opt,name=arena_coin" json:"arena_coin,omitempty"`
	TowerCoin        *int32          `protobuf:"varint,24,opt,name=tower_coin" json:"tower_coin,omitempty"`
	CropCoin         *int32          `protobuf:"varint,25,opt,name=crop_coin" json:"crop_coin,omitempty"`
	ItemCoin         *int32          `protobuf:"varint,26,opt,name=item_coin" json:"item_coin,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *S2C_LoadRoleInfo) Reset()                    { *m = S2C_LoadRoleInfo{} }
func (m *S2C_LoadRoleInfo) String() string            { return proto.CompactTextString(m) }
func (*S2C_LoadRoleInfo) ProtoMessage()               {}
func (*S2C_LoadRoleInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

const Default_S2C_LoadRoleInfo_Type LoginProtoType = LoginProtoType_S_2_C_LOAD_ROLE_INFO

func (m *S2C_LoadRoleInfo) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_S2C_LoadRoleInfo_Type
}

func (m *S2C_LoadRoleInfo) GetRet() int32 {
	if m != nil && m.Ret != nil {
		return *m.Ret
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetRoleId() int64 {
	if m != nil && m.RoleId != nil {
		return *m.RoleId
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetRoleName() string {
	if m != nil && m.RoleName != nil {
		return *m.RoleName
	}
	return ""
}

func (m *S2C_LoadRoleInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetHead() int32 {
	if m != nil && m.Head != nil {
		return *m.Head
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetExp() int32 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetVipLvl() int32 {
	if m != nil && m.VipLvl != nil {
		return *m.VipLvl
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetVipExp() int32 {
	if m != nil && m.VipExp != nil {
		return *m.VipExp
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetChargeCash() int32 {
	if m != nil && m.ChargeCash != nil {
		return *m.ChargeCash
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetChargeDiamond() int32 {
	if m != nil && m.ChargeDiamond != nil {
		return *m.ChargeDiamond
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetDiamond() int32 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetGold() int32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetStrength() int32 {
	if m != nil && m.Strength != nil {
		return *m.Strength
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetSkillPoint() int32 {
	if m != nil && m.SkillPoint != nil {
		return *m.SkillPoint
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetTrainPoint() int32 {
	if m != nil && m.TrainPoint != nil {
		return *m.TrainPoint
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetFirstLog() int64 {
	if m != nil && m.FirstLog != nil {
		return *m.FirstLog
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetVipGiftFirstLog() int64 {
	if m != nil && m.VipGiftFirstLog != nil {
		return *m.VipGiftFirstLog
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetTrialVipLevel() int32 {
	if m != nil && m.TrialVipLevel != nil {
		return *m.TrialVipLevel
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetTrialVipTime() int32 {
	if m != nil && m.TrialVipTime != nil {
		return *m.TrialVipTime
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetPower() int32 {
	if m != nil && m.Power != nil {
		return *m.Power
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetMorale() int32 {
	if m != nil && m.Morale != nil {
		return *m.Morale
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetArenaCoin() int32 {
	if m != nil && m.ArenaCoin != nil {
		return *m.ArenaCoin
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetTowerCoin() int32 {
	if m != nil && m.TowerCoin != nil {
		return *m.TowerCoin
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetCropCoin() int32 {
	if m != nil && m.CropCoin != nil {
		return *m.CropCoin
	}
	return 0
}

func (m *S2C_LoadRoleInfo) GetItemCoin() int32 {
	if m != nil && m.ItemCoin != nil {
		return *m.ItemCoin
	}
	return 0
}

type S2C_ServerTime struct {
	Type             *LoginProtoType `protobuf:"varint,1,opt,name=type,enum=protocol.LoginProtoType,def=110" json:"type,omitempty"`
	Time             *int32          `protobuf:"varint,2,req,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *S2C_ServerTime) Reset()                    { *m = S2C_ServerTime{} }
func (m *S2C_ServerTime) String() string            { return proto.CompactTextString(m) }
func (*S2C_ServerTime) ProtoMessage()               {}
func (*S2C_ServerTime) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

const Default_S2C_ServerTime_Type LoginProtoType = LoginProtoType_S_2_C_SERVER_TIME

func (m *S2C_ServerTime) GetType() LoginProtoType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_S2C_ServerTime_Type
}

func (m *S2C_ServerTime) GetTime() int32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*C2S_Login)(nil), "protocol.C2S_Login")
	proto.RegisterType((*S2C_Login)(nil), "protocol.S2C_Login")
	proto.RegisterType((*S2C_Login_Role)(nil), "protocol.S2C_Login.Role")
	proto.RegisterType((*C2S_RandName)(nil), "protocol.C2S_RandName")
	proto.RegisterType((*S2C_RandName)(nil), "protocol.S2C_RandName")
	proto.RegisterType((*C2S_CreateRole)(nil), "protocol.C2S_CreateRole")
	proto.RegisterType((*S2C_CreateRole)(nil), "protocol.S2C_CreateRole")
	proto.RegisterType((*C2S_LoadRoleInfo)(nil), "protocol.C2S_LoadRoleInfo")
	proto.RegisterType((*S2C_LoadRoleInfo)(nil), "protocol.S2C_LoadRoleInfo")
	proto.RegisterType((*S2C_ServerTime)(nil), "protocol.S2C_ServerTime")
	proto.RegisterEnum("protocol.LoginProtoType", LoginProtoType_name, LoginProtoType_value)
	proto.RegisterEnum("protocol.S2C_CreateRole_CreateRoleRet", S2C_CreateRole_CreateRoleRet_name, S2C_CreateRole_CreateRoleRet_value)
}

func init() { proto.RegisterFile("protocol_login.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x8e, 0x22, 0x37,
	0x10, 0x0e, 0x3f, 0xbd, 0x3b, 0x18, 0x86, 0xf1, 0x78, 0xd9, 0x1d, 0x67, 0x72, 0x59, 0x71, 0x09,
	0xca, 0x81, 0x03, 0x39, 0x44, 0x89, 0x72, 0x41, 0x2c, 0x89, 0x90, 0x76, 0x99, 0xa8, 0x21, 0x91,
	0x72, 0xb2, 0x9c, 0x6e, 0xd3, 0x38, 0xb8, 0xdb, 0x2d, 0xb7, 0x87, 0x6c, 0x5e, 0x27, 0x2f, 0x92,
	0x27, 0xc8, 0x31, 0xef, 0x93, 0xb2, 0x9b, 0x6e, 0x60, 0x77, 0x15, 0xcd, 0xcc, 0x09, 0xfc, 0x95,
	0xeb, 0xe7, 0xab, 0xf2, 0x57, 0x8d, 0x06, 0xb9, 0xd1, 0x56, 0x47, 0x5a, 0x31, 0xa5, 0x13, 0x99,
	0x8d, 0xfd, 0x91, 0x5c, 0x54, 0xe8, 0xf0, 0xef, 0x06, 0xea, 0xcc, 0x26, 0x2b, 0xf6, 0xd6, 0x59,
	0xc9, 0xd7, 0xa8, 0x6d, 0xff, 0xcc, 0x05, 0x6d, 0xbc, 0x6e, 0x8c, 0xfa, 0x13, 0x3a, 0xae, 0xae,
	0x8d, 0xbd, 0xf9, 0x27, 0x77, 0x5a, 0x83, 0xfd, 0xbb, 0xee, 0x8c, 0x4d, 0x18, 0x38, 0xdd, 0xfd,
	0xb8, 0x58, 0x12, 0x82, 0x90, 0x8f, 0xcd, 0x32, 0x9e, 0x0a, 0xda, 0x7c, 0xdd, 0x1c, 0x75, 0xc8,
	0x0b, 0xd4, 0xcd, 0x15, 0xb7, 0x1b, 0x6d, 0x52, 0x26, 0x63, 0xda, 0xf2, 0xe0, 0x00, 0xf5, 0x72,
	0x6e, 0xac, 0xb4, 0x52, 0x67, 0x0e, 0x6d, 0x7b, 0xb4, 0x8f, 0x9e, 0x15, 0xfa, 0xde, 0x44, 0x82,
	0x06, 0xfe, 0x7c, 0x8d, 0x3a, 0x56, 0xa6, 0xa2, 0xb0, 0x3c, 0xcd, 0xe9, 0x33, 0x80, 0x02, 0xd2,
	0x45, 0xad, 0x4d, 0x94, 0xd2, 0xe7, 0xfe, 0x00, 0xf7, 0xad, 0x8c, 0x76, 0xc2, 0xd2, 0x0b, 0x77,
	0x7f, 0xf8, 0x0f, 0x30, 0x58, 0x4d, 0x66, 0x8f, 0x65, 0xb0, 0x02, 0x06, 0xb3, 0x03, 0x03, 0x88,
	0x6f, 0x20, 0x5e, 0xd3, 0xc7, 0xff, 0x12, 0x05, 0x46, 0x2b, 0x51, 0x40, 0xd1, 0xad, 0x51, 0xf7,
	0x34, 0x44, 0x9d, 0x65, 0x1c, 0xc2, 0x85, 0xdb, 0x9f, 0x51, 0xdb, 0xfd, 0x92, 0x2b, 0xf4, 0xdc,
	0x39, 0x38, 0x46, 0x0d, 0x88, 0xd0, 0x22, 0x97, 0x28, 0x50, 0x62, 0x2f, 0xd4, 0x21, 0x60, 0x0f,
	0xb5, 0x7d, 0x67, 0xca, 0x26, 0xc0, 0x69, 0x2b, 0x78, 0x49, 0x3e, 0x70, 0x64, 0xa1, 0x47, 0x91,
	0x60, 0x7b, 0xae, 0x80, 0x7f, 0x63, 0x14, 0x0c, 0xd7, 0xa8, 0xe7, 0x06, 0x12, 0xf2, 0x2c, 0x5e,
	0x82, 0x1b, 0xf9, 0xe6, 0x81, 0x8c, 0xae, 0xca, 0x99, 0x84, 0xd3, 0xe5, 0x1b, 0xb6, 0x9c, 0xbe,
	0x9b, 0x3b, 0x56, 0x85, 0x78, 0x5f, 0x16, 0x31, 0xfc, 0x0d, 0xf5, 0x5c, 0xf9, 0x8f, 0x8f, 0x5a,
	0xf6, 0xe9, 0x2c, 0xea, 0xb1, 0x57, 0x67, 0xd4, 0x86, 0x09, 0xea, 0xbb, 0xca, 0x67, 0x46, 0x70,
	0x2b, 0x7c, 0x6b, 0xbe, 0x7d, 0x60, 0x96, 0xeb, 0xb2, 0xf6, 0x59, 0x38, 0x9f, 0xae, 0xe7, 0x2c,
	0xbc, 0x7b, 0x3b, 0xaf, 0x43, 0x37, 0xc1, 0xf5, 0xd8, 0xb5, 0x96, 0x6f, 0x11, 0x3c, 0xda, 0xbe,
	0x63, 0xf3, 0x94, 0x4c, 0x25, 0x9f, 0xd3, 0x4c, 0x67, 0x8c, 0x4e, 0x86, 0xe9, 0x48, 0xb5, 0x60,
	0x1c, 0x97, 0xc7, 0x34, 0xa1, 0xb0, 0xa4, 0x83, 0x82, 0xc2, 0x32, 0xbd, 0xc3, 0x9f, 0x91, 0x2f,
	0xd0, 0x0d, 0xfc, 0xf5, 0xf7, 0x5d, 0xad, 0x2c, 0xbe, 0xcf, 0x95, 0x8c, 0xe0, 0x6e, 0x8c, 0x1b,
	0x07, 0xe3, 0x86, 0x4b, 0xc5, 0x32, 0x6d, 0x99, 0xc8, 0xf4, 0x7d, 0xb2, 0x65, 0x32, 0xdb, 0x68,
	0xdc, 0x1c, 0x72, 0x84, 0x4b, 0xd5, 0xf1, 0xd8, 0xc5, 0x5d, 0x00, 0x4a, 0xbe, 0x7f, 0x20, 0x85,
	0x41, 0x25, 0xbe, 0xe9, 0x1b, 0x4f, 0x80, 0x2d, 0x96, 0x3f, 0xdc, 0x9d, 0x16, 0xde, 0xf4, 0x85,
	0xff, 0xd5, 0x46, 0xb8, 0x7c, 0xb1, 0x4f, 0xc9, 0x51, 0xc9, 0xe3, 0x2c, 0xc7, 0xff, 0x76, 0xca,
	0xbd, 0xe5, 0xba, 0x15, 0x07, 0x6d, 0xd7, 0x4a, 0x08, 0xaa, 0xe7, 0xe2, 0xa7, 0x58, 0xab, 0x5a,
	0xbc, 0xcf, 0x0f, 0xaa, 0x86, 0x68, 0x7b, 0x99, 0x33, 0xb5, 0x57, 0x5e, 0xd6, 0x35, 0xe0, 0x6e,
	0x74, 0x3c, 0x00, 0x2b, 0x25, 0xda, 0x72, 0x93, 0x08, 0x16, 0xf1, 0x62, 0x4b, 0x91, 0x07, 0x5f,
	0xa1, 0xfe, 0x01, 0x8c, 0x25, 0x4f, 0x75, 0x16, 0xd3, 0x6e, 0xe5, 0x5d, 0x01, 0xbd, 0x2a, 0x75,
	0xa2, 0x55, 0x4c, 0x2f, 0xfd, 0x09, 0xa3, 0x8b, 0xc2, 0x1a, 0x91, 0x25, 0x76, 0x4b, 0xfb, 0x55,
	0xf4, 0x62, 0x27, 0x95, 0x62, 0xb9, 0x96, 0x99, 0xa5, 0x57, 0x15, 0x68, 0x0d, 0x87, 0xcd, 0x56,
	0x82, 0xb8, 0x96, 0xac, 0x34, 0x30, 0x5a, 0x58, 0x7a, 0xf4, 0xda, 0x33, 0xbf, 0x45, 0xc4, 0xd5,
	0x9a, 0xc8, 0x0d, 0x0c, 0xbc, 0xb6, 0x11, 0x6f, 0xbb, 0x41, 0x57, 0xd6, 0x48, 0xae, 0x98, 0xa7,
	0xe7, 0x9b, 0xf1, 0xa2, 0x2a, 0xfd, 0x68, 0x70, 0x1b, 0x8f, 0x0e, 0x3c, 0x0e, 0x3d, 0xcb, 0xf5,
	0x1f, 0xc2, 0xd0, 0x97, 0xd5, 0xba, 0x4b, 0xb5, 0xe1, 0x4a, 0xd0, 0x57, 0xfe, 0x0c, 0xdb, 0x96,
	0x43, 0xe5, 0x9c, 0x45, 0x50, 0x13, 0xbd, 0x71, 0x7a, 0x70, 0x98, 0x75, 0x2e, 0x25, 0x46, 0x3d,
	0x06, 0x65, 0x46, 0x46, 0xe7, 0x25, 0xf4, 0x79, 0x05, 0x49, 0x2b, 0xd2, 0x12, 0xba, 0xf5, 0x4a,
	0xfa, 0xb5, 0x14, 0xd2, 0x4a, 0x98, 0xbd, 0x30, 0x6b, 0x28, 0xe2, 0xb1, 0x42, 0x5a, 0xcd, 0xc3,
	0x5f, 0xe6, 0x21, 0x5b, 0x2f, 0xde, 0x79, 0xc9, 0x7a, 0x1e, 0xfe, 0x7d, 0x7c, 0xf5, 0x2f, 0x88,
	0xf4, 0xdc, 0x07, 0xa6, 0x72, 0xfa, 0xe1, 0xc0, 0x1e, 0x38, 0xd9, 0xc3, 0x78, 0x03, 0x1d, 0xff,
	0x70, 0x8d, 0xe1, 0xc4, 0x81, 0x1f, 0x6c, 0x21, 0xbc, 0x25, 0x2f, 0xd1, 0xc7, 0x4b, 0x03, 0x4b,
	0x07, 0x7f, 0xa4, 0x70, 0xfc, 0x3b, 0xa1, 0xe8, 0x93, 0xaa, 0xc1, 0x3b, 0x67, 0xf9, 0xd4, 0x5b,
	0xc7, 0xea, 0x18, 0xea, 0x84, 0x23, 0xce, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x60, 0x3e, 0xf1,
	0x5c, 0x52, 0x07, 0x00, 0x00,
}
