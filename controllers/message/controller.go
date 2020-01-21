package message

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/droxo"
	"github.com/louisevanderlith/husk"
	"net/http"
)

func Get(c *gin.Context) {
	results := core.GetAllMessages(1, 10)

	c.JSON(http.StatusOK, results)
}

// @router /all/:pagesize [get]
func Search(c *gin.Context) {
	page, size := droxo.GetPageData(c.Param("pagesize"))
	results := core.GetAllMessages(page, size)

	c.JSON(http.StatusOK, results)
}

// @Title GetMessages
// @Description Gets all comments related to a node.
// @Param	typeID			path 	string 	true		"comment's type"
// @Param	nodeID			path	string 	true		"node's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router //:nodeKey?type= [get]
func View(c *gin.Context) {
	//commentType := commenttype.GetEnum(ctx.FindParam("type"))
	msgKey, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	result, err := core.GetMessageByKey(msgKey)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	c.JSON(http.StatusOK, result)
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func Create(c *gin.Context) {
	var entry core.Message
	err := c.Bind(&entry)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	rec := core.SubmitMessage(entry)

	if rec.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, rec.Error)
		return
	}

	c.JSON(http.StatusOK, rec.Record)
}

// @Title UpdateMessage
// @Description Updates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func Update(c *gin.Context) {
	key, err := husk.ParseKey(c.Param("key"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body := core.Message{}
	err = c.Bind(&body)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = core.UpdateMessage(key, body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "Saved")
}
