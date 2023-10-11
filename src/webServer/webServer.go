package webServer

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/boltdbwebeditor/boltdbwebeditor/src/helpers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

const dbPath = "/Users/cmeng/Work/devkit/data-ee/portainer.db"

func Start() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("/app/static", false)))

	router.GET("/api/db", func(c *gin.Context) {
		tempDbPath, err := helpers.CopyDbToTemp(dbPath)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err, "errorMsg": err.Error()},
			)
			return
		}

		all, err := boltdb.Read(tempDbPath, true)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err, "errorMsg": err.Error()},
			)
			return
		}

		c.PureJSON(http.StatusOK, all)
	})

	router.POST("/api/db", func(c *gin.Context) {
		var data map[string]interface{}
		err := c.BindJSON(&data)
		if err != nil {
			c.Error(err)
		}

		tempDbPath := helpers.GenerateDbTmpFilePath()

		err = boltdb.Create(tempDbPath, data)
		if err != nil {
			c.Error(err)
		}

		err = helpers.MoveFile(tempDbPath, dbPath)
		if err != nil {
			c.Error(err)
		}

		c.PureJSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
