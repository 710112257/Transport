package controllers

import (
	"18-mytransport/models"
	"github.com/astaxie/beego"
)

type InfoController struct {
	beego.Controller
}
func (c *InfoController) EditMsg() {
	var data models.RegisterData
	userid:=c.GetSession("id").(string)
	data.Phone=c.GetString("userPhone")
	data.WeiXin=c.GetString("weiXin")
	data.QQ=c.GetString("qQ")
	data.ChinaName=c.GetString("chinaName")
	err:=infoMsg.EditMsg(data,userid)
	if err!=nil{
		c.Data["success"]="账户或密码错误"
		c.Data["url"]="/user/info"
		c.Data["msg"]="返回个人信息"
		c.Data["statc"]="wrong"
		c.TplName="msg.html"
	}else {
		c.Data["success"]="修改成功"
		c.Data["url"]="/user/info"
		c.Data["statc"]="success"
		c.Data["msg"]="返回个人信息"
		c.TplName="msg.html"
	}
}
//修改密码
func (c *InfoController) EditPassword() {
	result:=infoMsg.EditPassword(c.GetSession("id").(string),c.GetString("oldPassword"),c.GetString("newPassword"))
	if result{
		c.Data["success"]="修改成功"
		c.Data["url"]="/user/info"
		c.Data["statc"]="success"
		c.Data["msg"]="返回个人信息"
		c.TplName="msg.html"

	}else {
		c.Data["success"]="旧密码错误"
		c.Data["url"]="/user/info"
		c.Data["msg"]="返回个人信息"
		c.Data["statc"]="wrong"
		c.TplName="msg.html"
	}

}