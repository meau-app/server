package internal

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/rafaelcn/meau/internal/config"
	"github.com/rafaelcn/meau/internal/handlers"
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
	e.GET("/pets", handlers.GetPet)
	e.GET("/pets/:id", handlers.GetPets)

	e.POST("/pets", handlers.InsertPet)

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)

	e.POST("/users", handlers.InsertUser)

	//
	e.GET("/health", handlers.Health)

	address := config.Hostname + ":" + config.Port

	return e.Start(address)
}
