package main
// 在源码文件的第10行处加入若干代码，使该文件不出现任何编译错误，并且运行该文件会使标准输出上出现true, &{Little C 2 In the house}
import "fmt"

type Animal interface {
    Grow()
	Move(string) string
}
type Cat struct {
    Name string
    Age uint8
    Address string
}

func (cat *Cat) Grow(){
    cat.Age++
}

func (cat *Cat) Move(new string) (old string){
    old = cat.Address
    cat.Address = new
    return old
}


func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v\n", ok, animal)
}