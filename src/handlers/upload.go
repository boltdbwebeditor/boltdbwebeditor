package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const dst = "./uploads/"

var filename string = ""

func UploadDBFile(c *gin.Context) {
	// Single file
	file, _ := c.FormFile("file")

	filename = file.Filename

	// Upload the file to specific dst.
	err := c.SaveUploadedFile(file, filepath.Join(dst, file.Filename))
	if err != nil {
		c.Error(err)
		return
	}

	c.PureJSON(200, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}
