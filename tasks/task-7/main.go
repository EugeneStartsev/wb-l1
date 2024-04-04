package main

import (
	"fmt"
	"sync"
)

type safeMap struct {
	rw   sync.Mutex
	data map[string]any
}

func newSafeMap() safeMap {
	return safeMap{
		rw:   sync.Mutex{},
		data: make(map[string]any),
	}
}

func (m *safeMap) set(key string, val any) {
	m.rw.Lock()
	defer m.rw.Unlock()

	m.data[key] = val
}

func (m *safeMap) get(key string) (any, bool) {
	m.rw.Lock()
	defer m.rw.Unlock()

	v, ok := m.data[key]

	return v, ok
}

func (m *safeMap) remove(key string) {
	m.rw.Lock()
	defer m.rw.Unlock()

	delete(m.data, key)
}

func main() {
	sm := newSafeMap()

	sm.set("first", "asdhkjas")
	sm.set("second", 2)
	sm.set("third", true)

	fmt.Println(sm.data)

	v, _ := sm.get("first")
	fmt.Println(v)

	sm.remove("second")
	_, ok := sm.get("second")
	fmt.Println(ok)
}
