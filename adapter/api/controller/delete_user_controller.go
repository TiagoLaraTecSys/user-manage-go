package controller

import (
	"net/http"
	"projeto-final/adapter/api/handler"
	"projeto-final/adapter/api/response"
	"projeto-final/core/erros"
	"projeto-final/core/usecase"
	"strconv"
)

type DeleteUserController struct {
	uc usecase.DeleteUser
}

func NewDeleteUserController(uc usecase.DeleteUser) *DeleteUserController {
	return &DeleteUserController{uc: uc}
}

func (c *DeleteUserController) Execute(w http.ResponseWriter, r http.Request) {
	userid := r.URL.Query().Get("userId")

	if userid == "" {
		handler.HandleError(w, erros.NewInvalidRequestErr())
		return
	}

	ctx := r.Context()
	id, _ := strconv.Atoi(userid)
	err := c.uc.Execute(&ctx, id)

	if err != nil {
		handler.HandleError(w, err)
	}

	response.NewSucess(http.StatusNoContent, nil)
}
