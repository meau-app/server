package internal

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/meau-app/server/internal/config"
	"github.com/meau-app/server/internal/handlers"
	"github.com/meau-app/server/internal/middleware"
)

func Serve() error {
	e := echo.New()
	e.Server.WriteTimeout = 1 * time.Minute
	e.Server.ReadTimeout = 1 * time.Minute

	config.InitEvironment()
	config.InitDatabase(e.Logger)

	e.Use(middleware.FirebaseAuthentication)

	// HTTP handlers
	e.GET("/pets", handlers.GetPets)
	e.GET("/pets/:id", handlers.GetPet)

	e.POST("/pets", handlers.InsertPet)

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)

	e.POST("/users", handlers.InsertUser)

	//  only for testing purposes, does not expose any information about the
	//  server itself
	e.GET("/health", handlers.Health)

	address := config.Hostname + ":" + config.Port

	return e.Start(address)
}
