package main

import "fmt"

type Student struct {
	Name string
	gender string
	age int
	id int
	score float64
}
func (student *Student)say()string{
	infostr:=fmt.Sprintf("stu name=[%v],gender=[%v],age=[%v],id=[%v],score=[%v]",
	student.Name,student.gender,student.age,student.id,student.score)
	return infostr
}

func main() {
	var stu=Student{
		Name: "tom",
		gender: "male",
		age: 20,
		id: 100,
		score: 99.25,
	}
	str:=stu.say()
	fmt.Println(str)
}
