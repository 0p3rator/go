package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	m := make(map[int]Person)
	Lee := Person{
		Name: "Lee",
		Age: 20,
	}
	m[1] = Lee
	Lee.Age = 21
	m[1] = Lee
	fmt.Println(m[1].Age)
}
