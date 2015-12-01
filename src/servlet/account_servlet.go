package servlet

import (
	//"fmt"
	"protocol"
	//	"server/server"
	"database/sql"
	"driver"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
	"golog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	G_dispatcher.Register(protocol.C2S_LOGIN, new(AccountLogin))
}

type AccountLogin struct{}

var count int32 = 0
var rediscount int32 = 0
var mongocount int32 = 0

func AccessRedis(loginId string, ts int32) {
	defer func() {
		msg := recover()
		if msg != nil {
			panic(msg)
		}
	}()
	conn := driver.RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("MULTI")
	if err != nil {
		golog.Error("AccessRedis", "AccessRedis", "DO MULTI", "msg", err)
		panic(err)
	}

	_, err = conn.Do("hget", "rank_arena_hash", "AI5653f448aff8bc6d00002351")
	if err != nil {
		golog.Error("AccessRedis", "AccessRedis", "DO hgetall", "msg", err)
		panic(err)
	}
	/*
		_, err = conn.Do("lrange", "rank_arena_list", 0, -1)
		if err != nil {
			golog.Error("AccessRedis", "AccessRedis", "Do lrange", "msg", err)
			panic(err)
		}
	*/
	_, err = conn.Do("EXEC")
	if err != nil {
		golog.Error("AccessRedis", "AccessRedis", "Do EXEC", "msg", err)
		panic(err)
	}
	rediscount++
	if rediscount%1000 == 0 {
		golog.Info("AccessRedis", "AccessRedis", "Access redis over", "count", rediscount)
	}

	/*
		results, err := redis.MultiBulk(reply, err)
		if err != nil {
			golog.Error("AccessRedis", "AccessRedis", "multibulk err", "msg", err)
			panic(err)
		}
		if len(results) < 2 {
			golog.Error("AccessRedis", "AccessRedis", "results len error", "len", len(results), "loginid", loginId, "index", ts)
			panic("results array len error")
		}
	*/
	/*
		{
			_, err := redis.IntMap(results[0], err)
				for k, v := range intMap {
					//fmt.Println(k, v)
					//golog.Debug("AccessRedis", "AccessRedis", "hgetall rank_arena_hash", "k", k, "v", v)
				}
			_, err = redis.Strings(results[1], err)
				for _, _ = range strs {
					//fmt.Println("lrange", val)
					//golog.Debug("AccessRedis", "AccessRedis", "lrange rank_arena_list", "val", val)
				}
		}
	*/
}

func AccessPostgresql(db *sql.DB, loginId string) {
	defer func() {
		msg := recover()
		if msg != nil {
			golog.Error("1", "1", "1", "msg", msg)
			panic(msg)
		}
	}()
	/*
		rows, err := db.Query("select loginname,accname from account limit 3") //"SELECT name FROM users WHERE age = $1", age)
		//	var cnt int
		//	err := db.QueryRow("select count(*) from account").Scan(&cnt)
		if err != nil {
			golog.Error("AccessPostgresql", "AccessPostgresql", "Access postgresql query", "err", err)
		}
		//	golog.Info("AccPostgres", "", "", "cnt", cnt, "loginId", loginId)

		defer rows.Close()
		var accname, loginname string
		for rows.Next() {
			err = rows.Scan(&loginname, &accname)
			if err != nil {
				golog.Error("AccessPostgresql", "AccessPostgresql", "Access postgresql scan", "err", err)
			}
			golog.Info("AccessPostgresql", "AccessPostgresql", "Access postgresql over", "loginname", loginname, "accname", accname)
		}
	*/

	var stmt *sql.Stmt
	var err error

	stmt, err = db.Prepare("select loginname,accname from account limit 3")
	if err != nil {
		golog.Error("AccessPostgresql", "AccessPostgresql", "Access postgresql prepare", "err", err)
		return
	}

	var rows *sql.Rows

	rows, err = stmt.Query()
	if err != nil {
		golog.Error("AccessPostgresql", "AccessPostgresql", "Access postgresql query", "err", err)
		return
	}

	defer stmt.Close()

	for rows.Next() {
		var loginname, accname string
		err = rows.Scan(&loginname, &accname)
		if err != nil {
			golog.Error("AccessPostgresql", "AccessPostgresql", "Access postgresql scan", "err", err)
			return
		}

		//			golog.Info("AccessPostgresql", "AccessPostgresql", "Access postgresql over", "loginname", loginname, "accname", accname)
	}
	count++
	if count%1000 == 0 {
		golog.Info("AccessPostgresql", "AccessPostgresql", "Access postgresql over", "count", count)
	}
}

func AccessMongo(session *mgo.Session, loginId string) {

	/*
		c := session.DB("card").C("role")
		pipe := c.Pipe([]bson.M{
			bson.M{"$match": bson.M{"level": bson.M{"$gte": 1}}},
			bson.M{"$group": bson.M{"_id": "$loginid", "total": bson.M{"$sum": "$money"}}},
			bson.M{"$sort": bson.M{"total": -1}},
		})
		resp := []bson.M{}
		err := pipe.All(&resp)
		if err != nil {
			//handle error
		}
	*/
	c := session.DB("card").C("role")
	var result interface{}
	err := c.Find(bson.M{"loginid": "g001_28_282843150017"}).One(&result)
	if err != nil {
		//		log.Fatal(err)
	}

	mongocount++
	if mongocount%1000 == 0 {
		golog.Info("AccessMongo", "AccessMongo", "Access mongo over", "count", mongocount)
	}
	//	for _, v := range resp {
	//		fmt.Println("loginid", v["_id"], "sum", v["total"])
	//	}
}

func (servlet *AccountLogin) DoRequest(session *PlayerSession, pack *protocol.Packet) bool {
	packData := pack.Data.(*protocol.C2SLogin)

	golog.Debug("AccountLogin", "DoRequest", "Test login", "loginId", packData.LoginId, "ts", packData.Ts)
	//fmt.Println("do account login ! loginId : ", packData.LoginId, ", ts :", packData.Ts)

	AccessPostgresql(driver.PGPool, packData.LoginId)
	//	golog.Info("AccountLogin", "DoRequest", "Access postgres over", "loginId", packData.LoginId, "ts", packData.Ts)

	AccessMongo(driver.Mgo, packData.LoginId)
	//golog.Info("AccountLogin", "DoRequest", "Access mongo over", "loginId", packData.LoginId, "ts", packData.Ts)

	AccessRedis(packData.LoginId, packData.Ts)
	//      golog.Info("AccountLogin", "DoRequest", "Access redis over", "loginId", packData.LoginId, "ts", packData.Ts, "cnt", rediscount)

	session.SendData([]byte("from account login"))
	return true
}
