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
	var n, m, s, e, res, check int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &s, &e)
		graph[s-1][e-1] = 1
	}
	for path := 0; path < n; path++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if graph[i][path] == 1 && graph[path][j] == 1 {
					graph[i][j] = 1
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		check = 0
		for j := 0; j < n; j++ {
			check += (graph[i][j] + graph[j][i])
		}
		if check == n-1 {
			res++
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
