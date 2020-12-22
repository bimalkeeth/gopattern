package main

import "fmt"

type Person struct {
	FirstNme, MiddleName, LastName string
}

func (p *Person) Names() []string {
	return []string{p.FirstNme, p.MiddleName, p.LastName}
}

func (p *Person) NameGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstNme
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

func main() {
	p := Person{"Alexander", "Graham", "Bell"}
	for name := range p.NameGenerator() {
		fmt.Println(name)
	}
}
