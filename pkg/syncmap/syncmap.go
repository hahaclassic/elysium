package syncmap

import (
	"sync"
)

type Map struct {
	mx   *sync.RWMutex
	data map[any]any
}

func NewMap() *Map {
	return &Map{
		data: make(map[any]any),
		mx:   &sync.RWMutex{},
	}
}

func (m *Map) Load(key any) (value any, ok bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.data[key]

	return val, ok
}

func (m *Map) Store(key any, value any) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.data[key] = value
}
