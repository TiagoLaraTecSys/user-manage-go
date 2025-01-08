package router

import (
	"projeto-final/adapter/api/controller"

	"github.com/gin-gonic/gin"
)

type (
	GinRouter interface {
		SetAppHandlers()
		GetRouter() *gin.Engine
	}

	ginEngine struct {
		router *gin.Engine
		saveU  *controller.SaveController
	}
)

func NewGinEngine(
	router *gin.Engine,
	saveU *controller.SaveController,
) *ginEngine {
	return &ginEngine{router: router, saveU: saveU}
}

func (e *ginEngine) SetAppHandlers() {
	e.router.POST("/v1/user", e.save())
}

func (e *ginEngine) save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.saveU.Execute(ctx.Writer, *ctx.Request)
	}
}

func (e *ginEngine) GetRouter() *gin.Engine {
	return e.router
}
