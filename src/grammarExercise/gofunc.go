package main
// 在命令源码文件index.go中的main函数中，我编写了三条go语句。其中的go函数会分别向标准输出打印一个整数。我希望打印的最终结果是

// 1
// 2
// 3
import (
    "fmt"
)

func main() {
    ch1 := make(chan int, 1)
    ch2 := make(chan int, 1)
    ch3 := make(chan int, 3)
    go func() {
        fmt.Println("1")
        ch1 <- 1
    }()
    go func() {
        <-ch1
        fmt.Println("2")
        ch2 <- 2
    }()
    go func() {
        <-ch2
        fmt.Println("3")
        ch3 <- 3
    }()
    <-ch3
}