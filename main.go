package main

import (
	"os"
	"path"
	"strconv"

	"github.com/louisevanderlith/comment/core"
	"github.com/louisevanderlith/comment/routers"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/do"
	"github.com/louisevanderlith/droxolite/element"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()


	r.GET("/message/:key", comment.View)

	messages := r.Group("/message")
	messages.POST("", message.Create)
	messages.PUT("/:key", message.Update)
	messages.DELETE("/:key", message.Delete)

	r.GET("/messages", message.Get)
	r.GET("/messages/:pagesize/*hash", message.Search)


	r.GET("/type/:key", message.View)

	types := r.Group("/type")
	types.POST("", type.Create)
	types.PUT("/:key", type.Update)
	types.DELETE("/:key", type.Delete)

	r.GET("/types", type.Get)
	r.GET("/types/:pagesize/*hash", type.Search)
	

	err := r.Run(":8084")

	if err != nil {
		panic(err)
	}
}

// func main() {
// 	keyPath := os.Getenv("KEYPATH")
// 	pubName := os.Getenv("PUBLICKEY")
// 	host := os.Getenv("HOST")
// 	httpport, _ := strconv.Atoi(os.Getenv("HTTPPORT"))
// 	appName := os.Getenv("APPNAME")
// 	pubPath := path.Join(keyPath, pubName)

// 	// Register with router
// 	srv := bodies.NewService(appName, "", pubPath, host, httpport, servicetype.API)

// 	routr, err := do.GetServiceURL("", "Router.API", false)

// 	if err != nil {
// 		panic(err)
// 	}

// 	err = srv.Register(routr)

// 	if err != nil {
// 		panic(err)
// 	}

// 	poxy := resins.NewMonoEpoxy(srv, element.GetNoTheme(host, srv.ID, "none"))
// 	routers.Setup(poxy)
// 	poxy.EnableCORS(host)

// 	core.CreateContext()
// 	defer core.Shutdown()

// 	err = droxolite.Boot(poxy)

// 	if err != nil {
// 		panic(err)
// 	}
// }
