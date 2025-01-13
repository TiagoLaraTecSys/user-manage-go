package setup

import (
	"context"
	"projeto-final/adapter/api/controller"
	"projeto-final/adapter/database"
	"projeto-final/core/usecase"
	"projeto-final/infrastructure/config"
	infraDb "projeto-final/infrastructure/database"
	"projeto-final/infrastructure/http/router"
	"projeto-final/infrastructure/http/server"
	"projeto-final/infrastructure/logger"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type configuration struct {
	configApp *config.AppConfig
	webServer server.Server
	db        *database.UserRepository
	router    router.GinRouter
}

func NewConfig() *configuration {
	return &configuration{}
}

func (c *configuration) InitLogger() *configuration {
	logger.NewZapLogger()

	logger.Info("Log has been successfully configured")
	return c
}

func (c *configuration) WithAppConfig() *configuration {
	var err error
	c.configApp, err = config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	return c
}

func (c *configuration) WithDB() *configuration {
	db, err := infraDb.NewSQLConnection(c.configApp.MySQL.Host)
	if err != nil {
		logger.Fatal(err)
	}

	c.db = database.NewUserRepository(db)
	logger.Info("DB has been successfully configured")
	return c
}

func (c *configuration) WithRouter() *configuration {
	su := controller.NewSaveController(usecase.NewSaveUser(c.db))
	fu := controller.NewFindByUserIdController(usecase.NewFindByUserId(c.db))
	au := controller.NewFindAllUsersController(usecase.NewFindAllUsers(c.db))
	du := controller.NewDeleteUserController(usecase.NewDeleteUser(c.db))
	c.router = router.NewGinEngine(gin.Default(), su, fu, au, du)
	return c
}

func (c *configuration) WithWebServer() *configuration {
	intPort, err := strconv.ParseInt(c.configApp.Application.Server.Port, 10, 64)
	if err != nil {
		logger.Fatal(err)
	}

	intDuration, err := time.ParseDuration(c.configApp.Application.Server.Timeout + "s")
	if err != nil {
		logger.Fatal(err)
	}

	c.webServer = server.NewWebServer(c.router, intPort, intDuration*time.Second)
	logger.Info("Web server has been successfully configurated")
	return c
}

func (c *configuration) Start(ctx context.Context, wg *sync.WaitGroup) {
	logger.Info("App running on port %s", c.configApp.Application.Server.Port)
	c.webServer.Listen(ctx, wg)

}
