//Problem 5 - Smallest multiple - https://projecteuler.net/problem=5
//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all the numbers from 1 to 20?
package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(scm(20))
	fmt.Println(time.Since(t1))
}

// smallest common multiplier
func scm(n int) int {
	prod := 1
	for i := 2; i <= n; i++ {
		prod = lcm(prod, i)
	}
	return prod
}

// lowest common multiplier
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
