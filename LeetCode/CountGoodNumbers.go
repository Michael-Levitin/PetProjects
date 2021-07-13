//https://leetcode.com/problems/count-good-numbers/
/* A digit string is good if the digits (0-indexed) at even indices are even and the digits at odd indices are prime (2, 3, 5, or 7).
For example, "2582" is good because the digits (2 and 8) at even positions are even and the digits (5 and 2) at odd positions are prime. However, "3245" is not good because 3 is at an even index but is not even. Given an integer n, return the total number of good digit strings of length n. Since the answer may be large, return it modulo 10^9 + 7.

A digit string is a string consisting of digits 0 through 9 that may contain leading zeros.*/

package main

import (
	"fmt"
	"time"
)

//a := strconv.FormatInt(n>>1, 2) // Convert to given base - string
func countGoodNumbers(n int64) int {
	m := 1000000007
	var power = 20
	var answer int
	if n&1 == 1 {
		answer = 5
	} else {
		answer = 1
	}
	for n >>= 1; n > 0; n >>= 1 {
		if n&1 != 0 {
			answer = (answer * power) % m
		}
		power = (power * power) % m
	}
	return answer
}

func fastPowMod(base int, pow int64, mod int) int { // (base^pow)%mod
	var answer int // answer = 5 to get  countGoodNumbers
	if pow&1 == 1 {
		answer = base
	} else {
		answer = 1
	} // if odd multiply by base
	// No need to convert to binary for bit shift :) !!
	for pow >>= 1; pow > 0; pow >>= 1 { // shifting n, 1 bit right (dividing by 2) == n := n>>1
		if pow&1 != 0 {
			answer = (answer * base) % mod
		} // if last digit in binary is 1
		base = (base * base) % mod // else - just multiply base
	} //(N^n)%m == ((((N%m)*N)%m *N)%m....
	return answer
}

func main() {
	timeA := time.Now()
	fmt.Println(countGoodNumbers(806166225460393))
	fmt.Println(fastPowMod(20, 806166225460393, 1000000007))
	//fmt.Println(countGoodNumbers(4)) // 400
	//fmt.Println(countGoodNumbers(50)) // 564908303
	fmt.Println(time.Since(timeA))
}
