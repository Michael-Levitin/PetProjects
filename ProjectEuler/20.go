//Problem 20 - Factorial digit sum - https://projecteuler.net/problem=20
//n! means n × (n − 1) × ... × 3 × 2 × 1
//For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
//and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
//Find the sum of the digits in the number 100!

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumFactDigits(1000))
}

func sumFactDigits(n int) int {
	m := float64(n)
	x := int(m*(math.Log10(m)-math.Log10(math.E))+math.Log10(2*math.Pi*m)/2) + 1 // 2 - just to be safe
	//fmt.Println(x)

	fact := make([]int, x)
	fact[0] = 1

	for i := 2; i <= n; i++ {
		remainder := make([]int, x)
		for j := 0; j < x; j++ { // TODO limit passes to non-zero interval
			fact[j] *= i
			for k := j; fact[j] > 0; k++ {
				remainder[k] += fact[j] % 10
				if remainder[k] > 9 { // single digits sum, hence...
					remainder[k] -= 10
					remainder[k+1]++
				}
				fact[j] /= 10
			}
		}
		fact = remainder
	}

	sum := 0
	for _, val := range fact {
		sum += val
	}
	return sum
}
