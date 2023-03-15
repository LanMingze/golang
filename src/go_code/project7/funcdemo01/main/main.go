package main

import (
	"fmt"
	util "go_code/project7/funcdemo01/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '-'
	res := util.Calculate(n1, n2, operator)
	fmt.Printf("res=%.2f", res)
}
