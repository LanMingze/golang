package main

import (
	"fmt"
)

func main() {
	var scores [3][5]float64
	sum1:=0.0
	sum2:=0.0
	for i:=0;i<len(scores);i++{
		sum1=0.0
		for j:=0;j<len(scores[i]);j++{
			fmt.Printf("input %v class %v student grade:\n",i+1,j+1)
			fmt.Scanln(&scores[i][j])
			sum1+=scores[i][j]
			sum2+=scores[i][j]
		}
		fmt.Printf("%v class ave grade:%v\n",i+1,sum1/float64(len(scores[i])))
	}
	fmt.Printf("all class ave grade:%v\n",sum2/float64(len(scores)*len(scores[0])))
	fmt.Println(scores)
}
