package main
// 请在命令源码文件index.go的第 12 和 15 行的圆括号中填入相应代码，使程序打印到标准输出上的内容为“数据已到达！”。提示，在第15行，只能使用针对通道值的某种操作。
import "fmt"

func main() {
    ch2 := make(chan string, 1)
	// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	go func() {
		ch2 <- ("已到达！")
	}()
	var value string = "数据"
	value = value + (<- ch2)
	fmt.Println(value)
}