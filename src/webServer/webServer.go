package webServer

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/src/cli/flags"
	"net/http"

	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/boltdbwebeditor/boltdbwebeditor/src/helpers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Start(flags *flags.Flags) {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("/app/static", false)))

	router.GET("/api/db", func(c *gin.Context) {
		tempDbPath, err := helpers.CopyDbToTemp(*(flags.DB))
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

		err = helpers.MoveFile(tempDbPath, *(flags.DB))
		if err != nil {
			c.Error(err)
		}

		c.PureJSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
