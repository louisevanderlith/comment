package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(audience, issuer string) http.Handler {

	mw := open.BearerMiddleware(audience, issuer)

	r := mux.NewRouter()
	r.Handle("/messages", mw.Handler(http.HandlerFunc(GetMessages))).Methods(http.MethodGet)
	r.Handle("/messages/{key:[0-9]+\\x60[0-9]+}", mw.Handler(http.HandlerFunc(ViewMessage))).Methods(http.MethodGet)

	r.Handle("/messages/{pagesize:[A-Z][0-9]+}", mw.Handler(http.HandlerFunc(SearchMessage))).Methods(http.MethodGet)
	r.Handle("/messages/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", mw.Handler(http.HandlerFunc(SearchMessage))).Methods(http.MethodGet)

	r.Handle("/messages", mw.Handler(http.HandlerFunc(CreateMessage))).Methods(http.MethodPost)

	r.Handle("/messages", mw.Handler(http.HandlerFunc(UpdateMessage))).Methods(http.MethodPut)

	/*"comment.messages.view","comment.messages.create","comment.messages.update","comment.messages.delete"*/

	//lst, err := middle.Whitelist(http.DefaultClient, securityUrl, "comment.messages.view", scrt)

	//if err != nil {
	//	panic(err)
	//}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
