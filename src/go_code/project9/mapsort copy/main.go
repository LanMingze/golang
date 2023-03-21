package main

import (
	"fmt"
	"sort"
)

func main() {
	map1:=make(map[int]int,11)
	map1[10]=100
	map1[1]=13
	map1[8]=9
	map1[4]=20
	fmt.Println(map1)

	var keys []int
	for k:=range map1{
		keys=append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _,k:=range keys{
		fmt.Printf("map[%v]=%v\n",k,map1[k])
	}
}
