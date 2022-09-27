// C. Декодирование строки (10/10 баллов)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	decode := map[string]string{"00": "a", "100": "b", "101": "c", "11": "d"}
	var sets int
	var row string
	fmt.Fscan(in, &sets)

	for i := 0; i < sets; i++ {
		fmt.Fscan(in, &row)
		for j := 0; j < len(row); {
			if val, exist := decode[row[j:j+2]]; exist {
				fmt.Fprint(out, val)
				j += 2
				continue
			}
			if val, exist := decode[row[j:j+3]]; exist {
				fmt.Fprint(out, val)
				j += 3
				continue
			}
		}
		fmt.Fprintln(out)
	}
}
