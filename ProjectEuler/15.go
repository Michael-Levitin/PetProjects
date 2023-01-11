//Problem 15- Lattice paths - https://projecteuler.net/problem=15
//Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly
// 6 routes to the bottom right corner. How many such routes are there through a 20×20 grid?

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()
	//fmt.Println(LatticeTravel(1, 2))
	//fmt.Println(LatticeTravel(2, 2))
	//fmt.Println(LatticeTravel(3, 2))
	//fmt.Println(LatticeTravel(3, 3))
	//fmt.Println(LatticeTravel(5, 5))
	//fmt.Println(LatticeTravel(10, 10))
	fmt.Println(LatticeTravel(20, 20)) // WA- 86089583437
	fmt.Println(time.Since(t1))
	t1 = time.Now()
	fmt.Println(LatticeTravelC(20)) // OK -  137846528820
	fmt.Println(time.Since(t1))
	t1 = time.Now()
	fmt.Println(LatticeTravelDPTab(20, 20)) //OK - 137846528820
	fmt.Println(time.Since(t1))
}

func LatticeTravel(n int, m int) int { // TODO Doesn't work, why?
	gridMap := make(map[string]int)
	return helper(n, m, gridMap)
}

func helper(n int, m int, gridMap map[string]int) int {
	if n == 0 || m == 0 {
		return 1
	}

	s := strconv.Itoa(n) + strconv.Itoa(m)
	if x, exists := gridMap[s]; exists {
		return x
	}
	t := helper(n-1, m, gridMap) + helper(n, m-1, gridMap)
	gridMap[s] = t
	return t
}

func LatticeTravelC(n int) int {
	c := 1
	for i := 1; i <= n; i++ {
		c = c * (n + i) / i
		//c *= (n + i) / i	// TODO doesn't work, why
	}
	return c
}

func LatticeTravelDPTab(n int, m int) int {
	grid := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		grid[i] = make([]int, m+1)
	}

	grid[0][0] = 1

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i > 0 {
				grid[i][j] += grid[i-1][j]
			}
			if j > 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[n][m]
}
