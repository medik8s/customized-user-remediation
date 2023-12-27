package script

import "sync"

type Manager interface {
	SetScript(script string)
	GetScript() string
}

func NewManager() Manager {
	return &manager{}
}

type manager struct {
	script string
	sync.RWMutex
}

func (m *manager) SetScript(newScript string) {
	m.Lock()
	defer m.Unlock()
	m.script = newScript
}

func (m *manager) GetScript() string {
	m.RLock()
	defer m.RUnlock()
	return m.script
}
