package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
		}
	}
	for k := 0; k < n; k++ {
		for a := 0; a < n; a++ {
			for b := 0; b < n; b++ {
				if graph[a][k] > 0 && graph[k][b] > 0 {
					graph[a][b] = 1
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for _, v := range graph[i] {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintf(writer, "\n")
	}
}
