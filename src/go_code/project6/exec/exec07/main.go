package main

import (
	"fmt"
)

func main() {
	var sum float32 = 0
	var grade float32=0
	var total float32=0
	var class int32=3
	var student int32=5
	for i:=0;i<int(class);i++{
		sum=0
		for j:=0;j<int(student);j++{
			fmt.Printf("输入第%d个班级的第%d同学的成绩\n",i+1,j+1)
			fmt.Scanln(&grade)
			sum+=grade
			
		}
		total+=sum
		fmt.Println("班级",i+1,"平均分是：",sum/float32(student))
	}
	fmt.Println("班级总成绩是：",total,"平均分是：",total/(float32(class)*float32(student)))
}
