// Create a program that reads the contents of a text file then prints its contents to the terminal.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	a := os.Args
	if len(a) < 2 {
		fmt.Println("File name was not provided, Exiting...")
		os.Exit(1)
	}
	f := a[1]
	fi, err := os.Open(f)
	if err != nil {
		fmt.Println("Error in opening file: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, fi)
}
