package main

import "fmt"

type Emp struct {
	Id    int
	Name  string
	Basic float32
	HRA   int
	DA    int
	Tax   int
}

type FulltimeEmployee struct {
	Emp
	Benefits float32
}

func main() {
	p := FulltimeEmployee{Emp{1, "Shubham", 15000, 2000, 500, 300}, 1500}
	p.FormatSalary()
	fmt.Println("Total Salary will be: ", p.CalculatorSalary())
}

func (p Emp) FormatSalary() {
	fmt.Printf("The employee %s with id %d is having Basic salary as %f,HRA as %d, DA as %d, so TAX will be %d\n", p.Name, p.Id, p.Basic, p.HRA, p.DA, p.Tax)
}

func (p Emp) CalculatorSalary() float32 {
	var total float32
	total = float32(int(p.Basic) + p.DA + p.HRA - p.Tax)
	return total
}

func (p FulltimeEmployee) CalculatorSalary() float32 {
	var total float32
	total = p.Emp.CalculatorSalary() + p.Emp.CalculatorSalary()*0.2 + p.Benefits
	return total
}

func (p FulltimeEmployee) FormatSalary() {
	fmt.Printf("The employee %s with id %d is having Basic salary as %f,HRA as %d, DA as %d, also he will be have some benefits total of %f so TAX will be %d\n", p.Name, p.Id, p.Basic, p.HRA, p.DA, p.Benefits, p.Tax)
}
