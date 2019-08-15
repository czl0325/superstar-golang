package dao

import (
	"github.com/go-xorm/xorm"
	"superstar/models"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {
	return &SuperstarDao{engine:engine}
}

func (d* SuperstarDao) Get (id int) *models.StarInfo {
	data := &models.StarInfo{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}