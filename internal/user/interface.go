package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByPhoneNumber(ctx context.Context, phoneNumber string) (*User, error)
	FindUserBySocialID(ctx context.Context, socialID string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
}
