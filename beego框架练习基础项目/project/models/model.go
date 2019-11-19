package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type  User struct {
	Id int `orm:"pk;auto"`
	Name string `orm:"unique"`
	Pwd string
}

type  Article struct {
	Id int `orm:"pk;auto"`
	Title string `orm:"size(20)"`
	Type string
	ATime time.Time `orm:"auto_now_add;type(date)"`
	Acount int `orm:"default(0)"`
	Acontent string
	Aimg string
}


func init()  {
	//设置数据库连接
	orm.RegisterDataBase("default", "mysql", "root:@(localhost:3306)/article?charset=utf8")
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))
	orm.RegisterModel(new(User),new(Article))
	//生成表(别名，更改字段是否重新生成表，是否显示创建过程)
	orm.RunSyncdb("default",false,true)
}