package controllers

import (
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Messages struct {
}

// @router /all/:pagesize [get]
func (req *Messages) GetAll(ctx context.Contexer) (int, interface{}) {
	page, size := ctx.GetPageData()
	results := core.GetAllMessages(page, size)

	return http.StatusOK, results
}

// @Title GetMessages
// @Description Gets all comments related to a node.
// @Param	typeID			path 	string 	true		"comment's type"
// @Param	nodeID			path	string 	true		"node's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /:type/:nodeID [get]
func (req *Messages) Get(ctx context.Contexer) (int, interface{}) {
	commentType := commenttype.GetEnum(ctx.FindParam("type"))
	nodeKey, err := husk.ParseKey(ctx.FindParam("nodeID"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result, err := core.GetMessage(nodeKey, commentType)

	if err != nil {
		return http.StatusNotFound, err
	}

	return http.StatusOK, result
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *Messages) Post(ctx context.Contexer) (int, interface{}) {
	var entry core.Message
	err := ctx.Body(&entry)

	if err != nil {
		return http.StatusBadRequest, err
	}

	rec := core.SubmitMessage(entry)

	if rec.Error != nil {
		return http.StatusInternalServerError, rec.Error
	}

	return http.StatusOK, rec.Record
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *Messages) Put(ctx context.Contexer) (int, interface{}) {
	body := core.Message{}
	key, err := ctx.GetKeyedRequest(&body)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = core.UpdateMessage(key, body)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, "Saved"
}
