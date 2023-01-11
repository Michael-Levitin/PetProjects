//Problem 16 - Power digit sum - https://projecteuler.net/problem=16
//2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
//What is the sum of the digits of the number 2^1000?
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(TwoPowSum(1000))
	fmt.Println(time.Since(t1))
}

func TwoPowSum(n int) int {
	var remainder, sum int
	digits := int(float64(n)*math.Log10(2) + 1)
	power := make([]int, digits)
	power[0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < len(power); j++ {
			power[j] = power[j]*2 + remainder
			remainder = 0
			if power[j] > 9 {
				power[j] -= 10
				remainder = 1
			}
		}
	}

	for i := 0; i < len(power); i++ {
		sum += power[i]
	}
	return sum
}
