package main
import (
	"fmt"
)
type Student struct{
	Name string
	age int
	
}
func(a *Student) Sayok(){
	fmt.Println("ok",a.Name)
}
func(a *Student) hello(){
	fmt.Println("hello",a.Name)
}
func(p *Pupil) hello(){
	fmt.Println("hello",p.Name)
}
type Pupil struct{
	Student
	Name string
}
type Brand struct{
	Name string
	Address string
}
type Goods struct{
	Name string
	Price float64
}
type TV struct{
	Goods
	Brand
}
type TV3 struct{
	*Goods
	*Brand
}
func main(){
	var pupil Pupil
	pupil.Student.Name="tom"
	pupil.Student.Sayok()
	pupil.Student.hello()
	//简化为
	pupil.Name="bob"
	pupil.age=100
	fmt.Println(pupil)
	pupil.Sayok()
	pupil.hello()
	tv:=TV{Goods{"电视机001",5000.3},Brand{"海尔","山东"}}
	tv2:=TV{
		Goods{
			Name: "电视机002",
			Price: 2000.02,
		},
		Brand{
			Name: "海尔",
			Address: "山东",
		},
	}
	fmt.Println(tv)
	fmt.Println(tv2)
	tv3:=TV3{&Goods{"电视机003",7000.3},&Brand{"创维","河南"},}
	fmt.Println(*tv3.Goods,*tv3.Brand)
	tv4:=TV3{
		&Goods{
			Name: "电视机004",
			Price: 5000.02,
		},
		&Brand{
			Name: "长虹",
			Address: "四川",
		},
	}
	fmt.Println(*tv4.Goods,*tv4.Brand)
}