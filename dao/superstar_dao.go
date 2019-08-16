package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/name5566/leaf/log"
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

func (d* SuperstarDao) GetAll() []models.StarInfo {
	dataList := make([]models.StarInfo, 0)
	err := d.engine.Desc("id").Find(&dataList)
	checkError(err)
	return dataList
}

func (d* SuperstarDao) Search(country string) []models.StarInfo  {
	dataList := make([]models.StarInfo, 0)
	err := d.engine.Where("country=?",country).Desc("id").Find(&dataList)
	checkError(err)
	return dataList
}

func (d* SuperstarDao) Delete(id int) error  {
	data := &models.StarInfo{Id:id, SysStatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	checkError(err)
	return err
}

func (d* SuperstarDao) Update(data *models.StarInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	checkError(err)
	return err
}

func (d* SuperstarDao) Create(data *models.StarInfo) *models.StarInfo {
	_, err := d.engine.Insert(data)
	checkError(err)
	return data
}

func checkError(err error)  {
	if err != nil {
		log.Fatal("发生错误,err=",err)
	}
}