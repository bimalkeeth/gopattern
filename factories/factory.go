package main

import "fmt"

type PersonF interface {
	SayHello()
}

type personF struct {
	name string
	age  int
}

func (p personF) SayHello() {
	fmt.Printf("Hi , My Name is %s,I am %d years old \n", p.name, p.age)
}
func NewPersonF(name string, age int) PersonF {
	return &personF{name, age}
}

func main() {
	p := NewPersonF("James", 22)
	p.SayHello()
}
