//Problem 7 - 10001st prime - https://projecteuler.net/problem=7
//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//What is the 10 001st prime number?
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t1 := time.Now()
	printNthPrime(10001)
	fmt.Println(time.Since(t1))
}

func printNthPrime(n float64) {
	primes := make([]int, int(n*math.Ceil(math.Logb(n))/2)) // approximate Nth prime minus all even
	primes[0] = 2

	for i, next := 1, 3; i < len(primes); i, next = i+1, next+2 { // populate  with odd
		primes[i] = next
	}

	for j := 1; j*j < primes[len(primes)-1]; j++ { // Sieve of Eratosthenes
		for i := j + 1; i < len(primes); i++ {
			if primes[i]%primes[j] == 0 {
				primes = append(primes[:i], primes[i+1:]...)
			}
		}
	}
	fmt.Println(primes[int(n)-1])
}
