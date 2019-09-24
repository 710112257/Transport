package controllers

import (
	"18-mytransport/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
	*baseControl
}
func (c *RegisterController) Personal() {
	c.TplName = "personal-reg.html"
}
func (c *RegisterController) Enterprise() {
	c.TplName = "enterprise-reg.html"
}
func (c *RegisterController) Register_Save() {
	var data models.RegisterData
	data.Useremail=c.GetString("usermail")
	data.Username=c.GetString("username")
	data.Password=c.GetString("password")
	data.Acount="0.00"
	result:=registerMsg.SaveMsg(data)
	if result{
		c.Data["success"]="注册用户成功"
		c.Data["url"]="/sign"
		c.Data["statc"]="success"
		c.Data["msg"]="返回登录"
		c.TplName="msg.html"
	}else {
		c.Data["success"]="注册用户失败"
		c.Data["url"]="/personal-reg"
		c.Data["statc"]="wrong"
		c.Data["msg"]="重新注册"
		c.TplName="msg.html"
	}

}