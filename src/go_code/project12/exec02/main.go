package main

import (
	"fmt"
	"math/rand"
	"sort"
)
type Hero struct{
	Name string
	Age int
}
type HeroSlice []Hero
func (hs HeroSlice)Len()int{
	return len(hs)
}
func (hs HeroSlice)Less(i,j int)bool{
	if hs[i].Age<hs[j].Age{
		return true 
	}else{
		return false
	}
		
}
func (hs HeroSlice)Swap(i,j int){
	// temp:=hs[i]
	// hs[i]=hs[j]
	// hs[j]=temp
	//等价于
	hs[i],hs[j]=hs[j],hs[i]
		
}
func main(){
	var intSlice=[]int{0,-1,10,7,90}
	sort.Ints(intSlice)//排序
	fmt.Println(intSlice)

	var heros HeroSlice
	for i:=0;i<10;i++{
		hero:=Hero{
			Name: fmt.Sprintf("英雄%d",rand.Intn(100)),
			Age: rand.Intn(100),
		}
		heros=append(heros,hero)
	}
	for _,v:=range heros{
		fmt.Println(v)
	}
	sort.Sort(heros)
	fmt.Println("heros排序后：")
	for _,v:=range heros{
		fmt.Println(v)
	}
}