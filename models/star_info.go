package models

type StarInfo struct {
	Id       int     `xorm:"not null pk autoincr comment('主键ID') INT(255)" form:"id"`
	NameZh   string  `xorm:"comment('中文名') VARCHAR(255)" form:"name_zh"`
	NameEn   string  `xorm:"comment('英文名') VARCHAR(255)" form:"name_en"`
	Avatar   string  `xorm:"comment('头像') VARCHAR(255)" form:"avatar"`
	Birthday string  `xorm:"comment('出生日期') VARCHAR(255)" form:"birthday"`
	Height   float64 `xorm:"comment('身高，单位cm') DOUBLE" form:"height"`
	Weight   float64 `xorm:"comment('体重，单位kg
') DOUBLE" form:"weight"`
	Club         string `xorm:"comment('俱乐部') VARCHAR(255)" form:"club"`
	Jersy        string `xorm:"comment('球衣号码以及主打位置') VARCHAR(255)" form:"jersy"`
	Country      string `xorm:"comment('国籍') VARCHAR(255)" form:"country"`
	Birthaddress string `xorm:"comment('出生地') VARCHAR(255)" form:"birthaddress"`
	Feature      string `xorm:"comment('个人特点') VARCHAR(255)" form:"feature"`
	Moreinfo     string `xorm:"comment('更多介绍') VARCHAR(255)" form:"moreinfo"`
	SysStatus    int    `xorm:"not null comment('状态，默认值 0 正常，1 删除') INT(255)" form:"-"`
	SysCreated   int    `xorm:"not null comment('创建时间') INT(11)" form:"-"`
	SysUpdated   int    `xorm:"not null comment('最后修改时间') INT(11)" form:"-"`
}
