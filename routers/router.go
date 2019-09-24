package routers

import (
	"18-mytransport/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	//静态资源处理
	beego.SetStaticPath("/static","D:\\go\\Codework\\src\\18-mytransport\\static")
	//不需要登录的页面
	beego.Router("/index", &controllers.IndexController{},"get:Index")
	beego.Router("/服务中心", &controllers.IndexController{},"get:Sever")
	beego.Router("/解决方案", &controllers.IndexController{},"get:Solve")
	beego.Router("/核心价值", &controllers.IndexController{},"get:Value")
	beego.Router("/关于中集", &controllers.IndexController{},"get:About")
	beego.Router("/关于中集", &controllers.IndexController{},"get:About")
	beego.Router("/sign", &controllers.LoginController{},"get:Login")
	beego.Router("/register", &controllers.LoginController{},"get:Register")
	//注册账户
	beego.Router("/register", &controllers.RegisterController{},"post:Register_Save")
	beego.Router("/personal-reg", &controllers.RegisterController{},"get:Personal")
	beego.Router("/enterprise-reg", &controllers.RegisterController{},"get:Enterprise")
	//用户中心
	beego.Router("/check", &controllers.UserController{},"post:Check")
	var FilterUser=func(ctx *context.Context) {
		if ctx.Input.Session("statc")!="ok"{
			ctx.Redirect(302,"/sign")
			return
		}
	}
	beego.InsertFilter("/user/*",beego.BeforeRouter,FilterUser)
	ns:=beego.NewNamespace("/user",
		beego.NSRouter("/index",&controllers.UserController{},"get:Index"),
		beego.NSRouter("/info",&controllers.UserController{},"get:Info"),
		beego.NSRouter("/order",&controllers.UserController{},"get:Order"),
		beego.NSRouter("/library",&controllers.UserController{},"get:Library"),
		beego.NSRouter("/capital",&controllers.UserController{},"get:Capital"),
		beego.NSRouter("/logout",&controllers.UserController{},"get:Logout"),
		beego.NSRouter("/order",&controllers.OrderController{},"post:Order"),
		beego.NSRouter("/info",&controllers.InfoController{},"post:EditMsg"),
		beego.NSRouter("/editpassword",&controllers.InfoController{},"post:EditPassword"),
	)
	beego.AddNamespace(ns)
}
