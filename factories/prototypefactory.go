package main

type EmployeeT struct {
	Name, Position string
	AnnualIncome   float32
}

const (
	Developer = iota
	Manager
)

func NewEmployeeT(role int) *EmployeeT {
	switch role {
	case Developer:
		return &EmployeeT{"", "Developer", 60000.25}
	case Manager:
		return &EmployeeT{"", "Manager", 80000.25}
	default:
		panic("Unsupported role")
	}
}

func main() {

}
