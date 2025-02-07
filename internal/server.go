package internal

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{
		router: gin.Default(),
	}
	server.routes()
	return server
}

func (s *Server) routes() {
	s.router.POST("/upload", s.UploadHandler)
	s.router.GET("/download/:filename", s.DownloadHandler)
}

func (s *Server) Start() {
	if err := s.router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *Server) UploadHandler(c *gin.Context) {
	// Placeholder for file upload handler
	c.String(http.StatusOK, "Upload handler")
}

func (s *Server) DownloadHandler(c *gin.Context) {
	// Placeholder for file download handler
	c.String(http.StatusOK, "Download handler")
}
