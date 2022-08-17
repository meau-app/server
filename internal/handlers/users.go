package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/meau-app/server/internal/cache"
	"github.com/meau-app/server/internal/dao"
)

var (
	UserCache *cache.Cache[dao.User]
)

func GetUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	user, err := dao.GetUser(ctx, c.Param("id"))
	if err != nil {
		c.Logger().Errorf(
			"failed to fetch user (%s), reason %v",
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

	users := UserCache.GetAll()

	// cache miss behaviour
	if len(users) == 0 {
		var err error

		users, err = dao.GetUsers(ctx)
		if err != nil {
			c.Logger().Error("failed to fetch users, reason %v", err)
			return err
		}

		err = UserCache.Save(users...)
		if err != nil {
			c.Logger().Warn(
				"failed to save user items to cache, reason %v",
				err,
			)
		}
	}

	if len(users) == 0 {
		return c.String(http.StatusNotFound, "")
	}

	return c.JSON(http.StatusOK, users)
}

func InsertUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	user := &dao.User{}

	if err := c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "failed to parse request")
	}

	dao.SaveUser(ctx, user)

	return c.String(http.StatusCreated, "")
}

func DeleteUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	return dao.DeleteUser(ctx, c.Param("id"))
}
