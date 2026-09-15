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

	fName := "Dennis"
	lName := "Mithamo"

	fmt.Printf("f_name l_name [before]%s %s\n", fName, lName)
	fName, lName = lName, fName

	fmt.Printf("f_name l_name [after]%s %s\n", fName, lName)
	const reuse = 20
	var i int64 = reuse
	var j float32 = reuse
	fmt.Printf("%d %f", i, j)
	var maxByte byte = math.MaxUint8
	var maxUint64 uint64 = math.MaxUint64
	var maxInt32 int32 = math.MaxInt32

	fmt.Printf("\n\n\n%d %d %d\n", maxByte+1, maxUint64+1, maxInt32+1)
}
