//H. Анализ игрового поля (30??/30 баллов)
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	t1 := time.Now()
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var sets, n, m int
	var row string
	fmt.Fscan(in, &sets)
	fmt.Fprintln(out, "==========================")

	for i := 0; i < sets; i++ {
		fmt.Fscan(in, &n, &m)
		field := make([][]string, n, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &row)
			field[j] = strings.Split(row, "")
		}
		nesting := []int{}
		findAreas(field, 0, 0, n, m, &nesting, 0)
		sort.Ints(nesting)
		fmt.Fprintln(out, nesting)
		//fmt.Fprintln(out, "--------------------------")
	}
	fmt.Fprintln(out, time.Since(t1))
}

func findAreas(field [][]string, xS, yS, xE, yE int, nesting *[]int, nestLevel int) {
	//fmt.Println(xS, yS, xE, yE, "NestLevel Start", nestLevel)
	for i := xS; i < xE-1; i++ { // looking for start corners
		for j := yS; j < yE-1; j++ {
			if field[i][j] == "*" &&
				field[i][j+1] == "*" &&
				field[i+1][j] == "*" {
				k, l := findCorner(field, i, j, xE, yE)
				//fmt.Println(i, j, k, l, "NestLevel Found", nestLevel)
				*nesting = append(*nesting, nestLevel)
				deleteArea(field, i, j, k, l)
				//print2D(field)
				if k-i > 5 && l-j > 5 {
					//fmt.Println("Next Area", i+2, j+2, k-2, l-2, nestLevel+1)
					findAreas(field, i+2, j+2, k-2, l-2, nesting, nestLevel+1)
				}
			}
		}
	}
}

func findCorner(field [][]string, xS, yS, xE, yE int) (int, int) {
	j := yS + 2         // min 3x3
	for ; j < yE; j++ { // top right corner
		if field[xS+1][j] == "*" {
			break
		}
	}

	i := xS + 2
	for ; i < xE; i++ { // looking for start corners
		if field[i][j-1] == "*" {
			break
		}
	}
	return i, j
}

func deleteArea(field [][]string, xS, yS, xE, yE int) {
	for i := xS; i <= xE; i++ {
		field[i][yS] = "."
		field[i][yE] = "."
	}
	for j := yS; j <= yE; j++ {
		field[xS][j] = "."
		field[xE][j] = "."
	}
}

func print2D(slice2D [][]string) {
	fmt.Print("  ")
	for j := 0; j < len(slice2D[0]); j++ {
		if j < 10 {
			fmt.Print("_", j)
		} else {
			fmt.Print(j)
		}

	}
	fmt.Print("\n")
	for i := 0; i < len(slice2D); i++ {
		if i < 10 {
			fmt.Print("_", i, " ")
		} else {
			fmt.Print(i, " ")
		}
		for j := 0; j < len(slice2D[i]); j++ {
			//fmt.Print(slice2D[i][j])
			fmt.Print(slice2D[i][j], " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
