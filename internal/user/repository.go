package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) FindAll() ([]User, error) {
	rows, err := r.pool.Query(context.Background(), "select id,name,email,created_at,updated_at from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *Repository) FindById(id string) (User, error) {
	var u User
	query := "select id,name,email,created_at,updated_at from users where id = $1"
	row := r.pool.QueryRow(context.Background(), query, id)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *Repository) Insert(name, email string) (User, error) {
	newUser := User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := "insert into users (id, name, email, created_at, updated_at) values ($1, $2, $3, $4, $5)"
	_, err := r.pool.Exec(context.Background(), query, newUser.ID, newUser.Name, newUser.Email, newUser.CreatedAt, newUser.UpdatedAt)
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (r *Repository) Update(id, name, email string) (User, error) {
	var u User
	query := "update users set name = $1, email = $2, updated_at = $3 where id = $4 returning id, name, email, created_at, updated_at"
	row := r.pool.QueryRow(context.Background(), query, name, email, time.Now(), id)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (r *Repository) Delete(id string) error {
	query := "delete from users where id = $1"
	_, err := r.pool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
