package protocol

import (
	"bytes"
	"encoding/binary"
	"log"
)

const C2S_LOGIN = ID_LOGIN_BASE + 1
const S2C_LOGIN = ID_LOGIN_BASE + 1

func init() {
	c2s_creators[C2S_LOGIN] = func() IUnpacker { return new(C2SLogin) }

	s2c_creators[S2C_LOGIN] = func() IPacker { return new(S2CLogin) }
}

type C2SLogin struct {
	LoginId string
	Ts      int32
}

func (pack *C2SLogin) FromData(buff []byte) bool {
	if len(buff) < 2 {
		log.Println("Invalid packet, len :", len(buff))
		return false
	}
	var strLen int16
	buffer := bytes.NewBuffer(buff)
	binary.Read(buffer, binary.BigEndian, &strLen)
	loginId := make([]byte, strLen)
	buffer.Read(loginId)
	pack.LoginId = string(loginId)
	binary.Read(buffer, binary.BigEndian, &pack.Ts)

	return true
}

type S2CLogin struct {
	ret      int32
	roleId   string
	roleName string
}

func (pack *S2CLogin) ToData() []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, pack.ret)
	binary.Write(buffer, binary.BigEndian, int16(len(pack.roleId)))
	buffer.Write([]byte(pack.roleId))
	binary.Write(buffer, binary.BigEndian, int16(len(pack.roleName)))
	buffer.Write([]byte(pack.roleName))
	return buffer.Bytes()
}
