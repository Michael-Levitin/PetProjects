// I. Дерево комментариев (30/30 баллов) works, not checked,
//TODO spare nl at the end,
//TODO tabsF are static - 100 levels deep
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func main() {
	t1 := time.Now()
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var sets, n, p, c int
	var s string
	fmt.Fscan(in, &sets)
	fmt.Fprintln(out, "==========================")

	for i := 0; i < sets; i++ {
		fmt.Fscan(in, &c)
		mapChat := make(map[int]map[int]string)
		for j := 0; j < c; j++ {
			fmt.Fscan(in, &n, &p)
			s, _ = in.ReadString('\n') // scanning string with spaces
			//s = strings.TrimSuffix(s, "\n") // remove newline
			if _, exists := mapChat[p]; !exists {
				mapChat[p] = make(map[int]string)
			}
			mapChat[p][n] = s[1:] // also deleting first space
		}

		tabsF := make([]bool, 100) // slice of tabs
		printMap(mapChat, -1, out, 0, &tabsF)
	}
	fmt.Fprintln(out, time.Since(t1))
}

func printMap(mapChat map[int]map[int]string, key int, out io.Writer, tab int, tabsF *[]bool) {
	//fmt.Fprintln(out, tabsF)
	mapKeys := make([]int, 0, len(mapChat[key]))
	for k, _ := range mapChat[key] {
		mapKeys = append(mapKeys, k)
	}
	sort.Ints(mapKeys) //

	for i, k := range mapKeys {
		printDivs(tab, tabsF, out)
		if tab != 0 {
			fmt.Fprintln(out, "|")
		}
		printDivs(tab, tabsF, out)
		if tab != 0 {
			fmt.Fprint(out, "|--")
		}

		fmt.Fprint(out, mapChat[key][k])
		if i < len(mapKeys)-1 {
			(*tabsF)[tab] = true
		} else {
			(*tabsF)[tab] = false
		}
		printMap(mapChat, k, out, tab+1, tabsF)
		if tab == 0 {
			fmt.Fprintln(out)
		}
	}
}

func printDivs(tab int, tabsF *[]bool, out io.Writer) {
	for j := 1; j < tab; j++ {
		if (*tabsF)[j] {
			fmt.Fprint(out, "|  ")
		} else {
			fmt.Fprint(out, "   ")
		}
	}
}
