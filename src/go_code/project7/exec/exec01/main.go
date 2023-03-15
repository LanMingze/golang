package main

import "fmt"

func fibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}
func main() {
	n := 5
	fb := fibo(n)
	fmt.Println(n, "的Fibonacci是：", fb)
	res := f(n)
	fmt.Println("res=", res)
}
