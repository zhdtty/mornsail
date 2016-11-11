package logic

import (
	//	"fmt"
	"sync"
)

type HeroBag struct {
	sync.RWMutex
	Heros map[int32]*Hero
}

func NewHeroBag() *HeroBag {
	heroBag := &HeroBag{
		Heros: make(map[int32]*Hero),
	}
	return heroBag
}

func (hb *HeroBag) GetHero(id int32) *Hero {
	hb.RLock()
	defer hb.RUnlock()
	if hero, ok := hb.Heros[id]; ok {
		return hero
	}
	return nil
}
