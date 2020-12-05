package main

import "fmt"

type Emoployee struct {
	Name, Position string
	AnnualIncome   float32
}

func NewEmployeeFactory(position string, annualIncome float32) func(name string) *Emoployee {
	return func(name string) *Emoployee {
		return &Emoployee{name, position, annualIncome}
	}
}

func main() {
	developerFactory := NewEmployeeFactory("Developer", 2500.25)
	developer := developerFactory("Adam")
	fmt.Println(developer)
}
