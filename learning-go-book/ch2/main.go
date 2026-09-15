package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 0.0
	var b = 0.0
	fmt.Println("Hello, world!")
	fmt.Printf("%[1]d\n%[1]o\n%.70f\n%c\n%d\n%f\n", 0o123456, 6.023e-23, '\x61', uint64(math.MaxUint64), b/a)
}
