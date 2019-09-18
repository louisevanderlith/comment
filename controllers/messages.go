package controllers

import (
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Messages struct {
}

func (req *Messages) Get(ctx context.Requester) (int, interface{}) {
	results := core.GetAllMessages(1, 10)

	return http.StatusOK, results
}

// @router /all/:pagesize [get]
func (req *Messages) Search(ctx context.Requester) (int, interface{}) {
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
// @router //:nodeKey?type= [get]
func (req *Messages) View(ctx context.Requester) (int, interface{}) {
	//commentType := commenttype.GetEnum(ctx.FindParam("type"))
	msgKey, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	result, err := core.GetMessageByKey(msgKey)

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
func (req *Messages) Create(ctx context.Requester) (int, interface{}) {
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

// @Title UpdateMessage
// @Description Updates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func (req *Messages) Update(ctx context.Requester) (int, interface{}) {
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	body := core.Message{}
	err = ctx.Body(&body)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = core.UpdateMessage(key, body)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, "Saved"
}

func (x *Messages) Delete(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}
