package routers

import (
	"github.com/louisevanderlith/comment/controllers"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	//Message

	e.JoinBundle("/", roletype.User, mix.JSON, &controllers.Types{})
	e.JoinBundle("/", roletype.Admin, mix.JSON, &controllers.Messages{})
	/*
		msgCtrl := &controllers.Messages{}
		msgGroup := routing.NewRouteGroup("message", mix.JSON)
		msgGroup.AddRoute("Create Message", "", "POST", roletype.User, msgCtrl.Post)
		msgGroup.AddRoute("Update Message", "", "PUT", roletype.User, msgCtrl.Put)
		msgGroup.AddRoute("All Messges", "/all/{pagesize:[A-Z][0-9]+}", "GET", roletype.Admin, msgCtrl.GetAll)
		msgGroup.AddRoute("Get Comments for Item", "/{type}/{nodeID:[0-9]+\x60[0-9]+}", "GET", roletype.Unknown, msgCtrl.Get)
		e.AddBundle(msgGroup)*/
}
