// D. Сломанный сервер  (15/15 баллов)
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

	var sets, requests int
	fmt.Fscan(in, &sets)
	//fmt.Fprintln(out, "==========================")
	for i := 0; i < sets; i++ {
		fmt.Fscan(in, &requests)
		reqSlice := make([]int, requests, requests)
		for j := 0; j < requests; j++ {
			fmt.Fscan(in, &reqSlice[j])
		}
		var max int
		for j := 0; j < requests; j++ {
			fruMap := make(map[int]struct{})
			for k := j; k < requests; k++ {
				fruMap[reqSlice[k]] = struct{}{}
				if len(fruMap) > 2 {
					//fmt.Fprintln(out, max, j, k-1)
					break
				}
				if k-j+1 > max {
					max = k - j + 1
				}
			}
		}
		fmt.Fprintln(out, max)
	}
	fmt.Fprintln(out, time.Since(t1))
}
