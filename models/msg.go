package models

type OrderData struct {
	Id string
	Username string//订单用户
	Type string	//邮件类型EMS
	Time string	//订单时间
	Amount string	//保额
	Cost string //花费
	Bill Bill	//发票信息
	Sender Sender	//发货人信息
	Receiver Receiver	//收件人信息
	Ting Ting//商品信息
}
type Sender struct{
	Id string
	Name string
	Phone string
	Country string
	City string
	Address string
	AddressOther string
	Code string

}
type Receiver struct {
	Id string
	Name string
	Phone string
	Country string
	City string
	Address string
	Code string
}
type Bill struct {
	Name string
	Business string
	Code string
	City string
	Address string
}
type Ting struct {
	Name string
	Weight string
	Volume string
	Shock string
	Moisture string
	Pressure string
}