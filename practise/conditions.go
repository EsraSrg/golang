package main

import (
	"fmt"
)

func main() {
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

// func main() {
// 	day := 8

// 	switch day {
// 	case 1:
// 		fmt.Println("Monday")
// 	case 2:
// 		fmt.Println("Tuesday")
// 	case 3:
// 		fmt.Println("Wednesday")
// 	case 4:
// 		fmt.Println("Thursday")
// 	case 5:
// 		fmt.Println("Friday")
// 	case 6:
// 		fmt.Println("Saturday")
// 	case 7:
// 		fmt.Println("Sunday")
// 	default:
// 		fmt.Println("Not a weekday")
// 	}
// }

// func main() {
// 	x := 30
// 	if x >= 10 {
// 		fmt.Println("x is larger than or equal to 10.")
// 	} else if x > 20 {
// 		fmt.Println("x is larger than 20.")
// 	} else {
// 		fmt.Println("x is less than 10.")
// 	}
// }
