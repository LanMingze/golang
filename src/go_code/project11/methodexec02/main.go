package main

import "fmt"

type MethodUtils struct{
	
	
}
func (m MethodUtils)Print(){
	for i:=1;i<=10;i++{
		for j:=1;j<=8;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
} 
func (me MethodUtils)Print2(m int,n int){
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
} 
func (me MethodUtils)area(m float64,n float64)float64{
	return m*n
} 
func (me *MethodUtils)judge(num int){
	if num%2==0{
		fmt.Println("oushu")
	}else{
		fmt.Println("jishu")
	}
}
func (me *MethodUtils)print3(m int,n int,s string){
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			fmt.Print(s)
		}
		fmt.Println()
	}
}
type Calculator struct{
	Num1 float64
	Num2 float64
}
func(cal *Calculator)getRes(operator byte)float64{
	res:=0.0
	switch operator{
	case '+':
		res=cal.Num1+cal.Num2
	case '-':
		res=cal.Num1-cal.Num2
	case '*':
		res=cal.Num1*cal.Num2
	case '/':
		res=cal.Num1/cal.Num2
	default:
		fmt.Println("error")
	}
	return res
}
func main(){
	var m MethodUtils
	m.Print()
	m.Print2(2,4)
	res:=m.area(5.5,6.2)
	fmt.Println(res)
	(&m).judge(11)
	(&m).print3(5,3,"%")
	var cal Calculator
	cal.Num1=4.5
	cal.Num2=6.3
	res=(&cal).getRes('*')
	fmt.Println(res)
}