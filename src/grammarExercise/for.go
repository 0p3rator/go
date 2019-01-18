package main
// 请把命令源码文件index.go的第 9 行写一条for语句，使该文件被执行时会在标准输出上打印出：

// 1: Golang
// 2: Java
// 3: Python
// 4: C
import (
    "fmt"
)

func main() {
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
    for i, _ := range map1 {
        fmt.Printf("%d: %s\n", i, map1[i])
    }
}