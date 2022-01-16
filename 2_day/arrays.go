package main

import "fmt"

func main() {

	// var noarray [5]int = [5]int{2, 45, 5, 6, 39}

	noarray := [5]int{5, 4, 78, 58, 63}
	fmt.Println(noarray)

	for i, v := range noarray {
		fmt.Println(i, v)
	}

	// Only index values
	for i := range noarray {
		fmt.Println(i)
	}

	// We can also do slicing
	fmt.Println("This will print from 1 to 3 index value")
	fmt.Println(noarray[1:4])
	var newarray []int
	newarray = append(newarray, 78, 86, 23)
	fmt.Println("Length of array: ", len(newarray))
	fmt.Println("Capacity of an array: ", cap(newarray))

	// Want to fix the capacity

	n := make([]int, 0, 5)

	n = append(n, 4, 3, 6)
	fmt.Println("Length and capacity of array is: ", len(n), cap(n))
}
