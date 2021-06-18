//Remove Duplicates from Sorted Array
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	n := 200
	var sorted []string
	sorted = createSortedSlice(n)
	fmt.Println(removeDuplicates(sorted))

}

func createSortedSlice(n int) []string {
	sorted := make([]string, n) // making slice
	var number int
	for i := 0; i < n; i++ {
		sorted[i] = strconv.Itoa(number)

		rand.Seed(time.Now().UnixNano() + rand.Int63()) // for truly? random
		decide := rand.Intn(100)
		//fmt.Println(decide)
		if decide > 30 { // Chance to change number
			number++
		}
	}
	//	fmt.Println(sorted)
	return sorted
}

func removeDuplicates(sorted []string) ([]string, int) {
	fmt.Println(sorted)

	previouseIndex := 0
	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[previouseIndex] {
			sorted[i] = "_"
		} else if sorted[i] != sorted[previouseIndex] && sorted[previouseIndex+1] == "_" {
			previouseIndex++
			sorted[previouseIndex] = sorted[i]
			sorted[i] = "_"
		} else {
			previouseIndex = i
		}
	}
	return sorted, previouseIndex + 1
}
