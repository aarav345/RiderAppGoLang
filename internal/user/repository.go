package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	_, err := r.pool.Exec(ctx, `INSERT INTO users (first_name, last_name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*User, error)
func (r *repository) FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (*User, error)
func (r *repository) FindUserBySocialID(ctx context.Context, socialID string) (*User, error)
func (r *repository) UpdateUser(ctx context.Context, user *User) error
