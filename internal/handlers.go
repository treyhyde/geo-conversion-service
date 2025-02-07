package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type Server struct {
	s3Client *s3.S3
	bucket   string
}

func NewServer() *Server {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))
	s3Client := s3.New(sess)

	return &Server{
		s3Client: s3Client,
		bucket:   "your-s3-bucket-name",
	}
}

func (s *Server) UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to upload file: %v", err)
		return
	}

	// Validate the uploaded file
	if !isValidFile(file.Filename) {
		c.String(http.StatusBadRequest, "Invalid file format")
		return
	}

	// Save the uploaded file to a temporary location
	tempFilePath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveUploadedFile(file, tempFilePath); err != nil {
		c.String(http.StatusInternalServerError, "Failed to save uploaded file: %v", err)
		return
	}

	// Convert the file to MBTile pack using ogr2ogr
	convertedFilePath, err := convertToMBTile(tempFilePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to convert file: %v", err)
		return
	}

	// Upload the converted file to S3
	s3Key := fmt.Sprintf("converted/%s", filepath.Base(convertedFilePath))
	if err := s.uploadToS3(convertedFilePath, s3Key); err != nil {
		c.String(http.StatusInternalServerError, "Failed to upload file to S3: %v", err)
		return
	}

	// Generate a pre-signed URL for the uploaded file
	presignedURL, err := s.generatePresignedURL(s3Key)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate pre-signed URL: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded and converted successfully",
		"url":     presignedURL,
	})
}

func (s *Server) DownloadHandler(c *gin.Context) {
	filename := c.Param("filename")
	s3Key := fmt.Sprintf("converted/%s", filename)

	// Generate a pre-signed URL for the file
	presignedURL, err := s.generatePresignedURL(s3Key)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate pre-signed URL: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": presignedURL,
	})
}

func isValidFile(filename string) bool {
	// Add validation logic to ensure the file is legal and understood by GDAL
	// For simplicity, we'll just check the file extension
	ext := filepath.Ext(filename)
	return ext == ".shp" || ext == ".geojson" || ext == ".kml"
}

func convertToMBTile(inputFilePath string) (string, error) {
	// Implement the logic to use ogr2ogr to convert the file to MBTile pack
	// For simplicity, we'll just return the input file path as the converted file path
	return inputFilePath, nil
}

func (s *Server) uploadToS3(filePath, s3Key string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s3Key),
		Body:   file,
	})
	return err
}

func (s *Server) generatePresignedURL(s3Key string) (string, error) {
	req, _ := s.s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s3Key),
	})
	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}
	return url, nil
}
