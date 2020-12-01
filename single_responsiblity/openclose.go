package main

import "fmt"

type Color int

const( red Color=iota; green; blue)
type Size int
const (small Size=iota;medium ;large)

type Product struct {
	name string
	color Color
	size Size
}
type Filter struct {}

func (f *Filter)FilterByColor(products[]Product,color Color) []*Product{
  result:=make([]*Product,0)
  for i,v :=range products{
  	if v.color==color{
  		result=append(result,&products[i])
	}
  }
  return result
}

type Specification interface {
	IsSatisfied(p *Product)bool
}
type ColorSpecification struct {
	color Color
}
func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color==c.color
}

type SizeSpecification struct {
	size Size
}
func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size ==s.size
}

type BetterFilter struct {}
func(f *BetterFilter)Filter(products []Product,spec Specification) []*Product{
	result:=make([]*Product,0)
	for i,item:=range products{
		if spec.IsSatisfied(&item){
			result=append(result,&products[i])
		}
	}
	return result
}


func main() {
	apple:=Product{"Apple",green,small}
	tree:=Product{"Tree",green,large}
	house:=Product{"House",blue,large}
	proucts:=[]Product{apple,tree,house}

	fmt.Printf("Green products (Old) \n")
	f:=Filter{}
	for _,v :=range f.FilterByColor(proucts,green){
		fmt.Printf(" - %s green \n",v.name)
	}

	fmt.Printf("Green products (New) \n")
    greenSpec:=ColorSpecification{green}
    bf:=BetterFilter{}
    for _,v:=range bf.Filter(proucts,greenSpec){
		fmt.Printf(" - %s green \n",v.name)
	}


}
