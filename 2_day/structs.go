package main

import "fmt"

type Products struct {
	ID    int
	Name  string
	Cost  float64
	Units int
}

func main() {
	p := Products{100, "Pen", 12, 5}
	Print(p)
	Applydiscount(&p, 0.1)
	Print(p)

}

// We are defining the object and the class in the function as follows

func Print(p Products) {
	fmt.Printf("Id = %d, Name = %s, Cost = %f, Units = %d\n", p.ID, p.Name, p.Cost, p.Units)
}

func Applydiscount(p *Products, discount float64) {
	(*p).Cost = ((*p).Cost) * (1 - discount)
}
