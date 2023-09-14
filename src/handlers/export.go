package handlers

import (
	"net/http"

	"github.com/boltdbwebeditor/boltdbwebeditor/src/boltdb"
	"github.com/gin-gonic/gin"
)

func Export(c *gin.Context) {
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
}
