package usecase

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/repository"
)

type (
	FindAllUsers interface {
		Execute(ctx *context.Context) (*[]domain.User, error)
	}
	findAllUsers struct {
		repo repository.UserRepository
	}
)

func NewFindAllUsers(repo repository.UserRepository) FindAllUsers {
	return &findAllUsers{repo: repo}
}

func (uc *findAllUsers) Execute(ctx *context.Context) (*[]domain.User, error) {

	out, err := uc.repo.GetUsers(ctx)

	if err != nil {
		return &[]domain.User{}, err
	}

	return &out, nil
}
