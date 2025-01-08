package repository

import (
	"context"
	"projeto-final/core/domain"
)

type UserRepository interface {
	Add(ctx *context.Context, user *domain.User) (domain.User, error)
	GetById(ctx *context.Context, id string) (domain.User, error)
	GetUsers(ctx *context.Context) ([]domain.User, error)
}
