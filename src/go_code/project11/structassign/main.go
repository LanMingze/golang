package main

import "fmt"

type Stu struct {
	Name string
	Age int
}


func main() {
	var stu1=Stu{"xiaoming",10}
	stu2:=Stu{"tom",20}
	var stu3=Stu{
		Name: "jack",
		Age: 20,
	}
	stu4:=Stu{
		Name: "mary",
		Age: 25,
	}
	fmt.Println(stu1,stu2,stu3,stu4)
	var stu5=&Stu{"smith",12}
	stu6:=&Stu{"bob",16}
	var stu7=&Stu{
		Name: "alice",
		Age: 17,
	}
	stu8:=&Stu{
		Name: "jerry",
		Age: 18,
	}
	fmt.Println(*stu5,*stu6,*stu7,*stu8)
}
