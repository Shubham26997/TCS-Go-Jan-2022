package main

import "fmt"

type Products struct {
	ID    int
	Name  string
	Cost  float64
	Units int
}

//association

type apple struct {
	p      Products
	Expiry int
}

//composition

type mango struct {
	Products
	Expiry int
}

func main() {
	p := Products{100, "Pen", 12, 5}

	// We do require to have any predefined object of the class in order to do inheritance or references

	fruit := apple{
		p:      Products{102, "Pencil", 15, 3},
		Expiry: 5,
	}

	// We don't require to have any predefined object of the class in order to do inheritance or references

	food := mango{
		Products: Products{105, "Rubber", 45, 30},
		Expiry:   4,
	}

	fmt.Println(p.Name)

	// This will require the Product struct to be mentioned
	fmt.Println(fruit.p.Name)

	// This will not require the Product struct to be mentioned
	fmt.Println(food.Name)

	Print(p)
	Print(fruit.p)

	Print(food.Products)
	Applydiscount(&p, 0.1)
	Print(p)

	// This is a method in the struct

	p.get_Nameunit("Pen")
	fruit.p.get_Nameunit("Pencil")

	// This method do not require to have Product struct to be defined in order to acces the get_Nameunit method

	food.get_Nameunit("Rubber")

}

// We are defining the object and the class in the function as follows

func Print(p Products) {
	fmt.Printf("Id = %d, Name = %s, Cost = %f, Units = %d\n", p.ID, p.Name, p.Cost, p.Units)
}

func Applydiscount(p *Products, discount float64) {
	(*p).Cost = ((*p).Cost) * (1 - discount)
}

// We can use these function as class (struct) methods by following method:

func (p *Products) get_Nameunit(name string) {
	fmt.Printf("The Product name is %s and number of units are %d\n", (*p).Name, (*p).Units)
}
