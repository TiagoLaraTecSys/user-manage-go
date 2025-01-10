package controller

import (
	"net/http"
	"projeto-final/adapter/api/handler"
	"projeto-final/adapter/api/response"
	"projeto-final/core/usecase"
	"projeto-final/core/usecase/input"
	"projeto-final/infrastructure/logger"
	"strconv"
)

type FindAllUsersController struct {
	uc usecase.FindAllUsers
}

func NewFindAllUsersController(uc usecase.FindAllUsers) *FindAllUsersController {
	return &FindAllUsersController{uc: uc}
}

func (c *FindAllUsersController) Execute(w http.ResponseWriter, r http.Request) {

	ctx := r.Context()
	page, _ := strconv.Atoi(r.URL.Query().Get("Page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	logger.Info("Valor de limit: %s", limit)
	i := &input.PaginationInput{Page: page, Limit: limit}

	users, err := c.uc.Execute(&ctx, i)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response.NewSucess(http.StatusOK, users).Send(w)
}
