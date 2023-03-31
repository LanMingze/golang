package main

import (
	"fmt"
)
type Monkey struct{
	Name string
}
//声明接口
type BirdAble interface{
	Flying()
}
type FishAble interface{
	Swimming()
}
func(m *Monkey)climbing(){
	fmt.Println(m.Name,"猴子会爬树。。。")
}
type LittleMonkey struct{
	Monkey
}
func(lm *LittleMonkey)Flying(){
	fmt.Println(lm.Name,"猴子会飞。。。")
}
func(lm *LittleMonkey)Swimming(){
	fmt.Println(lm.Name,"猴子会游泳。。。")
}
func main(){
	monkey:=LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}