package middleware

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"strings"
)

type FirebaseTokenVerifiedMiddleware interface {
	GetFirebaseTokenVerifiedMiddleware() echo.MiddlewareFunc
}

type firebaseTokenVerifiedMiddleware struct {
	client *auth.Client
}

func NewFirebaseTokenVerifiedMiddleware(client *auth.Client) FirebaseTokenVerifiedMiddleware {
	return &firebaseTokenVerifiedMiddleware{
		client,
	}
}

func (m *firebaseTokenVerifiedMiddleware) GetFirebaseTokenVerifiedMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqToken := c.Request().Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer")

			if len(splitToken) != 2 {
				return echo.NewHTTPError(400, "Bad Request")
			}

			_, err := m.client.VerifyIDToken(context.Background(), strings.TrimSpace(splitToken[1]))

			if err != nil {
				return echo.NewHTTPError(401, "UnAuthorize")
			}

			if err := next(c); err != nil {
				c.Error(err)
			}

			return nil
		}
	}
}
