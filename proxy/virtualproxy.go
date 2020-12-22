package main

import "fmt"

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, attack int, defense int) *Creature {
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s(%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreteModifier struct {
	creature *Creature
	next     Modifier
}

func (c CreteModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}

}

func (c CreteModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreteModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreteModifier: CreteModifier{creature: c}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreteModifier.Handle()
}

func main() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewDoubleAttackModifier(goblin)
	root.Add(NewDoubleAttackModifier(goblin))
	root.Handle()
	fmt.Println(goblin.String())
}
