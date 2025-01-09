package usecase

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/repository"
	"projeto-final/core/usecase/input"
)

type (
	FindByUserId interface {
		Execute(ctx *context.Context, i *input.FindByIdInput) (*domain.User, error)
	}
	findByUserId struct {
		repo repository.UserRepository
	}
)

func NewFindByUserId(repo repository.UserRepository) FindByUserId {
	return &findByUserId{repo: repo}
}

func (r *findByUserId) Execute(ctx *context.Context, i *input.FindByIdInput) (*domain.User, error) {
	out, err := r.repo.GetById(ctx, i.Id)

	if err != nil {
		return &domain.User{}, err
	}

	return &out, nil
}
