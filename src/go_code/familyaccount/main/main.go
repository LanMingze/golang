package main

import (
	"fmt"
	"go_code/familyaccount/utils"
)

func main() {
	fmt.Println("面向对象方式")
	utils.NewFamilyAccount().MainMenu()
	
}
