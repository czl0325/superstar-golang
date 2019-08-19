package routes

import (
	"github.com/kataras/iris/mvc"
	"superstar/bootstrap"
	"superstar/services"
	"superstar/web/controller"
	"superstar/web/middleware"
)

func Configure(b *bootstrap.Bootstrapper)  {
	superstarService := services.NewSuperstarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controller.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controller.AdminController))
}
