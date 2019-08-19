package models

type StarInfo struct {
	Id       int     `xorm:"not null pk autoincr comment('主键ID') INT(255)"`
	NameZh   string  `xorm:"comment('中文名') VARCHAR(255)"`
	NameEn   string  `xorm:"comment('英文名') VARCHAR(255)"`
	Avatar   string  `xorm:"comment('头像') VARCHAR(255)"`
	Birthday string  `xorm:"comment('出生日期') VARCHAR(255)"`
	Height   float64 `xorm:"comment('身高，单位cm') DOUBLE"`
	Weight   float64 `xorm:"comment('体重，单位kg
') DOUBLE"`
	Club         string `xorm:"comment('俱乐部') VARCHAR(255)"`
	Jersy        string `xorm:"comment('球衣号码以及主打位置') VARCHAR(255)"`
	Country      string `xorm:"comment('国籍') VARCHAR(255)"`
	Birthaddress string `xorm:"comment('出生地') VARCHAR(255)"`
	Feature      string `xorm:"comment('个人特点') VARCHAR(255)"`
	Moreinfo     string `xorm:"comment('更多介绍') VARCHAR(255)"`
	SysStatus    int    `xorm:"not null comment('状态，默认值 0 正常，1 删除') INT(255)"`
	SysCreated   int    `xorm:"not null comment('创建时间') INT(11)"`
	SysUpdated   int    `xorm:"not null comment('最后修改时间') INT(11)"`
}
