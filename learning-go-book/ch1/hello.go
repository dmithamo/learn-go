package main

import "fmt"

func main() {
	w := "Workl"
	fmt.Println(&w)
	w += "!"
	fmt.Println(&w)
	fmt.Printf("Hello, %v\n", w)
}
