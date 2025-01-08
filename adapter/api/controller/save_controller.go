package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"projeto-final/adapter/api/handler"
	"projeto-final/adapter/api/response"
	"projeto-final/core/usecase"
	"projeto-final/core/usecase/input"
)

type SaveController struct {
	uc usecase.SaveUser
}

func NewSaveController(uc usecase.SaveUser) *SaveController {

	return &SaveController{uc: uc}
}

func (c *SaveController) Execute(w http.ResponseWriter, r http.Request) {
	ctx := r.Context()

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

	u, err := c.uc.Execute(&ctx, &i)
	if err != nil {
		handler.HandleError(w, err)
		return
	}

	response.NewSucess(http.StatusOK, u).Send(w)

}
