package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var n, m int
	var w float64

	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var graph = make([][]float64, n)
	var res = make([][]float64, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]float64, m)
		res[i] = make([]float64, m)
		var s string
		fmt.Fscanf(reader, "%s\n", &s)
		sArr := strings.Split(s, "")
		for j := 0; j < len(sArr); j++ {
			a, _ := strconv.Atoi(sArr[j])
			graph[i][j] = float64(a)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				res[i][j] = graph[i][j]
			} else if graph[i][j] == 0 {
				res[i][j] = 0
			} else {
				res[i][j] = math.Min(math.Min(res[i-1][j-1], res[i-1][j]), res[i][j-1]) + 1
			}
			w = math.Max(res[i][j], w)
		}
	}
	fmt.Printf("%0.f\n", w*w)
}
