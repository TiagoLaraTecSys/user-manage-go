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
		findU  *controller.FindByUserIdController
	}
)

func NewGinEngine(
	router *gin.Engine,
	saveU *controller.SaveController,
	findU *controller.FindByUserIdController,
) *ginEngine {
	return &ginEngine{router: router, saveU: saveU, findU: findU}
}

func (e *ginEngine) SetAppHandlers() {
	e.router.GET("/v1/user", e.getByUserId())
	e.router.POST("/v1/user", e.save())
}

func (e *ginEngine) getByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("userId", ctx.Param("userId"))
		ctx.Request.URL.RawQuery = query.Encode()
		e.findU.Execute(ctx.Writer, *ctx.Request)
	}
}

func (e *ginEngine) save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.saveU.Execute(ctx.Writer, *ctx.Request)
	}
}

func (e *ginEngine) GetRouter() *gin.Engine {
	return e.router
}
