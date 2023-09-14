package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadDBFile(c *gin.Context) {
	// fileName := "portainer.db"

	// Open the file
	file, err := os.Open(filepath.Join(dst, filename))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the file contents
	fileContents, err := io.ReadAll(file)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Set the appropriate headers for file download
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", fileContents)
}
