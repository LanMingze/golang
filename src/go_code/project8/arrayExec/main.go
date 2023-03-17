package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	var mychars [26]byte
	for i:=0;i<26;i++{
		mychars[i]='A'+byte(i)
	}
	for i:=0;i<26;i++{
		fmt.Printf("%c\n",mychars[i])
	}
	
	var intArr [5]int=[5]int{1,-1,9,90,11}
	max:=intArr[0]
	index:=0
	for i:=1;i<len(intArr);i++{
		if intArr[i]>max{
			max=intArr[i]
			index=i
		}
	}
	fmt.Println("max=",max,"index=",index)

	var intArr2 [5]int=[5]int{1,-1,8,90,11}
	sum:=0
	for _,value:=range intArr2{
		sum+=value
	}
	ave:=fmt.Sprintf("%.2f",float64(sum)/float64(len(intArr2)))
	fmt.Println("sum=",sum,"ave=",ave)

	var intArr3 [5]int
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(intArr3);i++{
		intArr3[i]=rand.Intn(100)
	}
	fmt.Println(intArr3)
	for i:=0;i<len(intArr3)/2;i++{
		temp:=intArr3[i]
		intArr3[i]=intArr3[len(intArr3)-i-1]
		intArr3[len(intArr3)-i-1]=temp
	}
	fmt.Println(intArr3)
}