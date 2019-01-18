package main
// 在命令源码文件index.go中，我已经编写好了结构体类型MyInt及其两个方法。我希望在该文件被运行之后在标准输出上打印出true。但是好像哪里出了问题，致使标准输出上总是出现false。你能帮我找出问题所在并修正它吗？要求仅对一处做出修改
import "fmt"

type MyInt struct {
    n int
}

func (myInt *MyInt) Increase() {
	myInt.n++
}

func (myInt *MyInt) Decrease() {
	myInt.n--
}

func main() {
	mi := MyInt{}
	mi.Increase()
	mi.Increase()
	mi.Decrease()
	mi.Decrease()
	mi.Increase()
	fmt.Printf("%v\n", mi.n == 1)
}