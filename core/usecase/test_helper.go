package usecase

import (
	"context"
	"fmt"
	"projeto-final/core/domain"
	"projeto-final/core/usecase/input"

	"github.com/stretchr/testify/mock"
)

type databaseMock struct {
	mock.Mock
}

var databaseErr = fmt.Errorf("Database error")

func NewMockDatabase() *databaseMock {
	return &databaseMock{}
}

func (d *databaseMock) Add(ctx *context.Context, p *domain.User) (domain.User, error) {
	args := d.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

func (d *databaseMock) GetById(ctx *context.Context, uId string) (domain.User, error) {
	args := d.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

func (d *databaseMock) GetUsers(ctx *context.Context, i *input.PaginationInput) ([]domain.User, error) {
	args := d.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}
