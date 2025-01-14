package database

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/repository"
	"projeto-final/core/usecase/input"
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

func (r *UserRepository) Update(ctx *context.Context, user *domain.User) (domain.User, error) {

	u, err := r.db.Update(ctx, user)
	if err != nil {
		return u, err
	}
	return u, err
}

func (r *UserRepository) GetById(ctx *context.Context, id int) (domain.User, error) {

	u, err := r.db.GetById(ctx, id)
	if err != nil {
		return u, err
	}

	return u, err
}

func (r *UserRepository) GetUsers(ctx *context.Context, i *input.PaginationInput) (domain.Data, error) {

	u, err := r.db.GetUsers(ctx, i)
	if err != nil {
		return u, err
	}

	return u, err
}

func (r *UserRepository) DeleteUser(ctx *context.Context, Id int) error {

	err := r.db.DeleteUser(ctx, Id)

	if err != nil {
		return err
	}

	return nil
}
