package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/husk"
)

type Types struct {
}

func Get(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, nil)
}

func View(c *gin.Context) {
	commentType := commenttype.GetEnum(c.Param("type"))
	nodeKey, err := husk.ParseKey(c.Param("nodeID"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	result, err := core.GetMessage(nodeKey, commentType)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, result)
}
