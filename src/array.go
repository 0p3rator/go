package main
// 请在命令源码文件index.go的第11行的反引号中填入一个数字，使程序打印到标准输出上的内容为true
import "fmt"

func main() {
    var numbers2 [5]int
	numbers2[0] = 2
	numbers2[3] = numbers2[0] - 3
	numbers2[1] = numbers2[2] + 5
	numbers2[4] = len(numbers2)
	sum := (11)
	// “==”用于两个值的相等性判断
	fmt.Printf("%v\n", (sum == numbers2[0]+numbers2[1]+numbers2[2]+numbers2[3]+numbers2[4]))
}