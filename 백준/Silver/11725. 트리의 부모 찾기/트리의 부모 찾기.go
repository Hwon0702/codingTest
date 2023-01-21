package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	graph  = [][]int{}
	tree   = []int{}
)

func dfs(start int) {
	for _, v := range graph[start] {
		if tree[v] == 0 {
			tree[v] = start
			dfs(v)
		}
	}
}
func main() {
	defer writer.Flush()
	var n, a, b int
	fmt.Fscanf(reader, "%d\n", &n)
	graph = make([][]int, n+1)
	tree = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	dfs(1)
	for i := 2; i < n+1; i++ {
		fmt.Fprintf(writer, "%d\n", tree[i])
	}
}
