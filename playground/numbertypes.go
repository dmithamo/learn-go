package main

import "fmt"

func main() {
	var i byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Println(i+1, smallI+1, bigI+1)
}
