package team

import "time"

//Model
type Model struct {
	Id          uint
	Name        string
	Description string
	Manager     string
	Team        int64
	CreateAt    time.Time
	UpdateAt    time.Time
}

//interface CRUD

type Storage interface {
	Migrate() error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
