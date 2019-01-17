package main
// 在命令源码文件index.go的第 15 行和第 18 行添加一条语句，使该文件被执行时会在标准输出上打印出：

// No Data!
// 1
// End.
import "fmt"

func main() {
    ch4 := make(chan int, 1)
	for i := 0; i < 4; i++ {
		select {
		case e, ok := <-ch4:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Println(e)
			close(ch4)
		default:
			fmt.Println("No Data!")
			ch4 <- 1
		}
	}
}