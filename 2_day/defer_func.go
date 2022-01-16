package main

import "fmt"

func main() {
	defer fmt.Println("This is main cell defer!!")
	f1()
	f2()
}

func f1() {
	// defer statement will be executed after the f1 function return
	defer func() {
		fmt.Println("This is a defer statement!!")
	}()
	fmt.Println("This is 1st function!!")

}

func f2() {
	// if we want to use only print statement then we can execute it as below
	defer fmt.Println("This is a 2nd defer statement!!")
	fmt.Println("This is 2nd function!!")
	f1()
	fmt.Println("2nd function is completed!!")
}
