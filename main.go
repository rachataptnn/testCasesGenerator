package main

import (
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e.Start(":" + port)
}

// just dummy
func guessNumber(int, int) int { return 0 }
