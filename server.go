package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Retrieve environment variables that may be set within the container.
	ip := os.Getenv("IP")           // IP address of the container
	pod := os.Getenv("POD")         // Name of the pod where the container is running
	node := os.Getenv("NODE")       // Name of the node where the pod is scheduled
	namespace := os.Getenv("NAMESPACE") // Namespace where the pod is deployed

	// Create a new Echo instance to handle HTTP requests.
	e := echo.New()
	
	// Use CORS middleware to handle Cross-Origin Resource Sharing (CORS) headers.
	e.Use(middleware.CORS())
	
	// Define a route for handling incoming HTTP GET requests.
	e.GET("/", func(c echo.Context) error {
		// Create a response map containing some information about the container and its environment.
		res := map[string]interface{}{
			"message":   "Hello from backend version v1.0.0",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"ip":        ip,
			"pod":       pod,
			"node":      node,
			"namespace": namespace,
		}

		// Return the response map as a JSON response with a 200 OK status.
		return c.JSON(http.StatusOK, res)
	})
	
	// Start the Echo web server on port 8080 and handle incoming HTTP requests.
	e.Logger.Fatal(e.Start(":8080"))
}
