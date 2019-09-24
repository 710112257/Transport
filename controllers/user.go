package controllers

import (
	"18-mytransport/models"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}
func (c *UserController) Check() {
	var data models.LoginData
	data.Username=c.GetString("username")
	data.Password=c.GetString("password")
	result,sure:=loginMsg.Check(data)
	if sure{
		c.SetSession("statc","ok")
		c.SetSession("username",result.Username)
		c.SetSession("id",result.Id.Hex())
		c.Data["success"]="登录成功"
		c.Data["url"]="/user/index"
		c.Data["msg"]="点击前往个人中心"
		c.Data["statc"]="success"
		c.TplName="msg.html"
	}else {
		c.Data["success"]="账户或密码错误"
		c.Data["url"]="/sign"
		c.Data["msg"]="返回登录"
		c.Data["statc"]="wrong"
		c.TplName="msg.html"
	}

}
func (c *UserController) Index() {
	result:=loginMsg.Find(c.GetSession("id").(string))
	fmt.Println(result)
	data:=make(map[string]string)
	data["currentIndex"]="current"
	data["username"]=c.GetSession("username").(string)
	data["money"]=result.Acount
	c.Data["data"]=data
	c.TplName = "user/user-index.html"
}
func (c *UserController) Info() {
	result:=infoMsg.FindMsg(c.GetSession("id").(string))

	data:=make(map[string]string)
	data["userPhone"]=result.Phone
	data["chinaName"]=result.ChinaName
	data["weiXin"]=result.WeiXin
	data["qQ"]=result.QQ
	data["currentInfo"]="current"
	data["username"]=c.GetSession("username").(string)
	c.Data["data"]=data
	c.TplName = "user/user-info.html"
}

func (c *UserController) Order() {
	data:=make(map[string]string)
	data["currentOrder"]="current"
	data["username"]=c.GetSession("username").(string)
	c.Data["data"]=data
	c.TplName = "user/order.html"
}
func (c *UserController) Library() {
	data:=make(map[string]string)
	data["currentLibrary"]="current"
	data["username"]=c.GetSession("username").(string)
	c.Data["data"]=data
	c.TplName = "user/library-chu.html"
}
func (c *UserController) Capital() {
	data:=make(map[string]string)
	data["currentCapital"]="current"
	data["username"]=c.GetSession("username").(string)
	c.Data["data"]=data
	c.TplName = "user/capital.html"
}
func (c *UserController) Logout() {
	c.DelSession("username")
	c.DelSession("statc")
	c.Data["success"]="退出成功"
	c.Data["url"]="/index"
	c.Data["msg"]="返回首页"
	c.Data["statc"]="success"
	c.TplName="msg.html"

}