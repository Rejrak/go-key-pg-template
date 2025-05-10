package user

import (
	"be/internal/database/db"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	KcID      uuid.UUID
	FirstName string
	LastName  string
	Nickname  string
	Admin     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserWithPlans struct {
	ID        uuid.UUID
	KcID      uuid.UUID
	FirstName string
	LastName  string
	Nickname  string
	Admin     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository struct {
	DB *sql.DB
}

func NewRepository() *Repository {
	return &Repository{DB: db.DB.LD}
}

func (r *Repository) FindByID(ctx context.Context, userID string) (*UserWithPlans, error) {
	var user UserWithPlans
	return &user, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]User, error) {

	var users []User

	return users, nil
}

func (r *Repository) SaveUser(ctx context.Context, user UserWithPlans) (*UserWithPlans, error) {

	return &user, nil
}

func (r *Repository) DeleteUser(ctx context.Context, userID string) error {

	return nil
}
