package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var ans float64
	fmt.Fscanf(reader, "%d\n", &n)
	var cost = make([][]float64, n)
	var res = make([][]float64, n)
	for i := 0; i < n; i++ {
		res[i] = make([]float64, 3)
		cost[i] = make([]float64, 3)
	}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%f %f %f\n", &cost[i][0], &cost[i][1], &cost[i][2])
	}

	ans = 9999999999
	for f := 0; f < 3; f++ {
		for i := 0; i < 3; i++ {
			if i == f {
				res[0][i] = 9999999999
			} else {
				res[0][i] = cost[0][i]
			}
		}
		for i := 1; i < n; i++ {
			res[i][0] = math.Min(res[i-1][1], res[i-1][2]) + cost[i][0]
			res[i][1] = math.Min(res[i-1][0], res[i-1][2]) + cost[i][1]
			res[i][2] = math.Min(res[i-1][1], res[i-1][0]) + cost[i][2]
		}
		for i := 0; i < 3; i++ {
			if i == f {
				ans = math.Min(ans, res[n-1][i])
			}
		}
	}
	fmt.Printf("%0.f\n", ans)
}
