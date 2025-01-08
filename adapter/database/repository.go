package database

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/repository"
)

type UserRepository struct {
	db DbUser
}

func NewUserRepository(db DbUser) *UserRepository {
	return &UserRepository{db: db}
}

var _ repository.UserRepository = (*UserRepository)(nil)

func (r *UserRepository) Add(ctx *context.Context, user *domain.User) (domain.User, error) {

	u, err := r.db.Add(ctx, user)
	if err != nil {
		return u, err
	}

	return u, err
}

func (r *UserRepository) GetById(ctx *context.Context, id string) (domain.User, error) {

	u, err := r.db.GetById(ctx, id)
	if err != nil {
		return u, err
	}

	return u, err
}

func (r *UserRepository) GetUsers(ctx *context.Context) ([]domain.User, error) {

	u, err := r.db.GetUsers(ctx)
	if err != nil {
		return u, err
	}

	return u, err
}
