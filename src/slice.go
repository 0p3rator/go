package main
// 请在命令源码文件index.go的第 8 行和第 9 行的圆括号中分别填入一个数字，使程序打印到标准输出上的内容为true, true。
import "fmt"

func main() {
    var numbers3 = [5]int{1, 2, 3, 4, 5}
	slice3 := numbers3[2 : len(numbers3)]
	length := (3)
	capacity := (3)
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))
}