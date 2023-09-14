package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/gin-gonic/gin"
)

func Export(c *gin.Context) {
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file uploaded"})
		return
	}

	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		c.Error(err)
	}

	err = boltdb.ExportJSON(filepath.Join(dst, filename), data, false)
	if err != nil {
		c.Error(err)
	}

	c.PureJSON(http.StatusOK, data)
}
