package main

import "fmt"

func main() {

	var x = 106

	// This is a if and else block

	if x >= 6 {
		fmt.Println("Check the value again")
	} else {
		fmt.Printf("Good going!!")
	}

	// This is a for loop block

	for i := 0; i < 5; i++ {
		fmt.Printf("Shubham is great!%d\n", x)
	}

	// This is similar to while loop
	// In order to stop the loop at any stage just have a break statement

	x = 5
	for x < 10 {
		fmt.Printf("This is a similar case of while loop!")
		x++
	}

	// This is switch statement block

	switch x / 2 {
	case 0, 1, 2, 3:
		fmt.Printf("Not good")
	case 4, 5, 6:
		fmt.Printf("Ok to have progess!!")
	default:
		fmt.Printf("Try with different number!!")

	}

}
