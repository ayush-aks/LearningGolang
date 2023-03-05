package main

import (
	"fmt"
)

func main() {
	n, f := genSlice()
	if n%2 == 0 {
		fmt.Println("Number is even")
		f = 0
	} else {
		fmt.Println("Number is odd")
		f = 1
	}
	_ = f
}

func genSlice() (int, int) {
	var n int
	fmt.Scanln(&n)
	var fl int = 1
	return n, fl
}
