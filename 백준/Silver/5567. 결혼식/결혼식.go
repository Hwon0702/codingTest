package main

import (
	"bufio"
	"fmt"
	"os"
)

func Bfs(visited []bool, graphs [][]int) int {
	var depth = make(map[int]int)
	var queue = []int{}
	var cnt int
	queue = append(queue, 1)
	for len(queue) != 0 {
		visit := queue[0]
		queue = queue[1:]
		for _, v := range graphs[visit] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
				if visit == 1 {
					depth[v] = 1
				} else if _, exists := depth[visit]; exists {
					depth[v] = depth[visit] + 1
				}
			}
		}
	}
	for k, v := range depth {
		if k != 1 && v <= 2 {
			cnt++
		}
	}
	return cnt
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	var a, b int
	fmt.Fscanf(reader, "%d\n%d\n", &n, &m)
	var graphs = make([][]int, n+1)

	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		graphs[a] = append(graphs[a], b)
		graphs[b] = append(graphs[b], a)
	}
	var visited = make([]bool, n+1)
	visited[0] = true
	res := Bfs(visited, graphs)
	fmt.Fprintf(writer, "%d\n", res)

}
