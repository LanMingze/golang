package main

import "fmt"

func main() {
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0
	total := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	ave := fmt.Sprintf("%.2f", total/6) //返回小数点后两位,6为常量不需要转为float64
	fmt.Printf("ave=%v\n", ave)

	//定义一个数组
	var hens [6]float64
	//给数组初始化
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	sum := 0.0
	for i := 0; i < len(hens); i++ {
		sum += hens[i]
	}
	avg := fmt.Sprintf("%.2f", sum/float64(len(hens)))
	fmt.Println("avg=", avg)

}
