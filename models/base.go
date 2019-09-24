package models

import (
	"18-mytransport/conf"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
)


var (
	db *mongo.Database
)
func init(){
	client, err := mongo.Connect(context.TODO(),conf.DB_URL)
	if err != nil {
		panic("获得client失败")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic("连接mongodb失败")
	}
	fmt.Println("Connected to MongoDB!")
	db = client.Database(conf.DB_NAME)
}