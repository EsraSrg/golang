package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

// func main() {
// 	myslice1 := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	myslice1 = append(myslice1, 20, 21)
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// }

// func main() {
// 	arr1 := [6]int{10, 11, 12, 13, 14, 15}
// 	myslice := arr1[2:4]

// 	fmt.Printf("myslice = %v\n", myslice)
// 	fmt.Printf("length = %d\n", len(myslice))
// 	fmt.Printf("capacity = %d\n", cap(myslice))
// }

/*func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
} */
