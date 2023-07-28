package main

import (
	"fmt"
)

func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + "World"
	return
}

func main() {
	fmt.Println(myFunction(5, "Hello"))
}

// func myFunction(x int, y int) (result int) {
// 	result = x + y
// 	return
// }

// func main() {
// 	total := myFunction(1, 2)
// 	fmt.Println(total)
// }
