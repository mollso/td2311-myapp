package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ip := os.Getenv("IP")
	pod := os.Getenv("POD")
	node := os.Getenv("NODE")
	namespace := os.Getenv("NAMESPACE")

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		res := map[string]interface{}{
			"message":   "Hello from backend version v1.0.0",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"ip":        ip,
			"pod":       pod,
			"node":      node,
			"namespace": namespace,
		}

		// Return mock data as JSON response
		return c.JSON(http.StatusOK, res)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
