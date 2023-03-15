package main

import (
	"fmt"
)

func main() {
	/*var n int 
	fmt.Println("输入层数：")
	fmt.Scanln(&n)*/
/*	*
	**
	***
	for i:=0;i<n;i++{
		for j:=0;j<=i;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}*/
	//   *
	//  ***
	// *****
	/*for i:=1;i<=n;i++{
		for k:=1;k<=n-i;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=2*i-1;j++{
			fmt.Print("*")
		}
		fmt.Print("\n")
	}*/
	//     *         3空格，1*   n=4，i=1
	//    * *        2空格，1* 1空格 1*  n=4，i=2
	//   *   *       1空格，1* 3空格 1*   n=4，i=3 
	//  *******      0空格，7*         n=4，i=4
	// for i:=1;i<=n;i++{
	// 	for k:=1;k<=n-i;k++{
	// 		fmt.Print(" ")
	// 	}
	// 	for j:=1;j<=2*i-1;j++{
	// 		if j==1||j==2*i-1||i==n{
	// 			fmt.Print("*")
	// 		}else{
	// 			fmt.Print(" ")
	// 		}
			
	// 	}
	// 	fmt.Print("\n")
	// }

	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Print(j,"*",i,"=",i*j,"\t")
		}
		fmt.Println()
	}
}
