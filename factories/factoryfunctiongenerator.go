package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   float32
}

func NewEmployeeFactory(position string, annualIncome float32) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

func main() {
	developerFactory := NewEmployeeFactory("Developer", 2500.25)
	developer := developerFactory("Adam")
	fmt.Println(developer)
}
