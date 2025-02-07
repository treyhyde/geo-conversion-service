package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yourusername/geo-conversion-service/internal"
)

func TestUploadHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	server := internal.NewServer()
	router := gin.Default()
	router.POST("/upload", server.UploadHandler)

	// Create a temporary file to simulate file upload
	tempFile, err := os.CreateTemp("", "testfile.*.shp")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Create a new file upload request
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(tempFile.Name()))
	assert.NoError(t, err)
	_, err = io.Copy(part, tempFile)
	assert.NoError(t, err)
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, "/upload", body)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "File uploaded and converted successfully")
}

func TestDownloadHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	server := internal.NewServer()
	router := gin.Default()
	router.GET("/download/:filename", server.DownloadHandler)

	req, err := http.NewRequest(http.MethodGet, "/download/testfile.mbtiles", nil)
	assert.NoError(t, err)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "url")
}
