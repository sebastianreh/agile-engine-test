package server

import (
	"agile-engine-test/cmd/api/controller"
	"github.com/labstack/echo/v4"
	"fmt"
)

func SetupServer() {
	settings := InitializeSettings()
	server := echo.New()
	url := fmt.Sprintf("%s:%s", settings.Host, settings.Port)
	setupRoutes(server, settings)
	server.Logger.Fatal(server.Start(url))
}

func setupRoutes(server *echo.Echo, settings ProjectSettings) {
	statusController := controller.NewStatusController(settings.ProjectName, settings.ProjectVersion)
	userController := controller.NewUser()

	base := server.Group("/agile-engine-test")
	base.GET("/health", statusController.HealthCheck)

	user := base.Group("/user")
	user.GET("/history", userController.FetchHistory)
	user.POST("/transaction/commit", userController.CommitTransaction)
	user.GET("/history/:transactionID", userController.GetTransaction)
}