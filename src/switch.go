package main
// 请把命令源码文件index.go的第 10、11 行的“?”都替换为相应的表达式，以保证该文件每次被运行后都会在标准输出上打印出Case A。注意，第11行的“?”只能被替换成一个字面量。提示一下，表达式rand.Intn(4)结果会是一个范围在[0,4)的随机数
import (
    "fmt"
	"math/rand"
)

func main() {
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4) + 1]; interface{}(v).(type) {
	case int, uint, int32 :
		fmt.Printf("Case A.")
	case byte :
		fmt.Printf("Case B.")
	default:
		fmt.Println("Unknown!")
	}
}