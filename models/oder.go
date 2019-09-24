package models

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"golang.org/x/net/context"
)
type OrderMsg struct {

}
func (r *OrderMsg)SaveMsg(data OrderData,senderId interface{},receiverId interface{})interface{} {
	doc := bson.D{
		{"Username", data.Username},
		{"Type", data.Type},
		{"Time", data.Time},
		{"Amount", data.Amount},
		{"Bill", data.Bill},
		{"Sender", senderId},
		{"Receiver", receiverId},
		{"Ting", data.Ting},
	}
	collection := db.Collection("order")
	result,err:=collection.InsertOne(context.TODO(), doc)
	if err!=nil{
		fmt.Println(err)
	}

	return 	result.InsertedID

}
func (r *OrderMsg)SaveSender(data Sender)interface{} {
	doc := bson.D{
		{"senderName", data.Name},
		{"address", data.Address},
		{"addressOther", data.AddressOther},
		{"country", data.Country},
		{"city", data.City},
		{"code", data.Code},
		{"phone", data.Phone},
	}
	collection := db.Collection("Sender")
	result, err :=collection.InsertOne(context.TODO(),doc)
	if err!=nil{
		fmt.Println(err)
	}

	return 	result.InsertedID
}
func (r *OrderMsg)SaveReceiver(data Receiver)interface{} {
	doc := bson.D{
		{"receiverName", data.Name},
		{"receiverCountry", data.Country},
		{"receiverCity", data.City},
		{"receiverPhone", data.Phone},
		{"receiverAddress", data.Address},
		{"receiverCode", data.Code},

	}
	collection := db.Collection("Receiver")
	result, err :=collection.InsertOne(context.TODO(),doc)
	if err!=nil{
		fmt.Println(err)
	}

	return 	result.InsertedID
}
func (r *OrderMsg)SavetoAdmin(Id string,orderId interface{},senderId interface{},receiverId interface{})bool {
	tmpId,err := primitive.ObjectIDFromHex(Id)
	if err!=nil {
		fmt.Println(err)
		return false
	}
	//查出原有的数据
	filter := bson.D{{"_id",tmpId}}
	collection := db.Collection("admin")
	var result RegisterData
	err=collection.FindOne(context.TODO(),filter).Decode(&result)
	result.Order=append(result.Order,orderId)
	result.Sender=append(result.Sender,senderId)
	result.Receiver=append(result.Receiver,receiverId)
	//储存
	update := bson.M{"$set": bson.D{
		{"sender", result.Sender},
		{"receiver", result.Receiver},
		{"order", result.Order},
	}}
	result2, err :=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil {
		fmt.Println(err)
		return false
	}
	if result2.MatchedCount<=0{
		return false
	}
	return true
}