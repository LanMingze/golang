package service

import (
	"go_code/customermanager/model"
)
//增删该查
type CustomerService struct{
	customer []model.Customer
	//表示当前切片含有多少个客户,还可以作为新客户的编号
	customerNum int
}
func NewCustomerService()*CustomerService{
	cs:=&CustomerService{

	}
	cs.customerNum=1
	customer:=model.NewCustomer(1,"张三","男",20,"11021","sjefdhas@qq.com")
	cs.customer=append(cs.customer, customer)
	return cs
}
//返回客户切片
func (cs *CustomerService)List()[]model.Customer{
	return cs.customer
}