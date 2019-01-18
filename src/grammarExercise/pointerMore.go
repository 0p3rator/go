package main
// 在源码文件的第10行处加入若干代码，使该文件不出现任何编译错误，并且运行该文件会使标准输出上出现true, true
import "fmt"

type Pet interface {
    Name() string
	Age() uint8
}

type Dog struct {
    Nam string
    Ag uint8
}

func (dog Dog) Name() string {
    return dog.Nam
}

func (dog Dog) Age() uint8 {
    return dog.Ag
}
func main() {
	myDog := Dog{"Little D", 3}
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)
}