package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rafaelcn/meau/internal/dao"
)

func GetUser(c echo.Context) error {
	
	return nil
}

func GetUsers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	users, err := dao.GetUsers(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func InsertUser(c echo.Context) error {
	return nil
}
