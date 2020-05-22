package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

// Service base struct
// consist of db using gorm ORM
// gorm docs https://pkg.go.dev/github.com/jinzhu/gorm?tab=doc
type Service struct {
	db   *gorm.DB
	repo *Repository
}

// NewService initiate Service
func NewService() *Service {
	r := &Service{}

	r.repo = NewRepository()
	return r
}

// Resolve resolve all of todo
func (s *Service) Resolve() ([]Todo, error) {
	return s.repo.ResolveAll()
}

// ResolveByID resolve todo by id
func (s *Service) ResolveByID(id uuid.UUID) (Todo, error) {
	t, err := s.repo.ResolveByID(id)
	return t, err
}

// Create create todo and generate id
func (s *Service) Create(i Input) (Todo, error) {
	t := i.toTodo()
	err := s.repo.Store(t)
	return t, err
}

// Update update Todo if id is exist
func (s *Service) Update(id uuid.UUID, u Update) (Todo, error) {
	t, err := s.repo.ResolveByID(id)
	if err != nil {
		return t, err
	}

	t.Update(u)
	err = s.repo.Update(id, t)
	return t, err
}

// Remove delete todo data
func (s *Service) Remove(id uuid.UUID) error {
	_, err := s.repo.ResolveByID(id)
	if err != nil {
		return err
	}

	return s.repo.Remove(id)
}
