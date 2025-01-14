package usecase

import (
	"context"
	"projeto-final/core/repository"
)

type (
	DeleteUser interface {
		Execute(ctx *context.Context, Id int) error
	}
	deleteUser struct {
		repo repository.UserRepository
	}
)

func NewDeleteUser(repo repository.UserRepository) DeleteUser {
	return &deleteUser{repo: repo}
}

func (uc *deleteUser) Execute(ctx *context.Context, Id int) error {

	result, err := uc.repo.GetById(ctx, Id)

	if err != nil {
		return err
	}

	err = uc.repo.DeleteUser(ctx, result.Id)

	if err != nil {
		return err
	}

	return nil
}
