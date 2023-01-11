//Problem 17 - Number letter counts - https://projecteuler.net/problem=17
//If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
//NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and
//115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
package main

import (
	"fmt"
	"time"
)

func main() {
	//numbers := make(map[int]int)
	t1 := time.Now()
	fmt.Println(CountLetters(1000))
	fmt.Println(time.Since(t1))
}

func CountLetters(n int) int {
	numbers := map[int]int{
		1: 3, 2: 3, 3: 5, 4: 4, 5: 4, 6: 3, 7: 5, 8: 5, 9: 4, 10: 3,
		11: 6, 12: 6, 13: 8, 14: 8, 15: 7, 16: 7, 17: 9, 18: 8, 19: 8, 20: 6,
		30: 6, 40: 5, 50: 5, 60: 5, 70: 7, 80: 6, 90: 6, 100: 7, 1000: 8}
	//numbers := map[int]int{
	//	1: len("one"), 2: len("two"), 3: len("three"), 4: len("four"), 5: len("five"),
	//	6: len("six"), 7: len("seven"), 8: len("eight"), 9: len("nine"), 10: len("ten"),
	//	11: len("eleven"), 12: len("twelve"), 13: len("thirteen"), 14: len("fourteen"), 15: len("fifteen"),
	//	16: len("sixteen"), 17: len("seventeen"), 18: len("eighteen"), 19: len("nineteen"), 20: len("twenty"),
	//	30: len("thirty"), 40: len("forty"), 50: len("fifty"), 60: len("sixty"),
	//	70: len("seventy"), 80: len("eighty"), 90: len("ninety"), 100: len("hundred"), 1000: len("thousand")}
	var sum int

	for i := 1; i <= n; i++ {
		temp := i
		if temp > 99 && temp%100 > 0 {
			sum += 3 // and
		}
		if temp/1000 > 0 {
			sum += numbers[1000] + numbers[temp/1000]
			temp = temp % 1000
		}
		if temp/100 > 0 {
			sum += numbers[100] + numbers[temp/100]
			temp = temp % 100
		}
		if temp >= 20 && temp <= 99 {
			sum += numbers[temp-temp%10] + numbers[temp%10]
			continue
		}
		sum += numbers[temp]
	}
	return sum
}
