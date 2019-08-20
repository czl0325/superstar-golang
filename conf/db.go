package conf

const DriverName = "mysql"

type DbConf struct {
	Host 	string	//主机名称
	Port	int  //端口
	User    string  //用户名
	Pwd     string  //密码
	DbName  string	//数据库名字
}

//主表参数配置
var MasterDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "123",
	DbName: "superstar",
}

var SlaveDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "123",
	DbName: "superstar",
}