package utils
import "fmt"
type FamilyAccount struct{
	key string    //保存输入
	loop bool //是否退出循环
	//收支详情，使用字符串记录
	detail string
	//定义变量是否有收支
	flag bool
	//定义账户余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支说明
	note string
}
//绑定方法
//显示主菜单
func (fa *FamilyAccount)MainMenu(){
	for {
		fmt.Println("\n--------家庭收支记账软件--------")
		fmt.Println("        1.收支明细        ")
		fmt.Println("        2.登记收入        ")
		fmt.Println("        3.登记支出        ")
		fmt.Println("        4.退出软件        ")
		fmt.Print("请选择1-4：")
		fmt.Scanln(&fa.key)
		switch fa.key {
		case "1":
			fa.showDetail()			
		case "2":
			fa.income()
		case "3":
			fa.pay()
		case "4":
			fa.exit()		
		default:
			fmt.Println("输入错误...")
		}
		if !fa.loop {
			break
		}
	}
	fmt.Println("已退出！")
}
//显示主菜单
func (fa *FamilyAccount)showDetail(){
	fmt.Println("--------当前收支明细--------")
	if fa.flag{
		fmt.Println(fa.detail)
	}else{
		fmt.Println("当前没有收支，请来一笔吧！")
	}
}
//登记收入
func (fa *FamilyAccount)income(){
	fmt.Println("本次收入金额：")
	fmt.Scanln(&fa.money)
	fa.balance += fa.money //修改余额
	fa.flag=true
	fmt.Println("本次收入说明：")
	fmt.Scanln(&fa.note)
	//将收入情况记录到detail中
	fa.detail += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", fa.balance,fa.money, fa.note)
}
//登记支出
func (fa *FamilyAccount)pay(){
	fmt.Println("--------登记支出--------")
	fmt.Println("本次支出金额：")
	fmt.Scanln(&fa.money)
	if fa.money>fa.balance{
		fmt.Println("余额不足")
	}
	fa.balance -= fa.money //修改余额
	fa.flag=true
	fmt.Println("本次收入说明：")
	fmt.Scanln(&fa.note)
	fa.detail += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", fa.balance, fa.money, fa.note)
}
//退出软件
func (fa *FamilyAccount)exit(){
	fmt.Println("你确定要退出吗？ y/n")
	choice:=""
	for{
		fmt.Scanln(&choice)
		if choice=="y"||choice=="n"{
			break
		}
		fmt.Println("输入有误，重新输入")
	}
	if choice=="y"{
		fa.loop = false
	}
}

func NewFamilyAccount()*FamilyAccount{
	return &FamilyAccount{
		key: "",
		loop: true,
		balance: 100000.0,
		money: 0.0,
		note: "",
		flag: false,
		detail: "收支\t账户余额\t收支金额\t说明",
	}
}