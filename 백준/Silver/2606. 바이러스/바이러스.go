package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	var visitCnt int
	fmt.Fscanf(reader, "%d\n%d\n", &n, &m)
	var graphs = make([][]int, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		graphs[a] = append(graphs[a], b)
		graphs[b] = append(graphs[b], a)
	}
	var visited = make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		visited[i] = false
	}
	visited[0] = true
	var queue = []int{}
	queue = append(queue, 1)
	visited[1] = true
	for len(queue) != 0 {
		visit := queue[0]
		queue = queue[1:]
		for _, visitNode := range graphs[visit] {
			if visited[visitNode] != true {
				visited[visitNode] = true
				queue = append(queue, visitNode)
			}
		}
	}
	for i := 1; i <= n; i++ {
		if visited[i] {
			visitCnt++
		}
	}
	fmt.Fprintf(writer, "%d", visitCnt-1)
}
