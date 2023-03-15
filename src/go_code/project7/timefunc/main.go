package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now)
	//通过now可以获得年月日时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期和时间
	fmt.Printf("当前年月日%d-%d-%d %d：%d：%d\n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	datastr := fmt.Sprintf("当前年月日%d-%d-%d %d：%d：%d\n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("data:", datastr)

	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	for j:=1;j<10;j++{
		fmt.Println(j)
		time.Sleep(time.Microsecond*100)
	}

	fmt.Printf("unix时间戳=%v unixnano时间戳=%v",now.Unix(),now.UnixNano())
}
