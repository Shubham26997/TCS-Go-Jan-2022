package main

import (
	"fmt"
)

func main() {
	var x = 27
	fmt.Println(increment(x))
	fmt.Println(increment(x))
	fmt.Println(increment(x))
	fmt.Println(increment(x))
	fmt.Println(increment(x))
	fmt.Println(decrement(x))
	fmt.Println(decrement(x))
	fmt.Println(decrement(x))
	fmt.Println(decrement(x))
	fmt.Println(decrement(x))

}

func increment(x int) int {
	x += 1
	return x
}

func decrement(x int) int {
	x -= 1
	return x
}
