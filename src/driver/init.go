package driver

import (
	"config"
	"golog"
)

func Init() {
	golog.Info("Init", "Init", "-------init db driver")
	PGPool = newPostgres()
	RedisPool = newPool(*config.FLAG_REDIS_ADDR, *config.FLAG_REDIS_PASSWD)
	Mgo = newMongo()
}
