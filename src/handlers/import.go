package handlers

import (
	"net/http"

	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/gin-gonic/gin"
)

func Import(c *gin.Context) {
	all, err := boltdb.ImportJSON("./portainer.db", false)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.PureJSON(http.StatusOK, all)
}
