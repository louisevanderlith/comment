package handles

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"log"
	"net/http"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/core/commenttype"
)

func ViewType(w http.ResponseWriter, r *http.Request) {
	commentType := commenttype.GetEnum(drx.FindParam(r, "type"))
	itemKey, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	result, err := core.GetNodeMessages(itemKey, commentType)

	if err != nil {
		log.Println("Get Node Message Error", err)
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
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	token := r.Context().Value("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	entry.SubjectID = claims["sub"].(string)
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
	key, err := keys.ParseKey(drx.FindParam(r, "key"))

	if err != nil {
		log.Println("Parse Error", err)
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

	err = core.UpdateMessage(key, body)

	if err != nil {
		log.Println("Updated Message Error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON("Saved"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
