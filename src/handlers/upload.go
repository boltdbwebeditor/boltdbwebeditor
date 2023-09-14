package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const dst = "./uploads/"

func UploadDBFile(c *gin.Context) {
	// Single file
	file, _ := c.FormFile("file")

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, filepath.Join(dst, file.Filename))

	c.PureJSON(200, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}
