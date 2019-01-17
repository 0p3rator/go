package main
// 请在命令源码文件index.go的第 7、8、9 行中填入相应代码，使程序打印到标准输出上的内容为：25, 50, 0
import "fmt"

func main() {
    mm2 := map[string]int{"golang": 42, "java": 1, "python": 8}
    mm2["scala"] = 25
    mm2["erlang"] = 50
    delete(mm2, "python")

	fmt.Printf("%d, %d, %d \n", mm2["scala"], mm2["erlang"], mm2["python"])
}