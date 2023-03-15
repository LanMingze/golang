package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main(){
	str:="hello你"//编码统一为utf-8，ASCII字母和数字占一个字节，汉字占3个字节
	fmt.Println("str长度是：",len(str))

	str2:="hello你好"
	r:=[]rune(str2)
	for i:=0;i<len(r);i++{
		fmt.Printf("%c\n",r[i])
	}

	n,err:=strconv.Atoi("123")
	if err!=nil{
		fmt.Println("转换错误",err)
	}else{
		fmt.Println("n=",n)
	}

	str3:=strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n",str3,str3)

	var bytes=[]byte("hello go")
	fmt.Printf("bytes=%v\n",bytes)

	str4:=string([]byte{97,98,99})
	fmt.Println(str4)

	str5:=strconv.FormatInt(123,2)
	fmt.Printf("%v\n",str5)

	b:=strings.Contains("seafood","foo")
	fmt.Println(b)

	num:=strings.Count("chesses","e")
	fmt.Println(num)
	num1:=strings.Count("chesses","s")
	fmt.Println(num1)

	fmt.Println(strings.EqualFold("abc","Abc"))


	index:=strings.Index("NLT_abc","abc")
	fmt.Println(index)
	
	index=strings.LastIndex("go golang","go")
	fmt.Println(index)

	str=strings.Replace("go go golang","go","你好",1)
	fmt.Println(str)

	strArr:=strings.Split("hello,world,ok",",")
	fmt.Println(strArr)

	fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("go"))

	str=strings.TrimSpace(" tn s losfs sdfns ")
	fmt.Printf("%q\n",str)
	str=strings.Trim(" !tn s losfs sdfns! "," !")
	fmt.Printf("%q\n",str)
	str=strings.TrimLeft(" !tn s losfs sdfns! "," !")
	fmt.Printf("%q\n",str)
	str=strings.TrimRight(" !tn s losfs sdfns! "," !")
	fmt.Printf("%q\n",str)

	b=strings.HasPrefix("ftp://172.2.15.2","ftp")
	fmt.Println(b)
	b=strings.HasSuffix("ftp://172.2.15.2ftp","ftp")
	fmt.Println(b)

}