package main

import (
	"fmt"
	_"math/rand"
	_"time"
)

func main() {
	// var count int=0
	// rand.Seed(time.Now().UnixNano())
	// for{

		
	// 	//为了生成随机数需要给rand生成种子
		
	// 	n := rand.Intn(100) + 1
	// 	count++
	// 	if n==99{
	// 		break
	// 	}
	// }
	// fmt.Println("生成99一个使用了",count,"次")
	lable1:
	for i:=0;i<4;i++{
		//lable2:
		for j:=0;j<10;j++{
			if j==2{
				break lable1
			}
			fmt.Println("j=",j)
		}
	}
}
