package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/kong/tokens"
	"log"
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/husk"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	results, err := core.GetAllMessages(1, 10)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

func SearchMessage(w http.ResponseWriter, r *http.Request) {
	page, size := drx.GetPageData(r)
	results, err := core.GetAllMessages(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON(results))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title GetMessages
// @Description Gets all comments related to a node.
// @Param	typeID			path 	string 	true		"comment's type"
// @Param	nodeID			path	string 	true		"node's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router //:key [get]
func ViewMessage(w http.ResponseWriter, r *http.Request) {
	msgKey, err := husk.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.GetMessageByKey(msgKey)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = mix.Write(w, mix.JSON(result))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateMessage(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = mix.Write(w, mix.JSON("Comment Created"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

// @Title UpdateMessage
// @Description Updates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
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
