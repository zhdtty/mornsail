package player

import (
//	"fmt"
)

type Role struct {
	LoginId string
	RoleId  string
	Name    string
	Level   int32
	Exp     int32
	HeroBag *HeroBag
}

func NewRole() *Role {
	role := &Role{
		LoginId: "",
		RoleId:  "",
		Name:    "",
		Level:   1,
		Exp:     0,
		HeroBag: NewHeroBag(),
	}
	return role
}

func (role *Role) Init(loginId string, roleId string, name string,
	level int32, exp int32) {
	role.SetLoginId(loginId)
	role.SetRoleId(roleId)
	role.SetName(name)
}

func (role *Role) SetLoginId(val string) {
	role.LoginId = val
}
func (role *Role) SetRoleId(val string) {
	role.RoleId = val
}
func (role *Role) SetName(val string) {
	role.Name = val
}
func (role *Role) AddLevel(val int32) {
	role.Level += val
}
func (role *Role) AddExp(val int32) {
	role.Exp += val
}
