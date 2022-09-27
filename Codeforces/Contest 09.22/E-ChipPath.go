// E - Путь фишки (20/20 баллов)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var sets, rows, cols int
	var rowString string
	fmt.Fscan(in, &sets)
	fmt.Fprintln(out, "==========================")
	for i := 0; i < sets; i++ {
		fmt.Fscan(in, &rows, &cols)

		fieldMap := make([][]string, rows)
		for j := 0; j < rows; j++ {
			fieldMap[j] = make([]string, cols)
			fmt.Fscan(in, &rowString)
			fieldMap[j] = strings.Split(rowString, "")
		}

		startEnd := make([][]int, 0, 2)

		for j := 0; j < rows; j++ { // looking where to start
			for k := 0; k < cols; k++ {
				count := 0
				if fieldMap[j][k] == "*" {
					if j-1 < 0 || fieldMap[j-1][k] == "." {
						count++
					}
					if k-1 < 0 || fieldMap[j][k-1] == "." {
						count++
					}
					if j+1 >= rows || fieldMap[j+1][k] == "." {
						count++
					}
					if k+1 >= cols || fieldMap[j][k+1] == "." {
						count++
					}
				}
				if count == 3 {
					//x, y = j, k
					//fmt.Fprintln(out, "Start at", x, y)
					startEnd = append(startEnd, []int{j, k})
				}
			}

		}
		previousX, previousY := -1, -1
		for j, k := startEnd[0][0], startEnd[0][1]; j != startEnd[1][0] || k != startEnd[1][1]; {
			fmt.Println(j, k)
			if j-1 >= 0 && fieldMap[j-1][k] == "*" {
				if j-1 != previousX || k != previousY {
					fmt.Fprint(out, "U")
					previousX, previousY = j, k
					j--
					continue
				}
			}
			if j+1 < rows && fieldMap[j+1][k] == "*" {
				if j+1 != previousX || k != previousY {
					fmt.Fprint(out, "D")
					previousX, previousY = j, k
					j++
					continue
				}
			}
			if k-1 >= 0 && fieldMap[j][k-1] == "*" {
				if j != previousX || k-1 != previousY {
					fmt.Fprint(out, "L")
					previousX, previousY = j, k
					k--
					continue
				}
			}
			if k+1 < cols && fieldMap[j][k+1] == "*" {
				if j != previousX || k+1 != previousY {
					fmt.Fprint(out, "R")
					previousX, previousY = j, k
					k++
					continue
				}
			}
		}
		fmt.Fprintln(out)
		//fmt.Fprintln(out, startEnd)
		//fmt.Fprintln(out, fieldMap)
	}
}
