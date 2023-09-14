package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/gin-gonic/gin"
)

func Import(c *gin.Context) {
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file uploaded"})
		return
	}

	all, err := boltdb.ImportJSON(filepath.Join(dst, filename), false)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.PureJSON(http.StatusOK, all)
}
