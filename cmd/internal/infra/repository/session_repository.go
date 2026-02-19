package repository

import (
	"errors"
	"sync"

	"github.com/lhupalo/mcp-network-sentinel-go/cmd/internal/domain"
)

type MemorySessionRepository struct {
	sessionsList map[string]*domain.Session
	sync.RWMutex // nao para proteger a sessao em si, mas o mapa
}

func NewMemorySessionRepository() *MemorySessionRepository {
	return &MemorySessionRepository{sessionsList: make(map[string]*domain.Session, 20)}
}

// Implementaçao dos metodos

func (m *MemorySessionRepository) Save(session *domain.Session) error {

	if session.ID == "" {
		return errors.New("session ID is required")
	}

	m.Lock()
	defer m.Unlock()

	m.sessionsList[session.ID] = session

	return nil
}

func (m *MemorySessionRepository) GetById(id string) (*domain.Session, error) {

	m.RLock()
	defer m.RUnlock()

	if id == "" {
		return nil, errors.New("session id is required")
	}

	session, found := m.sessionsList[id]

	if found == false {
		return nil, errors.New("session not found")
	}

	return session, nil
}

func (m *MemorySessionRepository) Delete(id string) error {

	m.Lock()
	defer m.Unlock()

	if id == "" {
		return errors.New("session id is required")
	}

	delete(m.sessionsList, id)

	return nil
}

func (m *MemorySessionRepository) GetAll() ([]*domain.Session, error) {

	m.RLock()
	defer m.RUnlock()

	sessions := make([]*domain.Session, 0, len(m.sessionsList))

	for _, session := range m.sessionsList {
		sessions = append(sessions, session)
	}

	return sessions, nil
}
