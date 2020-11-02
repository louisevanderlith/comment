package main

import (
	"flag"
	"github.com/louisevanderlith/comment/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/comment/core"
)

func main() {
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	audience := flag.String("audience", "comment", "Token target 'aud'")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8084",
		Handler:      handles.SetupRoutes(*audience, *issuer),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
