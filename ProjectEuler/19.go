//Problem 19 - Counting Sundays - https://projecteuler.net/problem=19
//You are given the following information, but you may prefer to do some research for yourself.
//1 Jan 1900 was a Monday.
//Thirty days has September,
//April, June and November.
//All the rest have thirty-one,
//Saving February alone,
//Which has twenty-eight, rain or shine.
//And on leap years, twenty-nine.
//A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
//How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
package main

import "fmt"

func main() {
	var count int
	date := []int{31, 12, 1899}
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	for date[2] != 2001 {
		date[0] += 7
		if date[0] > months[date[1]-1] {
			date[0] -= months[date[1]-1]
			date[1]++
		}
		if date[1] > 12 {
			date[1] -= 12
			date[2]++
			if date[2]%400 == 0 || (date[2]%100 != 0 && date[2]%4 == 0) {
				months[1] = 29
			} else {
				months[1] = 28
			}
		}

		if date[0] == 1 && date[2] > 1900 {
			//fmt.Println(date)
			count++
		}
	}
	fmt.Println(count)
}
