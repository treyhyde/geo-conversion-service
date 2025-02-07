package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/geo-conversion-service/internal"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Initialize the server
	server := internal.NewServer()

	// Define routes
	router.POST("/upload", server.UploadHandler)
	router.GET("/download/:filename", server.DownloadHandler)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
