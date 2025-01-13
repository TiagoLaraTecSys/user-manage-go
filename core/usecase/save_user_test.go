package usecase

import (
	"context"
	"fmt"
	"projeto-final/core/domain"
	"projeto-final/core/usecase/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveUser(t *testing.T) {

	inPadrao := input.SaveUser{
		Email: "laratecsys@gmail.com",
		Idade: 26,
	}

	outPadrao := domain.User{
		Id:    1,
		Email: "laratecsys@gmail.com",
		Idade: 26,
	}

	tt := []struct {
		name     string
		input    input.SaveUser
		expected domain.User
		err      error
	}{
		{
			name:     "Sucesso",
			input:    inPadrao,
			expected: outPadrao,
			err:      nil,
		},
		{
			name:     "Error",
			input:    inPadrao,
			expected: domain.User{},
			err:      fmt.Errorf("Database error"),
		},
	}

	for _, su := range tt {
		t.Run(su.name, func(t *testing.T) {

			//Cria a estrtura de teste
			r := NewMockDatabase()
			uc := NewSaveUser(r)

			// Estabelece comportamento do mock
			r.On("Add", mock.Anything, mock.Anything).Return(su.expected, su.err)

			//Execução
			ctx := context.TODO()
			out, err := uc.Execute(&ctx, &su.input)

			//Validação
			assert.Equal(t, su.expected, *out)
			assert.Equal(t, su.err, err)
		})
	}
}
