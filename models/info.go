package models

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type InfoMsg struct {

}
//修改个人信息
func (r *InfoMsg)EditMsg(data RegisterData,id string)error {
	tmpId,err := primitive.ObjectIDFromHex(id)
	if err!=nil {
		return err
	}
	filter := bson.D{{"_id",tmpId}}
	update := bson.M{"$set": bson.D{
		{"chinaname", data.ChinaName},
		{"phone", data.Phone},
		{"weixin", data.WeiXin},
		{"qq", data.QQ},
	}}
	collection := db.Collection("admin")
	_, err =collection.UpdateOne(context.TODO(),filter,update)
	return err
}
//查找个人信息
func (r *InfoMsg)FindMsg(id string)RegisterData {
	tmpId,err := primitive.ObjectIDFromHex(id)
	if err!=nil {
		fmt.Println(err)
	}
	filter := bson.D{{"_id",tmpId}}
	collection := db.Collection("admin")
	var result RegisterData
	err=collection.FindOne(context.TODO(),filter).Decode(&result)
	if err!=nil {
		fmt.Println(err)
	}
	return result

}
//修改密码
func (r *InfoMsg)EditPassword(id string,old string,new string)bool {
	tmpId,err := primitive.ObjectIDFromHex(id)
	if err!=nil {
		fmt.Println(err)
		return false
	}
	filter := bson.D{{"_id",tmpId},{"password",old}}
	update := bson.M{"$set": bson.D{
		{"password", new},
	}}
	collection := db.Collection("admin")
	result, err :=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil {
		fmt.Println(err)
		return false
	}
	if result.MatchedCount<=0{
		return false
	}
	return true
}
