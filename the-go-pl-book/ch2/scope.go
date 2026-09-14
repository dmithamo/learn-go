package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name := "The quick brown fox jumped over a lazy Dog!"
	lower := ""
	upper := ""

	for _, char := range name {
		lower += fmt.Sprintf("%c", toLower(char))
		upper += fmt.Sprintf("%c", toUpper(char))
	}

	fmt.Println(upper)
	fmt.Println(lower)
}

func toUpper(c int32) int32 {
	if !isChar(c) {
		return c
	}

	if c > 'Z' {
		return c + 'A' - 'a'
	}

	return c
}

func toLower(c int32) int32 {
	if !isChar(c) {
		return c
	}

	if c > 'Z' {
		return c
	}

	return c - 'A' + 'a'
}

func isChar(c int32) bool {
	return c >= 'A' && c <= 'z'
}

var cwd string
var err error

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}

	fmt.Printf("cwd: %s\n", cwd)
	fmt.Printf("%v\n\n", 1e6)
}
