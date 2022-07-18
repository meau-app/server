package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rafaelcn/meau/internal/dao"
)

func GetPet(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	user, err := dao.GetPet(ctx, c.Param("id"))
	if err != nil {
		c.Logger().Errorf(
			"failed to fetch pet (%s), reason %v",
			c.Param("id"),
			err,
		)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func GetPets(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	pets, err := dao.GetPets(ctx)
	if err != nil {
		c.Logger().Error("failed to fetch pets, reason %v", err)
		return err
	}

	if len(pets) == 0 {
		return c.String(http.StatusNotFound, "")
	}

	return c.JSON(http.StatusOK, pets)
}

func InsertPet(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	ctx = context.WithValue(ctx, dao.ContextLoggerKey, c.Logger())

	pet := &dao.Pet{}

	if err := c.Bind(pet); err != nil {
		c.String(http.StatusBadRequest, "")
	}

	dao.SavePet(ctx, pet)

	return c.String(http.StatusCreated, "")
}
