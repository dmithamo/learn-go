/**
Computers typically compute the square root of x using a loop. Starting with some guess z, we can adjust z based on how close z² is to x, producing a better guess:

z -= (z*z - x) / (2*z)
Repeating this adjustment makes the guess better and better until we reach an answer that is as close to the actual square root as can be.

Implement this in the func Sqrt provided.
A decent starting guess for z is 1, no matter what the input.
To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.

Ref: https://en.wikipedia.org/wiki/Newton%27s_method
*/

package main

import (
	"math"
)

const TOLERANCE = 0.00001
const COUNT = 10

func sqrt(num float64) float64 {
	guess := num / 2

	for ; math.Abs(guess*guess-num) > TOLERANCE; guess -= (guess*guess - num) / (2 * guess) {
	}

	return guess
}

// func main() {

// 	for i := 50.0; i < 100.0; i += 1 {
// 		fmt.Printf("%v     %v\n", i, sqrt(i))
// 	}

// }
