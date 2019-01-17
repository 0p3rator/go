package main
// 为源码文件中声明的结构体类型Person添加相应的字段和方法，使得该文件不会导致任何编译错误并能够在标准输出上打印出Robert moved from Beijing to San Francisco.
import "fmt"

type Person struct {
    Name    string
	Gender  string
	Age     uint8
	Address string
}
func (person *Person) Move(newAddress string) string{
    oldAddress := person.Address
    person.Address = newAddress
    return oldAddress
}


func main() {
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}