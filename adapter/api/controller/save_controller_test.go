package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"projeto-final/core/domain"
	"projeto-final/core/usecase/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveController(t *testing.T) {
	var uc SaveControllerMock

	inputPadrao := input.SaveUser{
		Email: "teste@gmail.com",
		Idade: 26,
	}

	outputPadrao := domain.User{
		Id:    1,
		Email: "teste@gmail.com",
		Idade: 26,
	}

	tt := []struct {
		name               string
		requestBody        any
		mockUseCaseSetup   func()
		expectedStatusCode int
	}{
		{
			name:               "body read error",
			requestBody:        nil,
			mockUseCaseSetup:   func() {},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "Invalid Json",
			requestBody:        "{teste: 01",
			mockUseCaseSetup:   func() {},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "use case Error",
			requestBody: inputPadrao,
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&domain.User{}, errors.New("database error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "success",
			requestBody: inputPadrao,
			mockUseCaseSetup: func() {
				uc.On("Execute", mock.Anything, mock.Anything).Return(&outputPadrao, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, sc := range tt {
		t.Run(sc.name, func(t *testing.T) {
			c := NewSaveController(&uc)

			var body io.Reader

			if sc.requestBody != nil {
				jsonData, _ := json.Marshal(sc.requestBody)
				body = bytes.NewBuffer(jsonData)
			} else {
				body = &ErrorReader{}
			}

			req := httptest.NewRequest(http.MethodPost, "/v1/user", body)
			w := httptest.NewRecorder()

			sc.mockUseCaseSetup()

			c.Execute(w, *req)
			assert.Equal(t, sc.expectedStatusCode, w.Code)
		})
	}
}
