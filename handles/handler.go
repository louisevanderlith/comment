package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl string) http.Handler {
	/*
		e.JoinBundle("/", roletype.User, mix.JSON, &handles.Messages{})

		tps := &handles.Types{}
		e.JoinPath(e.Router().(*mux.Router), "/{type:[a-zA-Z]+}/{nodeID:[0-9]+\x60[0-9]+}", "View Article for Type", http.MethodGet, roletype.Unknown, mix.JSON, tps.View)
		e.JoinPath(e.Router().(*mux.Router), "", "Create Message", http.MethodPost, roletype.User, mix.JSON, tps.Create)
	*/

	r := mux.NewRouter()

	/*"comment.messages.view","comment.messages.create","comment.messages.update","comment.messages.delete"*/

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "comment.messages.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst, //you service is available and allowed for this base url
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
