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
		allU   *controller.FindAllUsersController
		delU   *controller.DeleteUserController
		upU    *controller.UpdateUserController
	}
)

func NewGinEngine(
	router *gin.Engine,
	saveU *controller.SaveController,
	findU *controller.FindByUserIdController,
	allU *controller.FindAllUsersController,
	delU *controller.DeleteUserController,
	upU *controller.UpdateUserController,
) *ginEngine {
	return &ginEngine{router: router, saveU: saveU, findU: findU, allU: allU, delU: delU, upU: upU}
}

func (e *ginEngine) SetAppHandlers() {
	e.router.GET("/v1/user", e.getByUserId())
	e.router.GET("/v1/users", e.getAllUsers())
	e.router.POST("/v1/user", e.save())
	e.router.DELETE("/v1/user", e.deleteUser())
	e.router.PUT("/v1/user", e.updateUser())

}

func (e *ginEngine) updateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("userId", ctx.Param("userId"))
		ctx.Request.URL.RawQuery = query.Encode()
		e.upU.Execute(ctx.Writer, *ctx.Request)
	}
}

func (e *ginEngine) deleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("userId", ctx.Param("userId"))
		ctx.Request.URL.RawQuery = query.Encode()
		e.delU.Execute(ctx.Writer, *ctx.Request)
	}
}

func (e *ginEngine) getByUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("userId", ctx.Param("userId"))
		ctx.Request.URL.RawQuery = query.Encode()
		e.findU.Execute(ctx.Writer, *ctx.Request)
	}
}

func (e *ginEngine) getAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		query := ctx.Request.URL.Query()
		query.Add("Page", ctx.Param("Page"))
		query.Add("Limit", ctx.Param("Limit"))
		ctx.Request.URL.RawQuery = query.Encode()
		e.allU.Execute(ctx.Writer, *ctx.Request)
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
