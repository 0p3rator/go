package main
// 命令源码文件index.go中代码的功能是打印出斐波那契数列的前10个数，即：0 1 1 2 3 5 8 13 21 34 。请把第 9 行的fmt.Printf函数调用语句修改成一条defer语句，使得该文件被执行之后标准输出上会出现：

// 0 1 1 2 3 5 8 13 21 34 34 21 13 8 5 3 2 1 1 0  
//     显然，这是关于斐波那契数列的一段回文。提示一下，这需要利用前面讲到的defer语句的几个特殊行为。尤其是最后的“特别提示”中讲到的一些内容。
import (
    "fmt"
)

func main() {
	for i := 0; i < 10; i++ {
	    defer func(n int){
	        fmt.Printf("%d ", fibonacci(n))
	    }(i)
		fmt.Printf("%d ", fibonacci(i))
	}
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}