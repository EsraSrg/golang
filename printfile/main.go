package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	//Args is an array of string that contains all the
	//command line arguments passed.
	os.Args[1]
}
