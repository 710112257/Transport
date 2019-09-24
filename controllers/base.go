package controllers

import "18-mytransport/models"

var (
	registerMsg=&models.RegisterMsg{}
	loginMsg=&models.LoginMsg{}
	orderMsg=&models.OrderMsg{}
	infoMsg=&models.InfoMsg{}
)
type baseControl struct {

}