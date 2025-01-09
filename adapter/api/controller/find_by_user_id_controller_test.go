package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"projeto-final/core/domain"
	"projeto-final/core/erros"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindByUserIdControllerTest(t *testing.T) {
	uc := FindByUserIdControllerMock{}

	tt := []struct {
		name               string
		input              string
		mockUseCaseSetup   func()
		expectedStatusCode int
	}{
		{
			name:               "erro de url",
			input:              "/v1/usuario?userId=",
			mockUseCaseSetup:   func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "erro not found",
			input: "/v1/user?userId=10",
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&domain.User{}, erros.NewNotFoundErr("User", "10")).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:  "sucesso",
			input: "/v1/usuario?userId=1",
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&domain.User{
					Id:    "1",
					Email: "teste@gmail.com",
					Idade: 26,
				}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			c := NewFindByUserIdController(&uc)

			r := httptest.NewRequest("GET", sc.input, &bytes.Reader{})
			w := httptest.NewRecorder()

			sc.mockUseCaseSetup()
			c.Execute(w, *r)

			assert.Equal(t, sc.expectedStatusCode, w.Code)

		})
	}
}
