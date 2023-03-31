package main
import "fmt"
type Account struct{
	AccountNo string
	Pwd string
	balance float64
}
func(a *Account)Deposite(money float64,pwd string){
	//存款
	if pwd!=a.Pwd{
		fmt.Println("pwd error")
		return
	}
	if money<=0{
		fmt.Println("money error")
		return
	}
	a.balance+=money
	fmt.Println("successful")
}
func (a *Account)WithDraw(money float64,pwd string){
	//取款
	if pwd!=a.Pwd{
		fmt.Println("pwd error")
		return
	}
	if money<=0||money>a.balance{
		fmt.Println("money error")
		return
	}
	a.balance-=money
	fmt.Println("successful")
}
func(a *Account)Query(pwd string){
	//查询余额
	if pwd!=a.Pwd{
		fmt.Println("pwd error")
		return
	}
	fmt.Println(a.AccountNo,a.balance)
}
func main(){
	account:=&Account{
		AccountNo: "zgyh1111",
		Pwd: "666666",
		balance: 100,
	}
	fmt.Println(*account)
	account.Deposite(200.2,"666666")
	account.Query("666666")
	account.WithDraw(100.35,"666666")
	account.Query("666666")
}