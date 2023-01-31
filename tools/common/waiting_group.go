package common

import (
	"sync"
)

type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

type WaitGroupMax struct {
	max   int
	count int
	sync.WaitGroup
	m sync.Mutex
	c *sync.Cond
}

func NewWaitingGroup(max int) WaitGroup {
	g := &WaitGroupMax{
		max: max,
	}
	g.c = sync.NewCond(&c.m)
	return g
}

func (wg *WaitGroupMax) Add(delta int) {
	wg.m.Lock()
	defer wg.m.Unlock()
	for wg.count >= wg.max {
		wg.c.Wait()
	}
	wg.count += delta
	wg.WaitGroup.Add(delta)
}

func (wg *WaitGroupMax) Done() {
	wg.m.Lock()
	defer wg.m.Unlock()
	wg.count--
	wg.c.Signal()
	wg.WaitGroup.Done()
}
