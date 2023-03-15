package main

import "fmt"

func main() {
	var heroes [3]string=[3]string{"宋江","吴用","卢俊义"}
	for i,v:=range heroes{
		fmt.Printf("i=%v,v=%v\n",i,v)
	}//不需要i时，可以用_占位

}
