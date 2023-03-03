// Даны 2 строки версий ПО вида x[.y][.z][-a] (x,y,z числа, а - что угодно), написать функцию, которая возвращает
// 0, если они равны; -1; если первая меньше, 1; если больше; 2, если ошибка.

package main

import (
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`^(\d+)(?:\.(\d+))?(?:\.(\d+))?(?:-(.+))?$`)

//var re = regexp.MustCompile(`^(\d+)(?:\.(\d+))?(?:\.(\d+))?(?:-(\w+))?$`) //last part consists of letters

type ver struct {
	d [3]int
	a string
}

func getVersion(input string) *ver {
	matches := re.FindStringSubmatch(input)
	if len(matches) < 4 {
		return nil
	}

	v := ver{[3]int{}, ""}
	for i := 0; i < 3; i++ {
		temp, _ := strconv.Atoi(matches[i+1])
		v.d[i] = temp
	}
	v.a = matches[4]
	return &v
}

func compareVersions(a, b string) int {
	v1, v2 := getVersion(a), getVersion(b)
	if v1 == nil || v2 == nil {
		return 2
	}

	for i := 0; i < 3; i++ {
		if v1.d[i] > v2.d[i] {
			return 1
		} else if v1.d[i] < v2.d[i] {
			return -1
		}
	}
	if v1.a > v2.a {
		return 1
	} else if v1.a < v2.a {
		return -1
	}
	return 0
}
