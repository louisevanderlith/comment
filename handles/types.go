package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/kong/tokens"
	"log"
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
	"github.com/louisevanderlith/husk"
)

func ViewType(w http.ResponseWriter, r *http.Request) {
	commentType := commenttype.GetEnum(drx.FindParam(r, "type"))
	nodeKey, err := husk.ParseKey(drx.FindParam(r, "nodeID"))

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

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title Create Comment
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateType(w http.ResponseWriter, r *http.Request) {
	var entry core.Message
	err := drx.JSONBody(r, &entry)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	tknInfo := drx.GetIdentity(r)

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

	err = mix.Write(w, mix.JSON("Comment Created"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func UpdateType(w http.ResponseWriter, r *http.Request) {
	key, err := husk.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := core.Message{}
	err = drx.JSONBody(r, &body)

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

	err = mix.Write(w, mix.JSON("Saved"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
