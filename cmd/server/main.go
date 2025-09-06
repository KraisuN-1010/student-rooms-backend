package main

import (
	"log"

	"gofr.dev/pkg/gofr"
	"github.com/KraisuN-1010/student-rooms-backend/internal/handlers"
)

func main() {
	// Create new GoFr app
	app := gofr.New()

	// Health check route
	app.GET("/", func(c *gofr.Context) (interface{}, error) {
		return "Student Rooms Backend is running!", nil
	})

	// Sample rooms route
	app.GET("/rooms", handlers.GetRooms)

	// Set server port explicitly to 8000
	const port = "8000"

	log.Println("Starting server on http://localhost:" + port)
	app.Run() // GoFr uses its internal default port (8000 in your case)
}
