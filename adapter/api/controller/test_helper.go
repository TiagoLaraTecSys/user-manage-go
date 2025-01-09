package controller

import (
	"context"
	"fmt"
	"projeto-final/core/domain"
	"projeto-final/core/usecase/input"

	"github.com/stretchr/testify/mock"
)

type (
	SaveControllerMock struct {
		mock.Mock
	}
	FindByUserIdControllerMock struct {
		mock.Mock
	}
	ErrorReader struct{}
)

func (cm *FindByUserIdControllerMock) Execute(ctx *context.Context, i *input.FindByIdInput) (*domain.User, error) {
	args := cm.Called()
	return args.Get(0).(*domain.User), args.Error(1)
}

func (uc *SaveControllerMock) Execute(ctx *context.Context, i *input.SaveUser) (*domain.User, error) {
	args := uc.Called()
	return args.Get(0).(*domain.User), args.Error(1)
}

func (e *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("read error")
}
