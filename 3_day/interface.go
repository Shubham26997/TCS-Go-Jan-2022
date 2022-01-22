package main

import "fmt"

type circle struct {
	radius float32
}

type rectangle struct {
	length int
	breath int
}

type Area interface {
	area() float32
}

func main() {
	c1 := circle{25}
	r1 := rectangle{2, 5}

	// We can also use interface as they both have same function

	// fmt.Println(c1.area())
	// fmt.Println(r1.area())

	// We have used interface since both the struct is having same function area()
	Printarea(c1)
	Printarea(r1)

}

func (c circle) area() float32 {
	return (3.14 * c.radius)
}

func (r rectangle) area() float32 {
	return (float32(r.length * r.breath))
}

// To call the function or value of the interface we need to call the object of it

func Printarea(c Area) {
	fmt.Println("The area is : ", c.area())
}
