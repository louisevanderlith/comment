package routers

import (
	"github.com/louisevanderlith/comment/controllers"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(poxy resins.Epoxi) {
	//Message
	msgCtrl := &controllers.Messages{}
	msgGroup := routing.NewRouteGroup("message", mix.JSON)
	msgGroup.AddRoute("Create Message", "", "POST", roletype.User, msgCtrl.Post)
	msgGroup.AddRoute("Update Message", "", "PUT", roletype.User, msgCtrl.Put)
	msgGroup.AddRoute("All Messges", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, msgCtrl.GetAll)
	msgGroup.AddRoute("Get Comments for Item", "/{type}/{nodeID:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, msgCtrl.Get)
	poxy.AddGroup(msgGroup)

	/*ctrlmap := EnableFilters(s, host)

	msgCtrl := controllers.NewMessageCtrl(ctrlmap)
	beego.Router("/v1/message", msgCtrl, "put:Put;post:Post")
	beego.Router("/v1/message/all/:pagesize", msgCtrl, "get:GetAll")
	beego.Router("/v1/message/:type/:nodeID", msgCtrl, "get:Get")*/
}

/*
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
*/
