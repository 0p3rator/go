package main

// 请在命令源码文件index.go的第9、10、11行中填写缺失的部分。这三行与第8行和第12行共同组成了对变量num1、num2和num3的声明（注意，不是声明并赋值
import (
    "fmt"
)

func main() {
    var (
        num1 int
        num2 int
        num3 int
	)
	num1, num2, num3 = 1, 2, 3
	// 打印函数调用语句。用于打印上述三个变量的值。
	fmt.Println(num1, num2, num3)
}
