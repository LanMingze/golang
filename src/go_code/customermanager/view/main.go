package main
import(
	"fmt"
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
			fmt.Println("添 加 客 户")
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			fmt.Println("删 除 客 户")
		case "4":
			//fmt.Println("客 户 列 表")
			cv.list()
		case "5":
			cv.loop=false
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