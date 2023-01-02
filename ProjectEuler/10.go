//Problem 10 - Largest product in a series - https://projecteuler.net/problem=10
//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//Find the sum of all the primes below two million.
package main

import (
	"fmt"
	"github.com/fxtlabs/primes"
	"time"
)

func main() {
	fmt.Print("Without library: ")
	t1 := time.Now()
	fmt.Println(sumPrimeToN(2000000))
	fmt.Println(time.Since(t1))

	fmt.Print("Using library  : ")
	sum := 0
	t1 = time.Now()
	s := primes.Sieve(2000000)
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	fmt.Println(sum)
	fmt.Println(time.Since(t1))
}

func sumPrimeToN(n int) int {
	pr := make([]int, n/2) // approximate Nth prime minus all even
	pr[0] = 2
	for i, next := 1, 3; i < len(pr); i, next = i+1, next+2 { // populate  with odd
		pr[i] = next
	}

	sqr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sqr[i] = i * i
	}

	for i := 1; sqr[i] < pr[len(pr)-1]; i++ { // Sieve of Eratosthenes

		count := 0
		for j := i + 1; j < len(pr); j++ {
			if pr[j]%pr[i] == 0 {
				pr[j] = 0
				count++
			}
		}

		primesTemp := make([]int, len(pr)-count)
		count = 0
		for j := 0; j < len(pr); j++ {
			if pr[j] != 0 {
				primesTemp[count] = pr[j]
				count++
			}
		}
		pr = primesTemp
	}

	sum := 0
	for i := 0; i < len(pr); i++ {
		sum += pr[i]
	}
	return sum
}
