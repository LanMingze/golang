package main

import (
	"fmt"
	_"math/rand"
	_"time"
)

func main() {
	fmt.Println("ok1")
	goto label
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	label:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
