package driver

import (
	"config"
	"github.com/garyburd/redigo/redis"
	"time"
)

func newPool(server, password string) *redis.Pool {
	//Set MaxIdle >= MaxActive, then will no report "Cannot assign requested address"
	//Or will create new connection, if MaxIdle not enough
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   50,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("PING"); err != nil {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
	}
}

var RedisPool *redis.Pool = newPool(*config.FLAG_REDIS_ADDR, *config.FLAG_REDIS_PASSWD)
