package main
// 请在命令源码文件index.go的第 8 行处添加一条语句，使该文件被执行之后在标准输出上会打印出39
import "fmt"

func main() {
    var number int = 5
	if number += 4; 10 > number {
		number += 27
		number += 3
		fmt.Print(number)
	} else if 10 < number {
		number -= 2
		fmt.Print(number)
	}
	fmt.Println(number)
}