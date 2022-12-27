// B - Проверка даты (10/10 баллов)
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

	var sets, day, month, year int
	fmt.Fscan(in, &sets)

	for i := 0; i < sets; i++ {
		fmt.Fscan(in, &day, &month, &year)

		switch month {
		case 1:
			fallthrough
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 10:
			fallthrough
		case 12:
			if day <= 31 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}

		case 4:
			fallthrough
		case 6:
			fallthrough
		case 9:
			fallthrough
		case 11:
			if day <= 30 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		case 2:
			if day == 29 && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
				fmt.Fprintln(out, "Yes")
			} else if day <= 28 {
				fmt.Fprintln(out, "Yes")
			} else {
				fmt.Fprintln(out, "No")
			}
		}
	}
}
