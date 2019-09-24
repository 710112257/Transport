package controllers

import (
	"18-mytransport/models"
	"github.com/astaxie/beego"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"time"
)

type OrderController struct {
	beego.Controller
}
func (c *OrderController) Order() {
	var data models.OrderData
	data.Time=time.Now().Format("2006-01-02 15:04:05")
	if c.GetString("tranSport")=="on"{
		data.Type="EMS"
		data.Amount="120.00$"
	}else {
		data.Type="其他快递"
	}

	data.Cost="12.00$"
	data.Username=c.GetSession("username").(string)
	//
	var senderId interface{}
	if c.GetString("senderId")=="" {
		data.Sender.Name = c.GetString("senderName")
		data.Sender.Country = c.GetString("senderCountry")
		data.Sender.Phone = c.GetString("senderPhone")
		data.Sender.Address = c.GetString("senderAddress")
		data.Sender.AddressOther = c.GetString("senderOther")
		data.Sender.Code = c.GetString("senderCode")
		data.Sender.City = c.GetString("senderCity")
		//储存发送人地址
		senderId = orderMsg.SaveSender(data.Sender)
	}else {
		senderId,_=primitive.ObjectIDFromHex(c.GetString("senderId"))
	}
	//
	data.Bill.Name=c.GetString("billName")
	data.Bill.Business=c.GetString("billBusiness")
	data.Bill.Code=c.GetString("billCode")
	data.Bill.City=c.GetString("billCity")
	data.Bill.Address=c.GetString("billAddress")
	//
	var receiverId interface{}
	if c.GetString("receiverId")=="" {
		data.Receiver.Name = c.GetString("receiverName")
		data.Receiver.Country = c.GetString("receiverCountry")
		data.Receiver.Phone = c.GetString("receiverPhone")
		data.Receiver.Address = c.GetString("receiverAddress")
		data.Receiver.Code = c.GetString("receiverCode")
		data.Receiver.City = c.GetString("receiverCity")
		//储存收件人地址
		receiverId = orderMsg.SaveReceiver(data.Receiver)
	}else {
		receiverId,_=primitive.ObjectIDFromHex(c.GetString("receiverId"))
	}
	//
	data.Ting.Name=c.GetString("tingName")
	data.Ting.Weight=c.GetString("tingWeight")
	data.Ting.Volume=c.GetString("tingVolume")
	data.Ting.Pressure=c.GetString("tingPressure")
	data.Ting.Shock=c.GetString("tingShock")
	data.Ting.Moisture=c.GetString("tingMoisture")
	//储存订单
	orderId:=orderMsg.SaveMsg(data,senderId,receiverId)
	//储存到个人信息内
	result:=orderMsg.SavetoAdmin(c.GetSession("id").(string),orderId,senderId,receiverId)
	if result{
		c.Data["success"]="提交订单成功"
		c.Data["url"]="/user/order"
		c.Data["statc"]="success"
		c.Data["msg"]="返回订单列表"
		c.TplName="msg.html"
	}else {
		c.Data["success"]="提交订单失败"
		c.Data["url"]="/user/order"
		c.Data["statc"]="wrong"
		c.Data["msg"]="返回订单列表"
		c.TplName="msg.html"
	}
}