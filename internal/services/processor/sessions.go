package service

import "sync"

type Session struct {
	lastOperation Operation
	currOperation Operation
	lastInterface UserInterface
	info          any
}

type SessionMap struct {
	data map[int64]*Session
	mx   *sync.RWMutex
}

func NewMap() *SessionMap {
	return &SessionMap{
		data: make(map[int64]*Session),
		mx:   &sync.RWMutex{},
	}
}

func (m *SessionMap) Load(userID int64) (*Session, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.data[userID]

	return val, ok
}

func (m *SessionMap) Store(userID int64, session *Session) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.data[userID] = session
}

func (m *SessionMap) Delete(userID int64) {
	m.mx.Lock()
	defer m.mx.Unlock()
	delete(m.data, userID)
}
