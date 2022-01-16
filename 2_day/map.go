package main

import "fmt"

func main() {

	dict_in_go := map[int]string{
		1: "Shubham",
		2: "Kashish",
		3: "Umesh",
		4: "Sanjana",
	}

	fmt.Println(dict_in_go[3])
	delete(dict_in_go, 4)
	fmt.Println(dict_in_go)

}
