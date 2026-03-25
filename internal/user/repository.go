package user

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Repository struct {
	mu sync.RWMutex
	users map[string]User
}

func NewRepository() *Repository {
	return &Repository{
		users: make(map[string]User),
	}
}

func (r *Repository) FindAll() []User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]User, 0, len(r.users))

	for _, user := range r.users {
		users = append(users, user)
	}

	return users
}

func (r *Repository) Insert(name, email string) (User, error){
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users{
		if user.Email == email {
			return User{}, errors.New("email already exists")
		}
	}

	newUser := User{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	r.users[newUser.ID] = newUser

	return newUser, nil
}