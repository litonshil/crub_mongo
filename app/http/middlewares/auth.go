package middlewares

import (
	"crud_mongo/infra/config"

	"github.com/labstack/echo/v4"
)

func Auth(middleware *JWTMiddleware) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(JWTConfig{
		SigningKey: []byte(config.Jwt().AccessTokenSecret),
		ContextKey: config.Jwt().ContextKey,
	})
}
