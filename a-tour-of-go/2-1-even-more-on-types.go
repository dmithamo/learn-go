package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Shorthand declaration and assignment
	a1 := 97
	a2 := 'b'
	a3 := '♬'
	var a4 byte = 'b'

	// Printing the value, unicode equivalent and type of the variable
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a1, a1, reflect.TypeOf(a1))
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a2, a2, reflect.TypeOf(a2))
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a3, a3, reflect.TypeOf(a3))
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a4, a4, reflect.TypeOf(a4))
}
