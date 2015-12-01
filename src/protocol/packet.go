package protocol

import (
	"bytes"
	"encoding/binary"
	"log"
)

const PACKET_MAX_SIZE = 2048

type IPacker interface {
	ToData() []byte
}

type IUnpacker interface {
	FromData(buff []byte) bool
}

type C2SCreator func() IUnpacker

var c2s_creators map[int32]C2SCreator = make(map[int32]C2SCreator)

type S2CCreator func() IPacker

var s2c_creators map[int32]S2CCreator = make(map[int32]S2CCreator)

type Packet struct {
	ptype int32 //0:custom 1:protobuf 2:json ...
	Cmd   int32 //protoId
	Data  interface{}
}

func (pack *Packet) ToData() []byte {
	if pack.Data == nil {
		log.Println("Invalid packet, cmd : ", pack.Cmd)
		return nil
	}
	if p, ok := pack.Data.(IPacker); ok {
		return p.ToData()
	}
	log.Println("Invalid packet, not IPacker")
	return nil
}

func (pack *Packet) FromData(buff []byte) bool {
	if creator, ok := c2s_creators[pack.Cmd]; ok {
		c := creator()
		r := c.FromData(buff)
		pack.Data = c
		return r
	}
	log.Println("Invalid packet, cmd not found, cmd :", pack.Cmd)
	return false
}

func ReceiveData(data []byte) *Packet {
	buffer := bytes.NewBuffer(data)
	var ptype, cmd int32
	binary.Read(buffer, binary.BigEndian, &ptype)
	binary.Read(buffer, binary.BigEndian, &cmd)

	pack := new(Packet)
	pack.Cmd = cmd
	pack.ptype = ptype
	if pack.FromData(data[8:]) == false {
		return nil
	}
	return pack
}
