package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/kong/tokens"
	"log"
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

func ViewType(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	commentType := commenttype.GetEnum(ctx.FindParam("type"))
	nodeKey, err := husk.ParseKey(ctx.FindParam("nodeID"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.GetNodeMessage(nodeKey, commentType)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON(result))
}

// @Title Create Comment
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateType(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	var entry core.Message
	err := ctx.Body(&entry)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	tknInfo := ctx.GetTokenInfo()

	if !tknInfo.HasUser() {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	k, err := husk.ParseKey(tknInfo.GetClaimString(tokens.UserKey))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	entry.UserKey = k
	err = entry.SubmitMessage()

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON("Comment Created"))
}

func UpdateType(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := core.Message{}
	err = ctx.Body(&body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.UpdateMessage(key, body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON("Saved"))
}
