package handles

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"log"
	"net/http"

	"github.com/louisevanderlith/comment/core"
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
	msgKey, err := keys.ParseKey(drx.FindParam(r, "key"))

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
	entry := core.Message{ItemKey: keys.CrazyKey()}

	err := drx.JSONBody(r, &entry)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	token := r.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	entry.SubjectID = claims["sub"].(string)
	err = entry.SubmitMessage()

	if err != nil {
		log.Println("Submit Error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := core.Message{ItemKey: keys.CrazyKey()}
	err = drx.JSONBody(r, &body)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	token := r.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	subj := claims["sub"].(string)

	if body.SubjectID != subj {
		log.Println("Wrong user", subj, "for", body.SubjectID)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	err = core.UpdateMessage(key, body)

	if err != nil {
		log.Println("Update Message Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON("Saved"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
