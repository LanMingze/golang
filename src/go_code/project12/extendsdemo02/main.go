package main
import (
	"fmt"
)
type Student struct{
	Name string
	Age int
	Score int
}
type Pupil struct{
	Student//匿名结构体
}
type Graduate struct{
	Student
}
func(s *Student)show(){
	fmt.Println(s.Name,s.Age,s.Score)
}
func(s *Student)SetScore(score int){
	s.Score=score
}
func(p *Pupil)testing(){
	fmt.Println("testing pupil...")
}
func(g *Graduate)testing(){
	fmt.Println("testing graduate...")
}

func main(){
	var pupil=&Pupil{}
	pupil.Student.Name="tom"
	pupil.Student.Age=8
	pupil.testing()
	pupil.Student.show()
	pupil.Student.SetScore(85)
	pupil.Student.show()
	pupil.testing()
}