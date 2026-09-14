package main

import (
	"fmt"
	"math"
)

const value = 20

func main() {
	nums()
	slices()

	// fmt.Println(len("😱"))
}

func nums() {
	var i = value
	var f float64 = value

	var (
		b      byte   = math.MaxUint8
		smallI uint32 = math.MaxUint32
		bigI   uint64 = math.MaxUint64
	)
	// untyped const can be assigned to more specific types
	fmt.Printf("int: %d\nfloat: %g\n", i, f)

	// 0 in all cases
	fmt.Printf("maxByte + 1: %v\nmaxInt32 + 1: %v\nmaxInt64 + 1: %v\n", b+1, smallI+1, bigI+1)
}

func slices() {
	users := []string{}
	fmt.Printf("len %v\ncap %v\n", len(users), cap(users))
	for i := range 100 {
		users = append(users, fmt.Sprintf("%d -- %v", i, "b"))
		fmt.Printf("len %v\ncap %v\n\n\n", len(users), cap(users))
	}
}
