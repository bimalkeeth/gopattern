package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type AddressSerial struct {
	StreetAddress, City, Country string
}

func (a *AddressSerial) DeepCopy() *AddressSerial {
	return &AddressSerial{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type PersonSerial struct {
	Name    string
	Address *AddressSerial
	Friends []string
}

func (p *PersonSerial) DeepCopy() *PersonSerial {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)
	fmt.Println(string(b.Bytes()))
	d := gob.NewDecoder(&b)
	result := PersonSerial{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	john := PersonSerial{"John",
		&AddressSerial{"123 London Road", "London", "UK"},
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	fmt.Println(jane, jane.Address)
}
