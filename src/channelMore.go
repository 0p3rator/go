package main
// 请在命令源码文件index.go的第 13 行的内层圆括号中填入一个正整数，使程序打印到标准输出上的内容为：

// Received! 6
// Sent!
import (
    "fmt"
	"time"
)

type Sender chan<- int

type Receiver <-chan int

func main() {
	var myChannel = make(chan int, (0))
	var number = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		fmt.Println("Received!", <-receiver)
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)
}