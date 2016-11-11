package logic

import (
	"golog"
	"protocol"
)

func init() {
	servlet := new(AccountServlet)
	G_dispatcher.Register(int32(protocol.Default_C2S_Login_Type), servlet)
}

type AccountServlet struct{}

func (servlet *AccountServlet) DoRequest(ps *PlayerSession, pack *protocol.Packet) bool {
	switch pack.Cmd {
	case int32(protocol.Default_C2S_Login_Type):
		return servlet.doLogin(ps, pack)
	}
	return false
}

func (servlet *AccountServlet) doLogin(ps *PlayerSession, pack *protocol.Packet) bool {
	packData := pack.Data.(*protocol.C2S_Login)

	golog.Debug("AccountLogin", "DoRequest", "Test login", "loginId", packData.GetLoginName(), "ts", packData.GetSource())

	ps.SendData([]byte("from account login"))
	return true
}
