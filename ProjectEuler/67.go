//Problem 67 - Maximum path sum II - https://projecteuler.net/problem=67
//By starting at the top of the triangle below and moving to adjacent numbers on the row below,
//the maximum total from top to bottom is 23. That is, 3 + 7 + 4 + 9 = 23.
//   3
//  7 4
// 2 4 6
//8 5 9 3
//Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'),
//a 15K text file containing a triangle with one-hundred rows. NOTE: This is a much more difficult version of Problem 18.
//It is not possible to try every route to solve this problem, as there are 2^99 altogether! If you could check one
//trillion (10^12) routes every second it would take over twenty billion years to check them all. There is an efficient
//algorithm to solve it. ;o)

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	triangle := make([][]int, 100)

	file, _ := os.Open("67.txt")
	defer file.Close()
	in := bufio.NewReader(file)
	for i := 0; i < 100; i++ {
		line, _, _ := in.ReadLine()
		words := bytes.Split(line, []byte{32})
		triangle[i] = make([]int, len(words))

		for j := 0; j < len(words); j++ {
			num, err := strconv.Atoi(string(words[j]))
			if err != nil {
				fmt.Println(err)
			}
			triangle[i][j] = num
		}
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	fmt.Println(triangle[0][0])
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
