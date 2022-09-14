package internal

import (
	"time"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"

	"github.com/meau-app/server/internal/cache"
	"github.com/meau-app/server/internal/config"
	"github.com/meau-app/server/internal/dao"
	"github.com/meau-app/server/internal/handlers"
	"github.com/meau-app/server/internal/middleware"
)

func Serve() error {
	e := echo.New()
	e.Server.WriteTimeout = 1 * time.Minute
	e.Server.ReadTimeout = 1 * time.Minute

	config.InitEvironment()
	config.InitDatabase(e.Logger)

	e.Use(m.Logger())
	e.Use(middleware.FirebaseAuthentication)

	// HTTP handlers
	e.GET("/pets", handlers.GetPets)
	e.GET("/pets/:id", handlers.GetPet)
	e.POST("/pets", handlers.InsertPet)
	e.DELETE("/pets/:id", handlers.DeletePet)

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.POST("/users", handlers.InsertUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	//  only for testing purposes, does not expose any information about the
	//  server itself
	e.GET("/health", handlers.Health)

	// initializing caches
	handlers.UserCache = cache.NewCache[dao.User](
		cache.TypeUser,
		cache.CacheConfig{
			TTL:    1 * time.Minute,
			Logger: e.Logger,
		},
	)
	handlers.PetCache = cache.NewCache[dao.Pet](
		cache.TypePet,
		cache.CacheConfig{
			TTL: 40 * time.Second,
		},
	)

	address := config.Hostname + ":" + config.Port

	return e.Start(address)
}
