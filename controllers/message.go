package controllers

import (
	"encoding/json"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/husk"

	"github.com/louisevanderlith/mango/control"
)

type MessageController struct {
	control.APIController
}

func NewMessageCtrl(ctrlMap *control.ControllerMap) *MessageController {
	result := &MessageController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetMessages
// @Description Gets all comments related to a node.
// @Param	typeID			path 	string 	true		"comment's type"
// @Param	nodeID			path	string 	true		"node's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router /:type/:nodeID[get]
func (req *MessageController) Get() {
	commentType := core.GetCommentType(req.Ctx.Input.Param(":type"))
	nodeKey, err := husk.ParseKey(req.Ctx.Input.Param(":nodeID"))

	if err != nil {
		req.Serve(nil, err)
		return
	}

	result, err := core.GetMessage(nodeKey, commentType)

	req.Serve(result, err)
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func (req *MessageController) Post() {
	var entry core.Message
	err := json.Unmarshal(req.Ctx.Input.RequestBody, &entry)

	if err != nil {
		req.Serve(nil, err)
		return
	}

	rec := core.SubmitMessage(entry)

	req.Serve(rec, nil)
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *MessageController) Put() {
	entry, err := req.GetKeyedRequest()

	if err != nil {
		req.Serve(nil, err)
		return
	}

	err = core.UpdateMessage(entry.Key, entry.Body.(core.Message))

	req.Serve(nil, err)
}
