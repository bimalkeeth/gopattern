package main

import "fmt"

type Personx struct {
	name, position string
}

type personMod func(*Personx)
type PersonBuilderx struct {
	actions []personMod
}

func (b *PersonBuilderx) Called(name string) *PersonBuilderx {
	b.actions = append(b.actions, func(p *Personx) {
		p.name = name
	})
	return b
}

func (b *PersonBuilderx) Build() *Personx {
	p := Personx{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

// extend PersonBuilder
func (b *PersonBuilderx) WorksAsA(position string) *PersonBuilderx {
	b.actions = append(b.actions, func(p *Personx) {
		p.position = position
	})
	return b
}

func main() {
	b := PersonBuilderx{}
	p := b.Called("Dmitri").WorksAsA("dev").Build()
	fmt.Println(*p)
}
