package models

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"golang.org/x/net/context"
	"log"
)
type LoginData struct {
	Id string
	Username string
	Password string
}
type LoginMsg struct {

}
func (r *LoginMsg)Check(data LoginData)(RegisterData,bool)  {
	collection := db.Collection("admin")
	matchD:=bson.M{
		"username":data.Username,
	}
	cur:=collection.FindOne(context.TODO(), matchD)
	var result=RegisterData{}
	cur.Decode(&result)
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	if result.Password!=data.Password {
		return result,false
	}else {
		return result,true
	}
}
//通过id查找
func (r *LoginMsg)Find(id string)RegisterData  {
	tmpId,err:= primitive.ObjectIDFromHex(id)
	if err!=nil{
		fmt.Println(err)
	}
	filter := bson.D{{
		"_id",
		tmpId,
	}}
	collection := db.Collection("admin")
	var result RegisterData
	err=collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
