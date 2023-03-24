package main
import (
	"fmt"
)
type Pupil struct{
	Name string
	Age int
	Score int
}
func(p *Pupil)show(){
	fmt.Println(p.Name,p.Age,p.Score)
}
func(p *Pupil)SetScore(score int){
	p.Score=score
}
func(p *Pupil)testing(){
	fmt.Println("testing...")
}

func main(){
	var pupil=&Pupil{
		Name: "tom",
		Age: 8,
	}
	pupil.show()
	pupil.SetScore(85)
	pupil.show()
	pupil.testing()
}