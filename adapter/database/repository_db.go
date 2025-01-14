package database

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/usecase/input"
)

type DbUser interface {
	Add(ctx *context.Context, user *domain.User) (domain.User, error)
	Update(ctx *context.Context, user *domain.User) (domain.User, error)
	GetById(ctx *context.Context, id int) (domain.User, error)
	GetUsers(ctx *context.Context, i *input.PaginationInput) (domain.Data, error)
	DeleteUser(ctx *context.Context, Id int) error
}
