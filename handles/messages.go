package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/kong/tokens"
	"log"
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)

	results, err := core.GetAllMessages(1, 10)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON(results))
}

func SearchMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()
	results, err := core.GetAllMessages(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON(results))
}

// @Title GetMessages
// @Description Gets all comments related to a node.
// @Param	typeID			path 	string 	true		"comment's type"
// @Param	nodeID			path	string 	true		"node's ID"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router //:key [get]
func ViewMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	msgKey, err := husk.ParseKey(ctx.FindParam("key"))

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

	ctx.Serve(http.StatusOK, mix.JSON(result))
}

// @Title CreateMessage
// @Description Creates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [post]
func CreateMessage(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	ctx.Serve(http.StatusOK, mix.JSON("Comment Created"))
}

// @Title UpdateMessage
// @Description Updates a comment
// @Param	body		body 	logic.MessageEntry	true		"comment entry"
// @Success 200 {map[string]string} map[string]string
// @Failure 403 body is empty
// @router / [put]
func UpdateMessage(w http.ResponseWriter, r *http.Request) {
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
