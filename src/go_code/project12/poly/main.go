package main
import (
	"fmt"
)
type Usb interface{
	Start()
	Stop()
}
type Phone struct{
	name string
}
//手机实现usb接口
func(p *Phone)Start(){
	fmt.Println("手机开始工作")
}
func(p *Phone)Stop(){
	fmt.Println("手机停止工作")
}
type Camera struct{
	name string
}
//相机实现usb接口
func(c *Camera)Start(){
	fmt.Println("相机开始工作")
}
func(c *Camera)Stop(){
	fmt.Println("相机停止工作")
}
type Computer struct{

}
func(c *Computer) Working(usb Usb){//接收usb接口类型的变量，
	//只要实现usb接口的都可以传入，即实现接口声明的所有方法
	usb.Start()
	usb.Stop()
}
func main(){
	//定义一个接口数组，可以存放phone和camera变量
	var usbArr [3]Usb
	fmt.Println(usbArr)
	usbArr[0]=&Phone{"xiaomi"}
	usbArr[1]=&Phone{"vivo"}
	usbArr[2]=&Camera{"nikang"}
	fmt.Println(usbArr[0],usbArr[1],usbArr[2],)
}