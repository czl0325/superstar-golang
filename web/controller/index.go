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
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map {
			"Title" : "球星库",
			"Data" : data,
		},
	}
}

func (c* IndexController) GetSearch() mvc.Result  {
	country := c.Ctx.URLParam("country")
	dataList := c.Service.Search(country)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map {
			"Title" : "球星库",
			"DataList" : dataList,
		},
	}
}