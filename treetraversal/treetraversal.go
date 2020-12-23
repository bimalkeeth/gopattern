package main

import "fmt"

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewNode(value int, left *Node, right *Node) *Node {
	n := &Node{Value: value, Left: left, Right: right}
	left.Parent = n
	right.Parent = n
	return n
}

func NewTerminalNode(value int) *Node {
	return &Node{Value: value}
}

type InOrderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	i := &InOrderIterator{root, root, false}
	for i.Current.Left != nil {
		i.Current = i.Current.Left
	}
	return i
}
func (i *InOrderIterator) Reset() {
	i.Current = i.root
	i.returnedStart = false
}
func (i *InOrderIterator) MoveNext() bool {
	if i.Current == nil {
		return false
	}
	if !i.returnedStart {
		i.returnedStart = true
		return true
	}

	if i.Current.Right != nil {
		i.Current = i.Current.Right
		for i.Current.Left != nil {
			i.Current = i.Current.Left
		}
		return true
	} else {
		p := i.Current.Parent
		for p != nil && i.Current == p.Right {
			i.Current = p
			p = p.Parent
		}
		i.Current = p
		return i.Current != nil
	}

}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}

func (b *BinaryTree) InOrder() *InOrderIterator {
	return NewInOrderIterator(b.root)
}

func main() {
	root := NewNode(1,
		NewTerminalNode(2),
		NewTerminalNode(3))
	it := NewInOrderIterator(root)

	for it.MoveNext() {
		fmt.Printf("%d,", it.Current.Value)
	}
	fmt.Println("\b")

	t := NewBinaryTree(root)
	for i := t.InOrder(); i.MoveNext(); {
		fmt.Printf("%d,", i.Current.Value)
	}
	fmt.Println("\b")
}
