package main

import "fmt"

func main() {
	names:=[4]string{"baimei","jinmao","zishan","qingyi"}
	fmt.Println("input:")
	var heroName=""
	fmt.Scanln(&heroName)
	for i:=0;i<len(names);i++{
		if heroName==names[i]{
			fmt.Printf("find %v,index %v\n",heroName,i)
			break
		}else if i==(len(names)-1){
			fmt.Printf("no found %v\n",heroName)
		}
	}

	//method 2
	index:=-1
	for i:=0;i<len(names);i++{
		if heroName==names[i]{
			fmt.Printf("find %v,index %v\n",heroName,i)
			index=i
			break
		}
	}
	if index==-1{
		fmt.Printf("no found %v\n",heroName)
	}
}
