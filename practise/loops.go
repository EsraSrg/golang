package main

import "fmt"

func main() {

	fruits := [3]string{"apple", "orange", "banana"}

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	// for i := 0; i < 5; i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)

}
