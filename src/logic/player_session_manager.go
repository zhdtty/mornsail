package logic

import (
//	"fmt"
)

type PlayerSessionManager struct {
	//	sync.Mutex
	sync     chan bool
	onlines  map[int64]*PlayerSession
	offlines map[int64]*PlayerSession
	caches   map[int64]*PlayerSession
}

func NewPlayerSessionManager() *PlayerSessionManager {
	psm := &PlayerSessionManager{
		sync:     make(chan bool, 1),
		onlines:  make(map[int64]*PlayerSession),
		offlines: make(map[int64]*PlayerSession),
		caches:   make(map[int64]*PlayerSession),
	}
	return psm
}

func (psm *PlayerSessionManager) OnlinePlayer(ps *PlayerSession) {
	//	psm.Lock()
	//	defer psm.Unlock()
	psm.sync <- true
	defer func() { <-psm.sync }()
	psm.onlines[ps.Id] = ps
	delete(psm.offlines, ps.Id)
	delete(psm.caches, ps.Id)
}

func (psm *PlayerSessionManager) OfflinePlayer(ps *PlayerSession) {
	//	psm.Lock()
	//	defer psm.Unlock()
	psm.sync <- true
	defer func() { <-psm.sync }()
	delete(psm.onlines, ps.Id)
	delete(psm.caches, ps.Id)
	psm.offlines[ps.Id] = ps
}

func (psm *PlayerSessionManager) CachePlayer(ps *PlayerSession) {
	//	psm.Lock()
	//	defer psm.Unlock()
	psm.sync <- true
	defer func() { <-psm.sync }()
	delete(psm.onlines, ps.Id)
	delete(psm.offlines, ps.Id)
	psm.caches[ps.Id] = ps
}

func (psm *PlayerSessionManager) Broadcast(b []byte) {
	//	psm.Lock()
	//	defer psm.Unlock()
	psm.sync <- true
	defer func() { <-psm.sync }()
	//Need not copy to new map, gc more busy, and also SendData just to writeCache
	for _, v := range psm.onlines {
		v.SendData(b)
	}
}

var G_playerSessionMgr *PlayerSessionManager
