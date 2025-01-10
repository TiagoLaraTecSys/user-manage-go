package usecase

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/repository"
	"projeto-final/core/usecase/input"
)

type (
	FindAllUsers interface {
		Execute(ctx *context.Context, i *input.PaginationInput) (*[]domain.User, error)
	}
	findAllUsers struct {
		repo repository.UserRepository
	}
)

func NewFindAllUsers(repo repository.UserRepository) FindAllUsers {
	return &findAllUsers{repo: repo}
}

func (uc *findAllUsers) Execute(ctx *context.Context, i *input.PaginationInput) (*[]domain.User, error) {

	out, err := uc.repo.GetUsers(ctx, i)

	if err != nil {
		return &[]domain.User{}, err
	}

	return &out, nil
}
