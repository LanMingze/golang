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
func (cs *CustomerService)Add(customer model.Customer)bool{
	cs.customerNum++
	customer.Id=cs.customerNum
	cs.customer=append(cs.customer, customer)
	return true
}
//根据id查找切片中的位置
func (cs *CustomerService)FindById(id int)int{
	index:=-1
	for i:=0;i<len(cs.customer);i++{
		if cs.customer[i].Id==id{
			index=i
		}
	}
	return index
}
//根据id删除切片中的该位置customer
func (cs *CustomerService)Delete(id int)bool{
	index:=cs.FindById(id)
	if index==-1{
		return false
	}else{
		cs.customer=append(cs.customer[:index],cs.customer[index+1:]... )
		return true
	}
}