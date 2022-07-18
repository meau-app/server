package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/rafaelcn/meau/internal/dao"
)

func GetUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	user, err := dao.GetUser(ctx, c.Param("id"))
	if err != nil {
		c.Logger().Errorf(
			"failed to find user (%s), reason %v",
			c.Param("id"),
			err,
		)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	users, err := dao.GetUsers(ctx)
	if err != nil {
		c.Logger().Error("failed to find users, reason %v", err)
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func InsertUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	user := &dao.User{}

	if err := c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "")
	}

	dao.SaveUser(ctx, user)

	return c.String(http.StatusCreated, "")
}
