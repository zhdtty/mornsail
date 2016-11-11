package logic

import (
//	"fmt"
)

type Hero struct {
	Id     int32
	Level  int32
	Skills map[int32]int32
}

func NewHero(id int32, level int32) *Hero {
	hero := &Hero{
		Id:     id,
		Level:  level,
		Skills: make(map[int32]int32),
	}
	return hero
}
