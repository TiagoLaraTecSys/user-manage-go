package controller

import (
	"net/http"
	"projeto-final/adapter/api/handler"
	"projeto-final/adapter/api/response"
	"projeto-final/core/erros"
	"projeto-final/core/usecase"
	"projeto-final/core/usecase/input"
)

type FindByUserIdController struct {
	uc usecase.FindByUserId
}

func NewFindByUserIdController(uc usecase.FindByUserId) *FindByUserIdController {
	return &FindByUserIdController{uc: uc}
}

func (c *FindByUserIdController) Execute(w http.ResponseWriter, r http.Request) {

	userid := r.URL.Query().Get("userId")

	if userid == "" {
		handler.HandleError(w, erros.NewInvalidRequestErr())
		return
	}

	i := &input.FindByIdInput{Id: userid}
	ctx := r.Context()

	user, err := c.uc.Execute(&ctx, i)

	if err != nil {
		handler.HandleError(w, err)
		return
	}

	//logger.Info("Response body", &user)
	response.NewSucess(http.StatusOK, user).Send(w)
}
