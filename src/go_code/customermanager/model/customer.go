package model

import "fmt"

//声明一个customer结构体，表示一个客户信息
type Customer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}
func NewCustomer(id int,name string,gender string,age int,
	phone string,email string)Customer{
		return Customer{
			Id: id,
			Name: name,
			Gender: gender,
			Age: age,
			Phone: phone,
			Email: email,
		}
}
func NewCustomer2(name string,gender string,age int,
	phone string,email string)Customer{
		return Customer{
			Name: name,
			Gender: gender,
			Age: age,
			Phone: phone,
			Email: email,
		}
}
func (c Customer)Getinfo ()string{
	info:=fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",c.Id,
	c.Name,c.Gender,c.Age,c.Phone,c.Email)
	return info
}