// is sum a power of 2 - count sums
//https://leetcode.com/problems/count-good-meals/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func countPairs(deliciousness []int) int {
	var count int
	numsMap := make(map[int]int)
	for _, v := range deliciousness {
		if _, exist := numsMap[v]; exist {
			numsMap[v]++
		} else {
			numsMap[v] = 1
		}
	}
	//fmt.Println(numsMap)

	keys := make([]int, 0, len(numsMap)) // contains keys of map
	for k := range numsMap {
		keys = append(keys, k)
	}

	secondPairs := make(map[int]int) // contains second pair - to eliminate dupe calc
	sort.Ints(keys)                  // Sorting the slice of keys
	//fmt.Println(keys)
	for i := 0; i <= 21; i++ {
		//target:= int(math.Pow(2, float64(i)))
		target := 1 << i
		for _, key := range keys {
			if key > target {
				break
			}

			if key == target-key { // nCr calculation
				//nCr = n!/(r!*(n-r)!), for r = 2 => n*(n-1)/2
				count += numsMap[key] * (numsMap[key] - 1) / 2
				//fmt.Println("Count 1:", count, ", Key1 = ", key, ", Key2 = ", target - key, ", Target = ", target)
			} else if value, exist := numsMap[target-key]; exist {
				if valuePair, existPair := secondPairs[key]; existPair &&
					valuePair == target-key {
					continue // if already counted that pair
				} // otherwise count pair and add to pairs
				secondPairs[target-key] = key
				count += numsMap[key] * value
				//fmt.Println("Count 2:", count, ",Key1 = ", key, ",Key2 = ", target - key, ",Target = ", target)
			}
		}
	}
	return count % (1e9 + 7) // problem statement
}

// for individual runs
func decodeJson(output interface{}, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the file", err)
	}
	// using file and address of var. var must be same structure as Json
	if err := json.NewDecoder(file).Decode(&output); err == nil {
		fmt.Println("Decode successful")
	} else {
		fmt.Println("Decode unsuccessful", err)
	}
	file.Close()
	//fmt.Println(output)
} // usage decodeJson(&someVar, "filename.ext")

func main() {
	fmt.Println(countPairs([]int{1, 3, 5, 7, 9}))
	fmt.Println(countPairs([]int{1, 1, 1, 3, 3, 3, 7}))
	start := time.Now()
	//array := []int{2160,1936,3,29,27,5,2503,1593,2,0,16,0,3860,28908,6,2,15,49,6246,1946,23,105,7996,196,0,2,55,457,5,3,924,7268,16,48,4,0,12,116,2628,1468}
	var array []int
	decodeJson(&array, "CountGoodMeals.json")
	//array := []int{1048576,1048576}
	//fmt.Println(math.Log2(float64(1048576*2)))
	//fmt.Println(len(array))
	fmt.Println(countPairs(array))
	fmt.Println(time.Since(start))
}
