package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func floydWarshal(bus [][]float64, n int) [][]float64 {
	var cost = bus
	for k := 1; k <= n; k++ {
		cost[k][k] = 0
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				cost[i][j] = math.Min(bus[i][j], bus[i][k]+bus[k][j])
			}
		}
	}
	return cost
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	var s, e int
	var c float64
	fmt.Fscanf(reader, "%d\n%d\n", &n, &m)
	var bus = make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		bus[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			bus[i][j] = math.MaxFloat64
		}
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %f\n", &s, &e, &c)
		bus[s][e] = math.Min(bus[s][e], c)
	}
	res := floydWarshal(bus, n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if res[i][j] == math.MaxFloat64 {
				res[i][j] = 0
			}
			fmt.Fprintf(writer, "%0.f ", res[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}
}
