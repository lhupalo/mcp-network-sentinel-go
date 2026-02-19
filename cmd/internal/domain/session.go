package domain

import "time"

// o channel no go é thread-safe por natureza

type Session struct {
	ID          string
	CreatedAt   time.Time
	MessageChan chan interface{}
}

// repositorio nao cria, apenas persiste. por isso nao tem create
type SessionRepository interface {
	Save(session *Session) error
	GetById(id string) (*Session, error)
	Delete(id string) error
	GetAll() ([]*Session, error)
}
