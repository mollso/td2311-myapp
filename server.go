package main

import (
	"crypto/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

var memConsumed []byte

func main() {
	// Retrieve environment variables that may be set within the container.
	ip := os.Getenv("IP")               // IP address of the container
	pod := os.Getenv("POD")             // Name of the pod where the container is running
	node := os.Getenv("NODE")           // Name of the node where the pod is scheduled
	namespace := os.Getenv("NAMESPACE") // Namespace where the pod is deployed

	// Create a new Echo instance to handle HTTP requests.
	e := echo.New()

	// Define a route for handling incoming HTTP GET requests on the root ("/") path.
	e.GET("/", func(c echo.Context) error {
		// Create a response message containing the information about the container and its environment.
		response := "Timestamp: " + time.Now().UTC().Format(time.RFC3339) + "\n" +
			"IP: " + ip + "\n" +
			"Pod: " + pod + "\n" +
			"Node: " + node + "\n" +
			"Namespace: " + namespace

		// Return the response as plain text with a 200 OK status.
		return c.String(http.StatusOK, response)
	})

	// Define a route for the "stress" endpoint.
	e.GET("/stress", func(c echo.Context) error {
		// Check if memory has already been consumed.
		if memConsumed != nil {
			return c.String(http.StatusOK, "RAM stress test is already running.")
		}

		// Define the target memory size in bytes (1 GB).
		const targetSize = 1024 * 1024 * 1024 // 1 GB

		// Allocate memory for the entire 1 GB.
		memConsumed = make([]byte, targetSize)

		// Fill the allocated memory with random data.
		rand.Read(memConsumed)

		// Return a response indicating that RAM has been consumed.
    return c.String(http.StatusOK, "RAM stress test initiated on Pod: "+pod)
	})


	// Define a route for the "kill" endpoint.
	e.GET("/kill", func(c echo.Context) error {
        // Return a message indicating that the server will be terminated, including the pod name.
        response := "Server in Pod: " + pod + " will be terminated."

        // Start a goroutine to allow the response to be sent before termination.
        go func() {
            time.Sleep(1 * time.Second) // Give some time for the response to be sent
            os.Exit(0) // Terminate the server after a delay
        }()

        return c.String(http.StatusOK, response)
	})

	// Start the Echo web server on port 8080 and handle incoming HTTP requests.
	e.Logger.Fatal(e.Start(":8080"))
}
