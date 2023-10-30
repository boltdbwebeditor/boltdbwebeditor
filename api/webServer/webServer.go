package webServer

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/api/cli/flags"
	"github.com/boltdbwebeditor/boltdbwebeditor/api/libs/boltdb"
	"github.com/boltdbwebeditor/boltdbwebeditor/api/libs/tempFile"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Start(flags *flags.Flags) {
	dbPath := *(flags.DbPath)

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("/app/static", false)))

	router.GET("/api/db/json", func(c *gin.Context) {
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

	router.POST("/api/db/json", func(c *gin.Context) {
		var data map[string]interface{}
		err := c.BindJSON(&data)
		if err != nil {
			c.Error(err)
		}

		tempDbPath := tempFile.GenerateDbTmpFilePath()

		err = boltdb.Create(tempDbPath, data)
		if err != nil {
			c.Error(err)
		}

		err = tempFile.MoveFile(tempDbPath, dbPath)
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

	router.GET("/api/db/file", func(c *gin.Context) {
		file, err := os.Open(dbPath)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fileContents, err := io.ReadAll(file)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Header("Content-Disposition", "attachment; filename="+filepath.Base(dbPath))
		c.Header("Content-Type", "application/octet-stream")
		c.Data(http.StatusOK, "application/octet-stream", fileContents)
	})

	router.Run(":8080")
}
