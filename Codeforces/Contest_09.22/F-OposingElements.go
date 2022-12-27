// E. Противоположные элементы (20/20 баллов) - create map first
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	t1 := time.Now()
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var sets, e, a, b, tripets int
	fmt.Fscan(in, &sets)
	fmt.Fprintln(out, "==========================")
	for i := 0; i < sets; i++ {
		fmt.Fscan(in, &tripets)
		triMap := make(map[int][]int)
		for j := 0; j < tripets; j++ {
			fmt.Fscan(in, &e, &a, &b)
			//fmt.Fprintln(out, e, a, b)
			triMap[e] = make([]int, 2, 2)
			triMap[e][0], triMap[e][1] = a, b
			//fmt.Fprintln(out, e, triMap[e])
		}
		//fmt.Fprintln(out, triMap)
		triSlice := make([]int, 3, tripets*2)
		triSlice[0], triSlice[1], triSlice[2] = a, e, b // using last input
		delete(triMap, e)
		//fmt.Fprintln(out, triSlice)
		//fmt.Fprintln(out, "before", triMap)

		for len(triMap) > 0 {
			temp := triMap[triSlice[len(triSlice)-1]]
			if triSlice[len(triSlice)-2] == temp[0] {
				triSlice = append(triSlice, temp[1])
			} else {
				triSlice = append(triSlice, temp[0])
			}
			delete(triMap, triSlice[len(triSlice)-2])
			//fmt.Fprintln(out, "after", triMap)
		}
		triSlice = triSlice[2:]
		//fmt.Fprintln(out, triSlice, cap(triSlice))
		for j := 0; j < len(triSlice)/2; j++ {
			fmt.Fprintln(out, triSlice[j], triSlice[len(triSlice)/2+j])
		}
		fmt.Fprintln(out)
	}
	fmt.Fprintln(out, time.Since(t1))
}
