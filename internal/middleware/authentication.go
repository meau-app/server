package middleware

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/meau-app/server/internal/config"
)

func FirebaseAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()
		if ip == "::1" || ip == "127.0.0.1" || config.Authenticate == "false" {
			return next(c)
		}

		ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
		defer cancel()

		client, err := config.Database.Auth(ctx)

		if err != nil {
			c.Logger().Errorf(
				"failed to acquire authentication client, reason %v",
				err,
			)
			return err
		}

		uid := c.FormValue("token")

		_, err = client.VerifyIDTokenAndCheckRevoked(ctx, uid)

		if err != nil {
			c.Logger().Errorf("token not valid, reason %v", err)
			return err
		}

		return next(c)
	}
}
