package main

import (
	"fmt"
	"go_code/customermanager/model"
	"go_code/customermanager/service"
)
type customerView struct{
	key string//接收用户输入
	loop bool//表示是否循环显示菜单
	cs *service.CustomerService
}
//显示所有客户信息
func (cv *customerView)list(){
	customers:=cv.cs.List()
	fmt.Println("-------客户列表--------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0;i<len(customers);i++{
		fmt.Println(customers[i].Getinfo())
	}
	fmt.Println("-------客户列表完成--------")
	fmt.Println()
}
//得到用户输入，构建用户信息，并完成添加
func (cv *customerView)add(){
	fmt.Println("-------添加客户--------")
	fmt.Println("姓名:")
	name:=""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender:=""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age:=0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone:=""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email:=""
	fmt.Scanln(&email)
	customer:=model.NewCustomer2(name,gender,age,phone,email)
	if cv.cs.Add(customer){
		fmt.Println("-------客户添加完成--------")
	}else{
		fmt.Println("-------客户添加失败--------")
	}
	
	fmt.Println()
}
//删除id对应客户
func (cv *customerView)delete(){
	fmt.Println("-------添加客户--------")
	fmt.Println("请输入id号，-1退出：")
	id:=-1
	fmt.Scanln(&id)
	fmt.Println("确认是否删除y/n：")
	choice:=""
	fmt.Scanln(&choice)
	if choice=="y"||choice=="n" {
		if id==-1{
			return
		}
		if cv.cs.Delete(id){
			fmt.Println("-------删除成功--------")
		}else{
			fmt.Println("id有误")
		}
	
	}
	
}
//退出功能
func(cv *customerView)quit(){
	fmt.Println("确认是否退出y/n：")
	for{
		fmt.Scanln(&cv.key)
		if cv.key=="y"||cv.key=="n"{
			break
		}
		fmt.Println("输入有误，请重新输入：y/n")
	}
	if cv.key=="y"{
		cv.loop=false
	}
	
}
func (cv *customerView)mainMenu(){
	for{
		fmt.Println("--------客户信息管理软件--------")
		fmt.Println("        1 添 加 客 户        ")
		fmt.Println("        2 修 改 客 户        ")
		fmt.Println("        3 删 除 客 户        ")
		fmt.Println("        4 客 户 列 表        ")
		fmt.Println("        5 退 出             ")
		fmt.Print("请选择1-5：")
		fmt.Scanln(&cv.key)
		switch cv.key{
		case "1":
			//fmt.Println("添 加 客 户")
			cv.add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			//fmt.Println("删 除 客 户")
			cv.delete()
		case "4":
			//fmt.Println("客 户 列 表")
			cv.list()
		case "5":
			cv.quit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !cv.loop{
			break
		}
	}
	fmt.Println("你已经退出系统！！！")
}
func main(){
	//创建一个customerview示例
	cv:=customerView{
		key: "",
		loop: true,
		
	}
	cv.cs=service.NewCustomerService()
	cv.mainMenu()
}