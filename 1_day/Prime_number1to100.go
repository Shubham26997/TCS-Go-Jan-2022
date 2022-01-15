package main

import "fmt"

func main() {
	// We used two pointer approach and having a flag
	for i := 3; i < 100; i++ {
		isprime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isprime = false
				break
			}
		}
		if isprime == true {
			fmt.Println(i)
		}
	}

}
