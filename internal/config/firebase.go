package config

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

var (
	Database *firebase.App
)

func InitDatabase(logger echo.Logger) {
	var err error

	file := GetEnvOrDefault("MEAU_CREDENTIALS", "firebase.json")
	credentials := option.WithCredentialsFile(
		file,
	)

	Database, err = firebase.NewApp(context.Background(), nil, credentials)

	if err != nil {
		logger.Fatalf("failed to initialize database", err)
	}
}
