package main

import (
	"fmt"
	"math"
)

func main() {
	// sum := sumUptoN(10)
	// fmt.Println("Sum of pairs upto n with n", 10, sum)

	pairMinMax := minMax([]int{20, 11, 3, -2, 4, 5, 100, 200, 100_000})
	pairMinMaxToo := minMaxToo([]int{20, 11, 3, -2, 4, 5, 100, 200, 100_000})
	fmt.Println("Min max here is", pairMinMax, pairMinMaxToo)
}

// Start both algos below
// Both algos below are O(N)
//
//	The number of ops necessary with each algo
//	grows linearly with the input
func minMax(numbers []int) [2]int {
	pairMinMax := [2]int{math.MaxInt, math.MinInt}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] < pairMinMax[0] {
			pairMinMax[0] = numbers[i]
		}
		if numbers[i] > pairMinMax[1] {
			pairMinMax[1] = numbers[i]
		}
	}

	return pairMinMax
}

func minMaxToo(numbers []int) [2]int {
	pairMinMax := [2]int{math.MaxInt, math.MinInt}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] < pairMinMax[0] {
			pairMinMax[0] = numbers[i]
		}
	}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > pairMinMax[1] {
			pairMinMax[1] = numbers[i]
		}
	}
	return pairMinMax
}

// End both algos below
