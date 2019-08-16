package datasource

import (
	"fmt"
	"log"
	"superstar/conf"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	masterEngine 	*xorm.Engine
	slaveEngine 	*xorm.Engine
	lock 			sync.Mutex
)

//主库  单例模式
func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		 return masterEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource )
	if err != nil {
		log.Fatal("xorm初始化数据库失败,错误=",err)
		return nil
	}
	engine.ShowSQL(true)
	engine.SetTZLocation(conf.SysTimeLocation)

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	masterEngine = engine
	return masterEngine
}

func InstanceSlave() *xorm.Engine  {
	if slaveEngine != nil {
		return slaveEngine
	}

	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource )
	if err != nil {
		log.Fatal("xorm初始化数据库失败,错误=",err)
		return nil
	}
	engine.ShowSQL(true)
	engine.SetTZLocation(conf.SysTimeLocation)

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	slaveEngine = engine
	return slaveEngine
}