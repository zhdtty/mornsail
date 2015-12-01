package servlet

import (
	//	"fmt"
	"sync"
)

type PlayerSessionManager struct {
	sync.Mutex
	onlines  map[int64]*PlayerSession
	offlines map[int64]*PlayerSession
	caches   map[int64]*PlayerSession
}

func NewPlayerSessionManager() *PlayerSessionManager {
	psm := &PlayerSessionManager{
		onlines:  make(map[int64]*PlayerSession),
		offlines: make(map[int64]*PlayerSession),
		caches:   make(map[int64]*PlayerSession),
	}
	return psm
}

func (psm *PlayerSessionManager) OnlinePlayer(ps *PlayerSession) {
	psm.Lock()
	defer psm.Unlock()
	psm.onlines[ps.Id] = ps
	delete(psm.offlines, ps.Id)
	delete(psm.caches, ps.Id)
}

func (psm *PlayerSessionManager) OfflinePlayer(ps *PlayerSession) {
	psm.Lock()
	defer psm.Unlock()
	delete(psm.onlines, ps.Id)
	delete(psm.caches, ps.Id)
	psm.offlines[ps.Id] = ps
}

func (psm *PlayerSessionManager) CachePlayer(ps *PlayerSession) {
	psm.Lock()
	defer psm.Unlock()
	delete(psm.onlines, ps.Id)
	delete(psm.offlines, ps.Id)
	psm.caches[ps.Id] = ps
}
