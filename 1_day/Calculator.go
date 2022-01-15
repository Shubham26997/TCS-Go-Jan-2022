package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Calculator!!")
	fmt.Println("1.ADD\n2.Subtract\n3.Multiply\n4.Divide\n5.Exit")
	fmt.Printf("Enter a key = ")
	var key int
	// With Scanf add \n at the end also with %d type verbs use Printf
	fmt.Scanf("%d\n", &key)
	var num1, num2 int
	var result int
	for {
		if key != 5 {
			switch key {
			case 1:
				fmt.Println("Enter two numbers")
				fmt.Scanf("%d %d\n", &num1, &num2)
				result = num1 + num2
				fmt.Printf("%d+%d = %d\n", num1, num2, result)

			case 2:
				fmt.Println("Enter two numbers")
				fmt.Scanf("%d %d\n", &num1, &num2)
				result = num1 - num2
				fmt.Printf("%d-%d = %d\n", num1, num2, result)
			case 3:
				fmt.Println("Enter two numbers")
				fmt.Scanf("%d %d\n", &num1, &num2)
				result = num1 * num2
				fmt.Printf("%d*%d = %d\n", num1, num2, result)
			case 4:
				fmt.Println("Enter two numbers")
				fmt.Scanf("%d %d\n", &num1, &num2)
				result = num1 / num2
				fmt.Printf("%d/%d = %d\n", num1, num2, result)

			default:
				fmt.Println("Enter a valid Key!!")

			}
			fmt.Println("1.ADD\n2.Subtract\n3.Multiply\n4.Divide\n5.Exit")
			fmt.Printf("Enter a key = ")
			fmt.Scanf("%d\n", &key)
		} else {
			fmt.Println("EXIT")
			break
		}
	}

}
