package routes

import (
	"tractionx/config"
	handlers "tractionx/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

var (
	setup handlers.OrgSetup
)

func ApiRoutes(router *echo.Echo) {
	router.GET("/", handlers.BaseRoute)

	router.GET("/query", setup.Query)
	router.POST("/invoke", setup.Invoke)
}

func Init() error {

	app := echo.New()

	// logger middleware
	log := logrus.New()
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(_ echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))

	// routes
	ApiRoutes(app)

	port := config.GetEnvVar("PORT")
	err := app.Start(":" + port)
	return err
}
