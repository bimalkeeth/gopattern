package main

type Document struct {
}
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {
	panic("implement me")
}

func (m MultiFunctionPrinter) Fax(d Document) {
	panic("implement me")
}

func (m MultiFunctionPrinter) Scan(d Document) {
	panic("implement me")
}

type OldFashionPrinter struct {
}

func (o OldFashionPrinter) Print(d Document) {
	panic("implement me")
}

func (o OldFashionPrinter) Fax(d Document) {
	panic("implement me")
}

//Deprecated:this method is deprecated
func (o OldFashionPrinter) Scan(d Document) {
	panic("Operation not supported")
}

func main() {

}
