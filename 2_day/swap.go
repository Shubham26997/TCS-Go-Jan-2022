package main

import "fmt"

func main() {
	var n1, n2 int
	n1 = 10
	n2 = 20
	fmt.Println("Before Swapping n1: ", n1, "and n2: ", n2)
	swap(&n1, &n2)
	fmt.Println("After Swapping n1: ", n1, "and n2: ", n2)
}

// "&" is for adderess and "*" is used for getting the value of that address

func swap(n1, n2 *int) {

	*n1, *n2 = *n2, *n1
}
