package main

import (
	"fmt"
	
)

func main() {
	var str1 string ="北京长城 123 hello world"
	fmt.Println("str1=",str1)

	//字符串是定义之后是不可变的
	//var str2="hello"
	//str2[0]='a'//这里不能修改string的内容

	//使用反引号``
	var str3 string = `saksdnadadaa""adada"`
	fmt.Println("str3=",str3)

	var str4="hello"+"world"
	str4+="haha"
	fmt.Println("str4=",str4)
	//当一行太长可以分行，+需要留在上一行
	var str5 ="hello"+"world"+
	"hahahha"+"sjdnsmcdnms"+
	"smdsmdc"
	fmt.Println("str5=",str5)

	var a int 
	var b float32
	var c float64
	var isMarried bool
	var name string
	//%v按照变量的值输出
	fmt.Printf("a=%d,b=%v,c=%v,isMArried=%v,name=%v\n",a,b,c,isMarried,name)
}
