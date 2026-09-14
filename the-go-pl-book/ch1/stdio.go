package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	sentences := make(map[string]int)

	// Use Ctrl + D to get out of .Scan
	for in.Scan() {
		sentences[in.Text()]++
	}

	for sentence, count := range sentences {
		fmt.Fprintf(os.Stderr, "%s\t%d\n", sentence, count)
	}
}
