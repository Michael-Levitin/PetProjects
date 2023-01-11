//Problem 14- Longest Collatz sequence - https://projecteuler.net/problem=14
//The following iterative sequence is defined for the set of positive integers:
//n → n/2 (n is even)
//n → 3n + 1 (n is odd)
//Using the rule above and starting with 13, we generate the following sequence:
//13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//It can be seen that this sequence (starting at 13 and finishing at 1) contains
//10 terms. Although it has not been proved yet (Collatz Problem), it is thought
//that all starting numbers finish at 1. Which starting number, under one
//million, produces the longest chain? NOTE: Once the chain starts the terms are
//allowed to go above one million.
package main

import (
	"fmt"
	"time"
)

// ToDo Why DP is slower??
func main() {
	t1 := time.Now()
	fmt.Println(Collatz(1000000))
	fmt.Println(time.Since(t1))
	t1 = time.Now()
	fmt.Println(CollatzDP(1000000))
	fmt.Println(time.Since(t1))
}

func Collatz(n int) int {
	var maxLen, maxLenNum int
	for i := 1; i < n; i++ {
		count, j := 0, i
		for j != 1 {
			if j%2 == 0 {
				j /= 2
			} else {
				j = j*3 + 1
			}
			count++
		}
		if count > maxLen {
			maxLen = count
			maxLenNum = i
			//fmt.Println("i:", i, ":count", count)
		}
	}
	return maxLenNum
}

func CollatzDP(n int) int {
	var maxLen, maxLenNum int
	collatzMap := make(map[int]int)
	collatzMap[1] = 1

	for i := 1; i < n; i++ {
		temp := helperC(i, collatzMap)
		if temp > maxLenNum {
			maxLenNum = temp
			maxLen = i
		}
	}
	return maxLen
}

func helperC(n int, collatzMap map[int]int) int {
	if _, exists := collatzMap[n]; !exists {
		if n%2 == 0 {
			collatzMap[n] = helperC(n/2, collatzMap) + 1
		} else {
			collatzMap[n] = helperC(n*3+1, collatzMap) + 1
		}
	}
	return collatzMap[n]
}
