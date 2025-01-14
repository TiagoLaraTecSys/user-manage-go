package usecase

import (
	"context"
	"projeto-final/core/domain"
	"projeto-final/core/usecase/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindByIdExecute(t *testing.T) {

	respDef := domain.User{
		Email: "teste@gmail.com",
		Idade: 26,
		Id:    1,
	}

	tt := []struct {
		name     string
		input    *input.FindByIdInput
		expected domain.User
		err      error
	}{
		{
			name:     "Sucesso",
			input:    &input.FindByIdInput{Id: 1},
			expected: respDef,
			err:      nil,
		},
		{
			name:     "Erro",
			input:    &input.FindByIdInput{Id: 1},
			expected: domain.User{},
			err:      databaseErr,
		},
	}

	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			r := NewMockDatabase()
			uc := NewFindByUserId(r)

			r.On("GetById", mock.Anything, mock.Anything).Return(sc.expected, sc.err)

			ctx := context.TODO()
			out, err := uc.Execute(&ctx, sc.input)

			assert.Equal(t, sc.expected, *out)
			assert.Equal(t, sc.err, err)
		})
	}
}
