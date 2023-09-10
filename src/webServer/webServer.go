package webServer

import (
	"net/http"

	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/api/all", func(c *gin.Context) {
		all, err := boltdb.ExportJSON("./portainer.db", false)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.PureJSON(http.StatusOK, all)
	})

	router.Run()
}
