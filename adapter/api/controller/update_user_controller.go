package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"projeto-final/adapter/api/handler"
	"projeto-final/adapter/api/response"
	"projeto-final/core/erros"
	"projeto-final/core/usecase"
	"projeto-final/core/usecase/input"
	"strconv"
)

type UpdateUserController struct {
	uc usecase.UpdateUser
}

func NewUpdateUserController(uc usecase.UpdateUser) *UpdateUserController {
	return &UpdateUserController{uc: uc}
}

func (c *UpdateUserController) Execute(w http.ResponseWriter, r http.Request) {

	userid := r.URL.Query().Get("userId")

	if userid == "" {
		handler.HandleError(w, erros.NewInvalidRequestErr())
		return
	}

	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	var i input.SaveUser
	if err := json.Unmarshal(jsonBody, &i); err != nil {
		handler.HandleError(w, err)
		return
	}

	ctx := r.Context()
	id, _ := strconv.Atoi(userid)
	u, err := c.uc.Execute(&ctx, &i, id)

	if err != nil {
		handler.HandleError(w, err)
	}

	response.NewSucess(http.StatusOK, u).Send(w)
}
