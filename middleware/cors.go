package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func CORSMiddleware() echo.MiddlewareFunc {
	config := middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderContentLength, echo.HeaderAccept},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	}
	return middleware.CORSWithConfig(config)
}