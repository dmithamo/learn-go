package main

import "fmt"

func main() {
	var r rune = 'a'
	var b byte = 255
	var i int = 33000
	var s string = "hello"
	fmt.Printf("rune='%c' byte='%b' int='%d' string='%s'\n", r, b, i, s)

	for i := range int32(10) {
		fmt.Printf("%c ", (r-65)+i)
	}
}
