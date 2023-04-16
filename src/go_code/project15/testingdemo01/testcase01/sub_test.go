package main

import (
	"fmt"
	"testing"
)

//编写测试用例
func TestGetsub(t *testing.T){
	res:=getSub(10,7)
	if res!=3 {
		t.Logf("res=%v",res)
	}else{
		fmt.Println("res=",res)
		t.Logf("successful")
	}
}