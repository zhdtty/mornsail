// example: /go/goexample/timer_fix.go
package timer

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type TimerItem struct {
	nextTime int64
	interval int64
	loop     bool
	f        func()
}
type TimerItemList []*TimerItem

func (list TimerItemList) Len() int {
	return len(list)
}
func (list TimerItemList) Less(i, j int) bool {
	return list[i].nextTime < list[j].nextTime
}
func (list TimerItemList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

type ServerTimer struct {
	sync.RWMutex
	timers TimerItemList
	ticker *time.Ticker
	close  bool
}

func NewServerTimer() *ServerTimer {
	st := &ServerTimer{
		timers: make(TimerItemList, 0),
		ticker: time.NewTicker(time.Millisecond * 1000),
		close:  false,
	}

	go func() {
		for _ = range st.ticker.C {
			if st.close == true {
				st.ticker.Stop()
				break
			}
			st.Run()
		}
	}()
	return st
}

func (st *ServerTimer) Run() {
	st.Lock()
	defer st.Unlock()
	nts := time.Now().Unix()
	resort := false
	for i := 0; i < len(st.timers); i++ {
		if st.timers[i].nextTime > nts {
			break
		}
		go st.timers[i].f()
		if st.timers[i].loop == false {
			st.timers = append(st.timers[:i], st.timers[i+1:]...)
			resort = true
		} else {
			st.timers[i].nextTime += st.timers[i].interval
			resort = true
		}
	}
	if resort == true {
		sort.Sort(st.timers)
	}
}

func (st *ServerTimer) AddDailyTimer(hour, min, sec int, f func(), loop bool) *TimerItem {
	now := time.Now()
	ts := time.Date(now.Year(), now.Month(), now.Day(), hour, min, sec, 0, time.Local).Unix()
	nts := now.Unix()
	nextTime := ts
	if ts < nts {
		nextTime = ts + int64(24*time.Hour/1000000000)
	}
	ti := &TimerItem{
		nextTime: nextTime,
		interval: int64(24 * time.Hour / 1000000000),
		loop:     loop,
		f:        f,
	}
	st.Lock()
	st.timers = append(st.timers, ti)
	sort.Sort(st.timers)
	st.Unlock()

	return ti
}

func (st *ServerTimer) AddHourlyTimer(min, sec int, f func(), loop bool) *TimerItem {
	now := time.Now()
	ts := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), min, sec, 0, time.Local).Unix()
	nts := now.Unix()
	nextTime := ts
	if ts < nts {
		nextTime = ts + int64(time.Hour/1000000000)
	}
	ti := &TimerItem{
		nextTime: nextTime,
		interval: int64(time.Hour / 1000000000),
		loop:     loop,
		f:        f,
	}

	st.Lock()
	st.timers = append(st.timers, ti)
	sort.Sort(st.timers)
	st.Unlock()

	return ti
}

func (st *ServerTimer) AddIntervalTimer(interval int, f func(), loop bool) *TimerItem {
	ti := &TimerItem{
		nextTime: time.Now().Unix() + int64(interval),
		interval: int64(interval),
		loop:     loop,
		f:        f,
	}

	st.Lock()
	st.timers = append(st.timers, ti)
	sort.Sort(st.timers)
	st.Unlock()

	return ti
}

func (st *ServerTimer) DeleteTimer(ti *TimerItem) {
	for i := 0; i < len(st.timers); i++ {
		if st.timers[i] == ti {
			st.timers = append(st.timers[:i], st.timers[i+1:]...)
			break
		}
	}
}

func (st *ServerTimer) PrintTimers() {
	for i := 0; i < len(st.timers); i++ {
		fmt.Println("nexttime ", st.timers[i].nextTime, "interval ", st.timers[i].interval)
	}
}

var SvrTimer *ServerTimer = NewServerTimer()
