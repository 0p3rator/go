package main
//  命令源码文件index.go中有三个函数：main、outerFunc和innerFunc。在innerFunc函数中，我准备通过调用panic函数引发一个运行时恐慌。请你在文件中的适当位置添加一条defer语句，以使得该运行时恐慌得到“恢复”。并且，该文件的运行应该使标准输出上打印出：

// Enter main
// Enter outerFunc
// Enter innerFunc
// Fatal error: Occur a panic!
import (
    "errors"
	"fmt"
)

func innerFunc() {
	fmt.Println("Enter innerFunc")
	panic(errors.New("Occur a panic!"))
	fmt.Println("Quit innerFunc")
}

func outerFunc() {
	fmt.Println("Enter outerFunc")
	innerFunc()
	fmt.Println("Quit outerFunc")
}

func main() {
	fmt.Println("Enter main")
	defer func(){
	    if p := recover(); p!= nil {
	        fmt.Printf("Fatal error: %s\n", p)
	    }
	}()
	outerFunc()
	fmt.Println("Quit main")
}