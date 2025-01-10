package handler

import (
	"errors"
	"net/http"
	"projeto-final/adapter/api/response"
	"projeto-final/core/erros"
)

func HandleError(w http.ResponseWriter, err error) {
	var status int

	switch {
	case errors.As(err, &erros.InvalidRequestErr{}), errors.As(err, &erros.NotUniqueError{}):
		status = http.StatusBadRequest
	case errors.As(err, &erros.NotFoundErr{}):
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
	}
	response.NewError(status, err).Send(w)
}
