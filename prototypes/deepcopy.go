package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type PersonP struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *PersonP) DeepCopy() *PersonP {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main() {
	john := PersonP{"John",
		&Address{"123 London Road", "London", "UK"},
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	fmt.Println(jane, jane.Address)
}
