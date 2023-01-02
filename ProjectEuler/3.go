//Problem 3 - Largest prime factor - https://projecteuler.net/problem=3
//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func main() {
	fmt.Println(lpf(600851475143))
}

func lpf(num int) int {
	div := 2
	for {
		if div*div > num {
			break
		}
		for {
			if num%div != 0 {
				break
			}
			num /= div
		}
		div++
	}
	return num
}
