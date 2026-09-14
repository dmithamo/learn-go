package main

import (
	"strings"
)

type Vertex = struct {
	lat, lng float64
}

// func main() {
// 	// fmt.Println(WordCount("Dennis is a Dennis is a dennis is good man good very very good man"))

// 	for i := 0; i < 11; i++ {
// 		fmt.Print(nthFib(i), " ")
// 	}
// 	fmt.Println()
// }

func WordCount(s string) map[string]int {
	wordMap := map[string]int{}
	for _, word := range strings.Fields(s) {
		word = strings.ToLower(word)
		if _, present := wordMap[word]; present {
			wordMap[word] += 1
		} else {
			wordMap[word] = 1
		}
	}

	return wordMap
}

func nthFib(n int) int {
	if n < 2 {
		return n
	}

	return nthFib(n-1) + nthFib(n-2)
}
