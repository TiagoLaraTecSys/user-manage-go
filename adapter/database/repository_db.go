package database

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/usecase/input"
)

type DbUser interface {
	Add(ctx *context.Context, user *domain.User) (domain.User, error)
	GetById(ctx *context.Context, id string) (domain.User, error)
	GetUsers(ctx *context.Context, i *input.PaginationInput) (domain.Data, error)
}
