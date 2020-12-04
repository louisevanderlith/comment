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

	r.HandleFunc("/{type:[a-z]+}/{key:[0-9]+\\x60[0-9]+}", ViewMessage).Methods(http.MethodGet)

	r.HandleFunc("/messages", GetMessages).Methods(http.MethodGet)
	r.HandleFunc("/messages/{key:[0-9]+\\x60[0-9]+}", ViewMessage).Methods(http.MethodGet)

	r.HandleFunc("/messages/{pagesize:[A-Z][0-9]+}", SearchMessage).Methods(http.MethodGet)
	r.HandleFunc("/messages/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchMessage).Methods(http.MethodGet)

	r.HandleFunc("/messages", CreateMessage).Methods(http.MethodPost)

	r.HandleFunc("/messages/{key:[0-9]+\\x60[0-9]+}", UpdateMessage).Methods(http.MethodPut)

	r.Use(mw.Handler)

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
