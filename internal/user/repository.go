package user

import "sync"

type Repository struct {
	mu sync.RWMutex
	users map[string]User
}

func NewRepository() *Repository {
	return &Repository{
		users: make(map[string]User),
	}
}