package webServer

import (
	"net/http"

	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/api/db", func(c *gin.Context) {
		all, err := boltdb.ImportJSON("./portainer.db", false)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.PureJSON(http.StatusOK, all)
	})

	router.POST("/api/db", func(c *gin.Context) {
		var data map[string]interface{}
		err := c.BindJSON(&data)
		if err != nil {
			c.Error(err)
		}

		err = boltdb.ExportJSON("./portainer.db", data, false)
		if err != nil {
			c.Error(err)
		}

		c.PureJSON(http.StatusOK, data)
	})

	// listen on :8080
	router.Run()
}
