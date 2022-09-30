package dto

import (
	"errors"
	"sync"
)

type Map struct {
	data map[int]string
	m    *sync.RWMutex
}

func NewMap() *Map {
	return &Map{
		data: map[int]string{},
		m:    new(sync.RWMutex),
	}
}

func (m *Map) Get(i int) (string, bool) {
	m.m.RLock()
	defer m.m.RUnlock()
	v, ok := m.data[i]
	return v, ok
}

func (m *Map) Set(i int, v string) {
	m.m.Lock()
	defer m.m.Unlock()
	m.data[i] = v
}

func (m *Map) Delete(i int) error {
	m.m.Lock()
	defer m.m.Unlock()
	if _, ok := m.data[i]; ok {
		delete(m.data, i)
		return nil
	}
	return errors.New("order not found")
}
