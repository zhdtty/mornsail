package protocol

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
)

const PACKET_MAX_SIZE = 2048

//for format custom
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

//for format protobuf
var c2s_protobufs map[int32]proto.Message = make(map[int32]proto.Message)

var s2c_protobufs map[int32]proto.Message = make(map[int32]proto.Message)

const (
	FORMAT_CUSTOM = 1 << iota
	FORMAT_PROTOBUF
	FORMAT_JSON
)

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
	if pack.ptype == FORMAT_CUSTOM {
		if p, ok := pack.Data.(IPacker); ok {
			return p.ToData()
		}
	} else if pack.ptype == FORMAT_PROTOBUF {
		data, err := proto.Marshal(*pack.Data.(*proto.Message))
		if err != nil {
			log.Fatal("marshaling error: ", err)
			return nil
		}
		return data
	}
	log.Println("Invalid packet, not IPacker")
	return nil
}

func (pack *Packet) FromData(buff []byte) bool {
	if pack.ptype == FORMAT_CUSTOM {
		if creator, ok := c2s_creators[pack.Cmd]; ok {
			c := creator()
			r := c.FromData(buff)
			pack.Data = c
			return r
		}
	} else if pack.ptype == FORMAT_PROTOBUF {
		if msg, ok := c2s_protobufs[pack.Cmd]; ok {
			err := proto.Unmarshal(buff, msg)
			if err != nil {
				log.Fatal("unmarshaling error: ", err)
				return false
			}
			return true
		}
	}
	log.Println("Invalid packet, cmd not found, cmd :", pack.Cmd)
	return false
}

func ReceiveData(data []byte) *Packet {
	//for custom
	// buffer := bytes.NewBuffer(data)
	// var ptype, cmd int32
	// binary.Read(buffer, binary.BigEndian, &ptype)
	// binary.Read(buffer, binary.BigEndian, &cmd)

	// pack := new(Packet)
	// pack.Cmd = cmd
	// pack.ptype = ptype
	// if pack.FromData(data[8:]) == false {
	// 	return nil
	// }
	// return pack

	buffer := bytes.NewBuffer(data)
	var cmd int32
	binary.Read(buffer, binary.BigEndian, &cmd)

	pack := new(Packet)
	pack.Cmd = cmd
	pack.ptype = FORMAT_PROTOBUF
	if pack.FromData(data[8:]) == false {
		return nil
	}
	return pack

}
