package model
import "fmt"
type account struct{
	accountNo string
	pwd string
	balance float64
}
func NewAccount(accountNo string,pwd string,balance float64)*account{
	if len(accountNo)<6||len(accountNo)>10{
		fmt.Println("账号长度不对")
		return nil
	}
	if len(pwd)!=6{
		fmt.Println("密码长度不对")
		return nil
	}
	if balance<20{
		fmt.Println("余额数目不对")
		return nil
	}
	return &account{
		accountNo: accountNo,
		pwd: pwd,
		balance: balance,
	}
}
func(a *account)Deposite(money float64,pwd string){
	if pwd!=a.pwd{
		fmt.Println("pwd error")
		return
	}
	if money<=0{
		fmt.Println("money error")
		return
	}
	a.balance+=money
	fmt.Println("success")
}
func(a *account)Withdraw(money float64,pwd string){
	if pwd!=a.pwd{
		fmt.Println("pwd error")
		return
	}
	if money<=0||money>a.balance{
		fmt.Println("money error")
		return
	}
	a.balance-=money
	fmt.Println("success")
}
func(a *account)Query(pwd string){
	if pwd!=a.pwd{
		fmt.Println("pwd error")
		return
	}
	fmt.Println(a.accountNo,a.balance)
}