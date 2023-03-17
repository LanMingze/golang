package main

import (
	"fmt"
	
)
func fibo(n int)([]uint64){
	fiboSlice:=make([]uint64,n)
	
	if n==1{
		fiboSlice[0]=1
	}else if n==2{
		fiboSlice[0]=1
		fiboSlice[1]=1
	}else if n>2{
		fiboSlice[0]=1
		fiboSlice[1]=1
		for i:=2;i<n;i++{
			fiboSlice[i]=fiboSlice[i-1]+fiboSlice[i-2]
		}
	}else{
		fmt.Println("n error")
	}	
	return fiboSlice

}
func main(){
	res:=fibo(4)
	fmt.Println(res)
}