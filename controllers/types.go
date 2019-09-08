package controllers

import (
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

type Types struct {
}

func (x *Types) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

func (x *Types) View(ctx context.Requester) (int, interface{}) {
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
