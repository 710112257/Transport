package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}
func (c *IndexController) Index() {
	data:=make(map[string]interface{})
	data=touxiang(c,data)
	data["currentIndex"]="current"
	c.Data["data"]=data
	c.TplName = "index.html"
}
func (c *IndexController) Sever() {
	data:=make(map[string]interface{})
	data=touxiang(c,data)
	data["currentSever"]="current"
	c.Data["data"]=data
	c.TplName = "服务中心.html"
}
func (c *IndexController) Solve() {
	data:=make(map[string]interface{})
	data=touxiang(c,data)
	data["currentSolve"]="current"
	c.Data["data"]=data
	c.TplName = "解决方案.html"
}
func (c *IndexController) Value() {
	data:=make(map[string]interface{})
	data=touxiang(c,data)
	data["currentValue"]="current"
	c.Data["data"]=data
	c.TplName = "核心价值.html"
}
func (c *IndexController) About() {
	data:=make(map[string]interface{})
	data=touxiang(c,data)
	data["currentAbout"]="current"
	c.Data["data"]=data
	c.TplName = "关于中集.html"
}
func touxiang(c *IndexController,data map[string]interface{})map[string]interface{}{
	username:=c.GetSession("username")
	if username!=nil{
		data["username"]=username
		data["url"]="/user/index"
	}else {
		data["username"]="登录"
		data["url"]="/sign"
	}
	return data
}

