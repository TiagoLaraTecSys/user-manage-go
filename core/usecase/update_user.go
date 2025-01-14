package usecase

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/repository"
	"projeto-final/core/usecase/input"
)

type (
	UpdateUser interface {
		Execute(ctx *context.Context, i *input.SaveUser, Id int) (*domain.User, error)
	}
	updateUser struct {
		repo repository.UserRepository
	}
)

func NewUpdateUser(repo repository.UserRepository) UpdateUser {
	return &updateUser{repo: repo}
}

func (u *updateUser) Execute(ctx *context.Context, i *input.SaveUser, Id int) (*domain.User, error) {

	user, err := u.repo.GetById(ctx, Id)

	if err != nil {
		return &domain.User{}, err
	}

	userNew, err := domain.NewUser(
		domain.WithEmail(i.Email),
		domain.WithIdade(i.Idade),
		domain.WithName(i.Name),
	)

	if err != nil {
		return &domain.User{}, err
	}

	if userNew.Idade != 0 {
		user.Idade = userNew.Idade
	}

	out, err := u.repo.Add(ctx, &user)

	if err != nil {
		return &domain.User{}, err
	}

	return &out, nil
}
