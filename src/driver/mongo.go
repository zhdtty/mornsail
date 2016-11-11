package driver

import (
	"config"
	"fmt"
	"gopkg.in/mgo.v2"
)

func newMongo() *mgo.Session {
	if len(*config.FLAG_MONGO_ADDR) <= 0 {
		return nil
	}
	session, err := mgo.Dial(*config.FLAG_MONGO_ADDR)
	if err != nil {
		fmt.Println("Invalid mongo address")
		return nil
	}
	session.SetMode(mgo.Monotonic, true)
	//	session.SetPoolLimit(100)
	return session
}

var Mgo *mgo.Session // = newMongo()
