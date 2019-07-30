package controllers

import (
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

type MessageController struct {
	xontrols.APICtrl
}

// @router /all/:pagesize [get]
func (req *MessageController) GetAll() {
	page, size := req.GetPageData()
	results := core.GetAllMessages(page, size)

	req.Serve(http.StatusOK, nil, results)
}

// @Title GetMessages
// @Description Gets all comments related to a node.
// @Param	typeID			path 	string 	true		"comment's type"
// @Param	nodeID			path	string 	true		"node's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /:type/:nodeID [get]
func (req *MessageController) Get() {
	commentType := commenttype.GetEnum(req.FindParam("type"))
	nodeKey, err := husk.ParseKey(req.FindParam("nodeID"))

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	result, err := core.GetMessage(nodeKey, commentType)

	if err != nil {
		req.Serve(http.StatusNotFound, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, result)
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *MessageController) Post() {
	var entry core.Message
	err := req.Body(&entry)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	rec := core.SubmitMessage(entry)

	if rec.Error != nil {
		req.Serve(http.StatusInternalServerError, rec.Error, nil)
		return
	}

	req.Serve(http.StatusOK, nil, rec.Record)
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *MessageController) Put() {
	body := core.Message{}
	key, err := req.GetKeyedRequest(&body)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = core.UpdateMessage(key, body)

	if err != nil {
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, "Saved")
}
