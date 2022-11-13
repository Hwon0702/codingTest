package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var DfsVisited []bool
var BfsVisited []bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m, s int
	fmt.Fscanf(reader, "%d %d %d\n", &n, &m, &s)
	var graphs = make([][]int, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		graphs[a] = append(graphs[a], b)
		graphs[b] = append(graphs[b], a)
	}
	for i, _ := range graphs {
		sort.Ints(graphs[i])
	}
	Dfs(s, graphs, writer)
	Bfs(s, graphs, writer)

}

func Dfs(start int, graph [][]int, writer *bufio.Writer) {
	if len(DfsVisited) == 0 {
		DfsVisited = make([]bool, len(graph), len(graph))
	}
	DfsVisited[start] = true
	fmt.Fprintf(writer, "%d ", start)
	for _, v := range graph[start] {
		if DfsVisited[v] != true {
			Dfs(v, graph, writer)
		}
	}
}

func Bfs(start int, graph [][]int, writer *bufio.Writer) {

	var queue = []int{}

	queue = append(queue, start)
	if len(BfsVisited) <= 0 {
		BfsVisited = make([]bool, len(graph), len(graph))
		BfsVisited[start] = true
		fmt.Fprintf(writer, "\n%d ", start)
	}
	for len(queue) != 0 {
		visit := queue[0]
		queue = queue[1:]
		for _, v := range graph[visit] {
			if !BfsVisited[v] {
				BfsVisited[v] = true
				fmt.Fprintf(writer, "%d ", v)
				queue = append(queue, v)
			}
		}
	}
}
