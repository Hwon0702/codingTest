package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, m, s, e int
	var c, ans float64
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var roads = make([][]float64, n+1)
	for i := 0; i < n+1; i++ {
		roads[i] = make([]float64, n+1)
		for j := 0; j < n+1; j++ {
			roads[i][j] = math.MaxFloat64
		}
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %f\n", &s, &e, &c)
		roads[s][e] = c
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			for k := 1; k < n+1; k++ {
				roads[j][k] = math.Min(roads[j][k], roads[j][i]+roads[i][k])
			}
		}
	}
	ans = math.MaxFloat64
	for i := 1; i < n+1; i++ {
		ans = math.Min(ans, roads[i][i])
	}
	if ans >= math.MaxFloat64 {
		ans = -1
	}
	fmt.Fprintf(writer, "%0.f", ans)
}
