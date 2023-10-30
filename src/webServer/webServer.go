package webServer

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/src/cli/flags"
	"github.com/boltdbwebeditor/boltdbwebeditor/src/helpers"
	"github.com/boltdbwebeditor/boltdbwebeditor/src/libs/boltdb"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start(flags *flags.Flags) {
	dbPath := *(flags.DbPath)

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("/app/static", false)))

	router.GET("/api/db", func(c *gin.Context) {
		data, err := boltdb.ForceRead(dbPath, true)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err, "errorMsg": err.Error()},
			)
			return
		}

		c.PureJSON(http.StatusOK, data)
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
			return
		}

		c.PureJSON(http.StatusOK, data)
	})

	router.POST("/api/db/file", func(c *gin.Context) {
		file, err := c.FormFile("upload.bolt.db")
		if err != nil {
			c.Error(err)
			return
		}

		err = c.SaveUploadedFile(file, dbPath)
		if err != nil {
			c.Error(err)
			return
		}

		data, err := boltdb.ForceRead(dbPath, true)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err, "errorMsg": err.Error()},
			)
			return
		}

		c.PureJSON(http.StatusOK, gin.H{"data": data})
	})

	router.Run(":8080")
}
