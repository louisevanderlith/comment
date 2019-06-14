package routers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/comment/controllers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service, host string) {
	ctrlmap := EnableFilters(s, host)

	msgCtrl := controllers.NewMessageCtrl(ctrlmap)
	beego.Router("/v1/message", msgCtrl, "put:Put;post:Post")
	beego.Router("/v1/message/all/:pagesize", msgCtrl, "get:GetAll")
	beego.Router("/v1/message/:type/:nodeID", msgCtrl, "get:Get")
}

func EnableFilters(s *mango.Service, host string) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.User
	emptyMap["PUT"] = roletype.User

	ctrlmap.Add("/v1/message", emptyMap)

	adminMap := make(secure.ActionMap)
	adminMap["GET"] = roletype.Admin

	ctrlmap.Add("/v1/message/all", adminMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)
	allowed := fmt.Sprintf("https://*%s", strings.TrimSuffix(host, "/"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{allowed},
		AllowMethods: []string{"GET", "PUT", "POST", "OPTIONS"},
	}), false)

	return ctrlmap
}
