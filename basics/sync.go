package main

import (
	"fmt"
	"sync"
	"time"
)

type cityMap struct {
	data map[string]uint32
	mu   sync.RWMutex
}

func initMap() *cityMap {
	now := now()
	data := map[string]uint32{
		"kyiv": now, "kharkiv": now, "lviv": now, "odessa": now,
	}
	return &cityMap{data: data}
}

func now() uint32 {
	return uint32(time.Now().Unix())
}

func (c *cityMap) update(city string, timestamp uint32) {
	c.mu.Lock()
	c.data[city] = timestamp
	c.mu.Unlock()
}

func (c *cityMap) get(city string) uint32 {
	c.mu.RLock()
	var value = c.data[city]
	c.mu.RUnlock()

	return value
}

func runSyncExample() {
	cm := initMap()
	fmt.Println(cm)
	cm.update("kyiv", now() + 54321)
	fmt.Println(cm)
}
