// G. Звёзды для отелей (40??/25+15 баллов)
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

type hotelData struct {
	number int
	votes  int
	stars  int
}

func main() {
	t1 := time.Now()
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var sets, hotels, temp int
	fmt.Fscan(in, &sets)
	fmt.Fprintln(out, "==========================")

nextHotel:
	for i := 0; i < sets; i++ {
		fmt.Fscan(in, &hotels)
		hotelsV := make([]hotelData, hotels, hotels)
		for j := 0; j < hotels; j++ {
			fmt.Fscan(in, &temp)
			hotelsV[j] = hotelData{j, temp, 0}

		}
		fmt.Fprintln(out, "Scan completed in :", time.Since(t1))
		t1 = time.Now()

		sort.Slice(hotelsV, func(i, j int) (less bool) {
			return hotelsV[i].votes < hotelsV[j].votes
		})
		fmt.Fprintln(out, "Sort completed in :", time.Since(t1))
		t1 = time.Now()
		//fmt.Fprintln(out, "Sorted")
		//for o := 0; o < hotels; o++ {
		//	fmt.Fprint(out, hotelsV[o].votes, " ")
		//}
		//fmt.Fprintln(out, "\nEnd")

		for n := 1; 5*n <= hotels-10; n++ { // n+4 + n+3 + n+2 + n+1 + n  <= hotels
			//fmt.Fprintln(out, "Now I'm here")
			star := 5
			retry := false
			starsCount := make([]int, 7, 7) // sC[6], sC[0] always 0

			for l := 0; l < n; l++ { // awarding 5 stars
				hotelsV[len(hotelsV)-l-1].stars = star
				starsCount[star]++
			}
			//fmt.Fprintln(out, hotelsV)

			for j := len(hotelsV) - n - 1; j >= 0; j-- {
				if star > 1 &&
					hotelsV[j].votes < hotelsV[j+1].votes &&
					starsCount[star] > starsCount[star+1] {
					star--
				}
				starsCount[star]++
				hotelsV[j].stars = star
			}
			n = starsCount[5] // advancing n to eliminate dupe counts !!!

			for j := 1; j < 5; j++ {
				if starsCount[j] < starsCount[j+1] {
					retry = true
					break
				}
			}

			if !retry {
				sort.Slice(hotelsV, func(i, j int) (less bool) {
					return hotelsV[i].number < hotelsV[j].number
				})

				for j := 0; j < hotels-1; j++ {
					fmt.Fprint(out, hotelsV[j].stars, " ")
				}
				fmt.Fprintln(out, hotelsV[len(hotelsV)-1].stars)
				fmt.Fprintln(out, starsCount)
				continue nextHotel
			}
			//fmt.Fprintln(out, starsCount)
		}
		for p := 0; p < hotels-1; p++ { // print for not solvable
			fmt.Fprint(out, "-1 ")
		}
		fmt.Fprintln(out, "-1")
	}
	fmt.Fprintln(out, time.Since(t1))
}
