package usecase

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/repository"
	"projeto-final/core/usecase/input"
)

type (
	SaveUser interface {
		Execute(ctx *context.Context, u *input.SaveUser) (*domain.User, error)
	}
	saveUser struct {
		repo repository.UserRepository
	}
)

func NewSaveUser(repo repository.UserRepository) SaveUser {
	return &saveUser{repo: repo}
}

func (s *saveUser) Execute(ctx *context.Context, u *input.SaveUser) (*domain.User, error) {

	user, err := domain.NewUser(
		domain.WithEmail(u.Email),
		domain.WithIdade(u.Idade),
		domain.WithName(u.Name),
	)

	if err != nil {
		return &domain.User{}, err
	}

	out, err := s.repo.Add(ctx, user)

	if err != nil {
		return &domain.User{}, err
	}

	return &out, nil
}
