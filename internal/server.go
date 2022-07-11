package internal

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/rafaelcn/meau/internal/config"
	"github.com/rafaelcn/meau/internal/middleware"
)

func Serve() error {
	e := echo.New()
	e.Server.WriteTimeout = 1 * time.Minute
	e.Server.ReadTimeout = 1 * time.Minute

	config.InitEvironment()
	config.InitDatabase(e.Logger)

	e.Use(middleware.FirebaseAuthentication)

	// HTTP handlers
	e.GET("/animals", getAnimals)
	e.GET("/animals/:id", getAnimal)
	e.POST("/animals", insertAnimal)

	e.GET("/health", health)

	address := config.Hostname + ":" + config.Port

	return e.Start(address)
}
