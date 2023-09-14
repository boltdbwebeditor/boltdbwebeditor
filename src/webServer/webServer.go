package webServer

import (
	"github.com/boltdbwebeditor/boltdbwebeditor/src/handlers"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/api/db", handlers.Import)

	router.POST("/api/db", handlers.Export)

	router.POST("/api/db/upload", handlers.UploadDBFile)

	router.GET("/api/db/download", handlers.DownloadDBFile)

	// listen on :8080
	router.Run()
}
