package main

import "fmt"

func binaryfind(arr *[6]int,left int, right int,findval int){
	if left>right{
		fmt.Println("no found")
		return
	}
	middle:=(left+right)/2
	if(*arr)[middle]>findval{
		//left------middle-1
		binaryfind(arr,left,middle-1,findval)
	}else if (*arr)[middle]<findval{
		//moddle=1-----right
		binaryfind(arr,middle+1,right,findval)
	}else{
		fmt.Printf("find it index:%v\n",middle)
	}
}
func main() {
	arr:=[6]int{1,8,10,89,1000,1024}
	binaryfind(&arr,0,len(arr)-1,1024)
}
