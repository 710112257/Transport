package models

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"golang.org/x/net/context"
)

type RegisterData struct { //注册个人信息
	Id primitive.ObjectID `bson:"_id"`
	Username string
	Password string
	Useremail string
	Acount string
	Phone string
	ChinaName string
	WeiXin string
	QQ string
	Logo string
	Order []interface{}
	Sender []interface{}
	Receiver []interface{}
}

type RegisterMsg struct {

}

func (r *RegisterMsg)SaveMsg(data RegisterData)bool {
	doc := bson.D{
		{"username", data.Username},
		{"password", data.Password},
		{"useremail", data.Useremail},
		{"acount", data.Acount},
		{"phone", data.Phone},
		{"chinaName", data.ChinaName},
		{"weiXin", data.WeiXin},
		{"qq", data.QQ},
		{"logo", data.Logo},
		{"sender", data.Sender},
		{"receiver", data.Receiver},
	}
	collection := db.Collection("admin")
	_,err:=collection.InsertOne(context.TODO(), doc)
	if err!=nil {
		fmt.Println("注册用户出错:",err)
		return false
	}
	return true
}
