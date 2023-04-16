package main

import (
	"fmt"
	"testing"
)

//编写测试用例
func TestAddUpper(t *testing.T){
	res:=AddUpper(10)
	if res!=55 {
		t.Logf("res=%v",res)
	}else{
		fmt.Println("res=",res)
		t.Logf("successful")
	}
}