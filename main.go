package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louisevanderlith/comment/controllers/message"
	"github.com/louisevanderlith/comment/core"
)

func main() {
	core.CreateContext()
	defer core.Shutdown()

	r := gin.Default()

	r.GET("/message/:key", message.View)

	messages := r.Group("/message")
	messages.POST("", message.Create)
	messages.PUT("/:key", message.Update)

	r.GET("/messages", message.Get)
	r.GET("/messages/:pagesize/*hash", message.Search)

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
