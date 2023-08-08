package datastructures

import "sync"

type MapLock struct {
	m    map[string]*sync.Mutex
	lock *sync.Mutex
}

func NewMapLock() *MapLock {
	return &MapLock{
		m:    make(map[string]*sync.Mutex),
		lock: &sync.Mutex{},
	}
}

func (m *MapLock) Lock(key string) {
	m.lock.Lock()

	lock, ok := m.m[key]
	if !ok {
		lock = &sync.Mutex{}
		m.m[key] = lock
	}

	m.lock.Unlock()
	lock.Lock()
}

func (m *MapLock) Unlock(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	lock, ok := m.m[key]
	if !ok {
		return
	}

	lock.Unlock()
	delete(m.m, key)
}
