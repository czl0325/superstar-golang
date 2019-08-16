package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"superstar/services"
)

type IndexController struct {
	Ctx 		iris.Context
	Service 	services.SuperstarService
}

func (c* IndexController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.View{
		Name: "index.html",
		Data: iris.Map {
			"Title" : "球星库",
			"DataList" : dataList,
		},
	}
}

func (c* IndexController) GetBy(id int) mvc.Result  {
	return nil
}

func (c* IndexController) GetSearch() mvc.Result  {
	return nil
}