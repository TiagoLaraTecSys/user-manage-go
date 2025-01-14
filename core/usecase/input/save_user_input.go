package input

import (
	"fmt"
	"projeto-final/core/erros"

	"github.com/go-playground/validator/v10"
)

type SaveUser struct {
	Name  string `json:"Name" validate:"max=100"`
	Email string `json:"Email" validate:"required,email"`
	Idade int    `json:"Idade" validate:"required"`
}

func (s *SaveUser) ValidateRequestBody() error {

	validate := validator.New()

	err := validate.Struct(s)

	var campos []string

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			campos = append(campos, fmt.Sprintf("Campo: %s | Erro: %s:%s\n", e.Field(), e.ActualTag(), e.Param()))
		}
		return erros.NewInvalidRequestErr(campos...)
	}

	return nil
}
