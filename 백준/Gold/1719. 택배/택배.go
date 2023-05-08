package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	n, m   int
)

func floydWarshal(graph [][]int) [][]int {
	var res = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		res[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			if i != j {
				res[i][j] = j
			}
		}
	}
	for k := 1; k < n+1; k++ {
		res[k][k] = 0
		for i := 1; i < n+1; i++ {
			for j := 1; j < n+1; j++ {
				if graph[i][j] > graph[i][k]+graph[k][j] {
					graph[i][j] = graph[i][k] + graph[k][j]
					res[i][j] = res[i][k]
				}
			}
		}
	}
	return res
}
func main() {
	defer writer.Flush()
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var graph = make([][]int, n+1)
	var a, b, c int
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			if i != j {
				graph[i][j] = 99999999
			} else {
				graph[i][j] = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		graph[a][b] = c
		graph[b][a] = c
	}
	res := floydWarshal(graph)
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if i == j {
				fmt.Fprintf(writer, "%s ", "-")
			} else {
				fmt.Fprintf(writer, "%d ", res[i][j])
			}
		}
		fmt.Fprint(writer, "\n")
	}
}
