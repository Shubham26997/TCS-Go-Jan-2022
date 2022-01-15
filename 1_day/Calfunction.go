package main

import "fmt"

func main() {
	var key, num1, num2 int
	key = menu()

	for {
		if key != 5 {
			switch key {
			case 1:
				fmt.Println("1. Add")
				num1, num2 = user_num()
				Add(num1, num2)

			case 2:
				fmt.Println("1. Subtact")
				num1, num2 = user_num()
				Subtract(num1, num2)

			case 3:
				fmt.Println("1. Multiply")
				num1, num2 = user_num()
				Multiply(num1, num2)

			case 4:
				fmt.Println("1. Divide")
				num1, num2 = user_num()
				Divide(num1, num2)

			default:
				fmt.Println("Enter a valid input!!")

			}
			menu()
		} else {
			fmt.Println("EXIT")
			break
		}
	}
}

func menu() int {
	fmt.Println("Welcome to the Calculator!!")
	fmt.Println("1.ADD\n2.Subtract\n3.Multiply\n4.Divide\n5.Exit")
	fmt.Printf("Enter a key = ")
	var key int
	fmt.Scanf("%d\n", &key)
	return key
}
func user_num() (int, int) {
	var num1, num2 int
	fmt.Println("Enter two numbers: ")
	fmt.Scanf("%d %d\n", &num1, &num2)
	return num1, num2
}
func Add(x, y int) (result int) {
	result = x + y
	fmt.Printf("%d\n", result)
	return
}
func Subtract(x, y int) (result int) {
	result = x - y
	fmt.Printf("%d\n", result)
	return
}
func Multiply(x, y int) (result int) {
	result = x * y
	fmt.Printf("%d\n", result)
	return
}
func Divide(x, y int) (result int) {
	result = x / y
	fmt.Printf("%d\n", result)
	return
}
