package main

import "fmt"

func main() {
	var intArr [3]int //int占8字节
	//数组已有默认值为0
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr地址是：%p,intArr[0]地址是：%p,intArr[1]地址是：%p,intArr[2]地址是：%p\n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])

	// var score [5]float64
	// for i := 0; i < len(score); i++ {
	// 	fmt.Println("请输入第", i, "个元素数据:")
	// 	fmt.Scanln(&score[i])
	// }
	// for i := 0; i < len(score); i++ {
	// 	fmt.Println("第", i, "个元素数据值为:", score[i])

	// }
	//四种数组初始化方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numarr01=", numArr01)
	var numArr02 = [3]int{3, 4, 5}
	fmt.Println("numarr02=", numArr02)
	var numArr03 = [...]int{7, 8, 9}
	fmt.Println("numarr03=", numArr03)
	var numArr04 = [...]int{1: 800, 0: 200, 2: 999}
	fmt.Println("numarr04=", numArr04)
	var numArr05 = [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("numarr05=", numArr05)

}
