package main

import (
	"net/http"

	"testcases-gen/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hi this is route '/' of my testcase gen")
	})
	e.POST("/gen-testcases", handler.GenTestcasesHandler)

	e.Start(":8080")
}

// just dummy
func guessNumber(int, int) int { return 0 }
