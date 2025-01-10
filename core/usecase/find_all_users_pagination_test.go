package usecase

import (
	"context"
	"projeto-final/core/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAllUsers(t *testing.T) {

	resDef := []domain.User{
		{
			Email: "teste1@gmail.com",
			Idade: 26,
			Id:    "1",
		},
		{
			Email: "teste2@gmail.com",
			Idade: 26,
			Id:    "2",
		},
	}

	tt := []struct {
		name     string
		expected []domain.User
		err      error
	}{
		{
			name:     "Sucesso",
			expected: resDef,
			err:      nil,
		},
		{
			name:     "Error",
			expected: []domain.User{},
			err:      databaseErr,
		},
	}

	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			r := NewMockDatabase()
			uc := NewFindAllUsers(r)

			r.On("GetUsers", mock.Anything).Return(sc.expected, sc.err)

			ctx := context.TODO()
			out, err := uc.Execute(&ctx)

			assert.Equal(t, sc.expected, *out)
			assert.Equal(t, sc.err, err)
		})
	}
}
