package usecase

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/repository"
	"projeto-final/core/usecase/input"
)

type (
	FindAllUsers interface {
		Execute(ctx *context.Context, i *input.PaginationInput) (*domain.Data, error)
	}
	findAllUsers struct {
		repo repository.UserRepository
	}
)

func NewFindAllUsers(repo repository.UserRepository) FindAllUsers {
	return &findAllUsers{repo: repo}
}

func (uc *findAllUsers) Execute(ctx *context.Context, i *input.PaginationInput) (*domain.Data, error) {

	out, err := uc.repo.GetUsers(ctx, i)

	if err != nil {
		return &domain.Data{}, err
	}

	return &out, nil
}
