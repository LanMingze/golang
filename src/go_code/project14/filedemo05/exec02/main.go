package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//创建一个新文件
	filepath := "d:/abc.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer file.Close()
	reader:=bufio.NewReader(file)
	for{
		str1,err:=reader.ReadString('\n')
		if err==io.EOF{
			break
		}
		fmt.Print(str1)
	}
	str := "hello world\n"
	//写入时使用带缓存的writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//writer带缓存，因此在调用writerstring时，内容先写入到缓存中
	//需要调用flush方法将缓存数据写入到文件中
	writer.Flush()
}
