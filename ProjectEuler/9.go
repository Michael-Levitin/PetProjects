//Problem 9 - Largest product in a series - https://projecteuler.net/problem=9
//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//a^2 + b^2 = c^2
//For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.
package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println(findPit(1000))
	fmt.Println(time.Since(t1))
}

func findPit(n int) int {
	sqr := make([]int, n+1)
	var rem int
	for i := 1; i <= n; i++ {
		sqr[i] = i * i
	}

	for i := n - 3; i > 3; i-- {
		rem = n - i
		for j := 1; j <= rem/2; j++ {
			k := rem - j
			if sqr[i] == sqr[j]+sqr[k] {
				//fmt.Println(i, j, k, i*j*k, i+j+k)
				return i * j * k
			}
		}
	}
	return 0
}
