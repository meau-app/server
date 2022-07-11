package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// this file contains handlers that are exposed to the web, these handlers are
// uses HTTP.

func getAnimal(c echo.Context) error {
	return nil
}

func getAnimals(c echo.Context) error {
	return nil
}

func insertAnimal(c echo.Context) error {
	return nil
}

func health(c echo.Context) error {
	type Response struct {
		Message string `json:"message"`
	}

	response := Response{
		Message: "ok",
	}

	return c.JSON(http.StatusOK, response)
}
