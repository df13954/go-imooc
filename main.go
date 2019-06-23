package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //一定要导入数据库
	_ "imooc/routers"
)

func main() {
	//第三个参数,root : 密码 @tpc(127.0.0.1:3306)/数据库名字?charset u8
	orm.RegisterDataBase("default",
		"mysql", "root:root@tcp(127.0.0.1:3306)/imooc?charset=utf8")
	beego.Run()
}
