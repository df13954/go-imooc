package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//请求,转移到这个default这个控制器处理
func (c *MainController) Get() {
	//向模板传递数据,也就是tpl中应该有对应的字段 Website  Email

	//从get方法中的URL获取传递的参数,URL的传递是在地址追加?key=value
	// 例子
	//http://localhost:8080/?name=home
	//从地址中,获取name的值,然后做处理.还有配置默认值
	name := c.GetString("name","\t dong ~~")

	c.Data["Website"] = "golang beego.me\t" + name
	c.Data["Email"] = "astaxie@gmail.com"
	//渲染的模板文件
	c.TplName = "index.tpl"
}
