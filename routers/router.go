// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/louisevanderlith/comment/controllers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(service *mango.Service) {
	ctrlmap := EnableFilters(service)

	msgCtrl := controllers.NewMessageCtrl(ctrlmap)
	beego.Router("v1/message", msgCtrl, "put:Put;post:Post")
	beego.Router("v1/message/:type/:nodeID", msgCtrl, "get:Get")
}

func EnableFilters(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	emptyMap["POST"] = roletype.User
	emptyMap["PUT"] = roletype.User

	ctrlmap.Add("/message", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterAPI, false)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}), false)

	return ctrlmap
}
