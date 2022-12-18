package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	cnt     int
	visited []float64
)

func check(x int) bool {
	for i := 0; i < x; i++ {
		if visited[x] == visited[i] || math.Abs(visited[x]-visited[i]) == float64(x-i) {
			return false
		}
	}
	return true
}
func putQueen(x, n int) {
	if x == n {
		cnt += 1
	} else {
		for i := 0; i < n; i++ {
			visited[x] = float64(i)
			if check(x) {
				putQueen(x+1, n)
			}
		}
	}
}
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	cnt = 0
	visited = make([]float64, n)
	putQueen(0, n)
	fmt.Fprintf(writer, "%d\n", cnt)
}
