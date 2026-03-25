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

func (r *Repository) FindById(id string) (User, error){
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	} 

	return user, nil
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

func (r *Repository) Update(id,name,email string) (User, error){
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists{
		return User{}, errors.New("user not found")
	}

	user.Name = name
	user.Email = email
	user.UpdatedAt = time.Now()

	r.users[id] = user

	return user, nil
}

func (r * Repository) Delete(id string) error{
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}