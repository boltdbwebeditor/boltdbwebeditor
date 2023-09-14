package webServer

import (
	"net/http"

	"github.com/boltdbwebeditor/boltdbwebeditor/src/handlers"
	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("/app/static", false)))

  router.GET("/api/db", handlers.Import)

	router.POST("/api/db", handlers.Export)

	router.POST("/api/db/upload", handlers.UploadDBFile)

	router.GET("/api/db/download", handlers.DownloadDBFile)

	router.Run(":8080")
}
