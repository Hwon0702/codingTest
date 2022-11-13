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
	var r, g, b float64
	fmt.Fscanf(reader, "%d\n", &n)
	var selected = make([][]float64, n)
	for i := 0; i < n; i++ {
		selected[i] = make([]float64, 3)
	}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%f %f %f\n", &r, &g, &b)
		if i == 0 {
			selected[0][0], selected[0][1], selected[0][2] = r, g, b
		} else {
			selected[i][0] = math.Min(selected[i-1][1], selected[i-1][2]) + r
			selected[i][1] = math.Min(selected[i-1][0], selected[i-1][2]) + g
			selected[i][2] = math.Min(selected[i-1][0], selected[i-1][1]) + b
		}
	}
	fmt.Println(math.Min(math.Min(selected[n-1][0], selected[n-1][1]), selected[n-1][2]))
}
