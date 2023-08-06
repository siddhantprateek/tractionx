package routes

import (
	"tractionx/config"

	"github.com/labstack/echo/v4"
)

func Init() error {

	app := echo.New()

	port := config.GetEnvVar("PORT")
	err := app.Start(":" + port)
	return err
}
