package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	m, n   int
	dx     = []int{-1, 1, 0, 0}
	dy     = []int{0, 0, -1, 1}
	graph  = [][]int{}
	routes = [][]int{}
)

func dfs(x, y int) int {
	if x == m-1 && y == n-1 {
		return 1
	}
	if routes[x][y] != -1 {
		return routes[x][y]
	}
	routes[x][y] = 0
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if 0 <= nx && nx < m && 0 <= ny && ny < n && graph[nx][ny] < graph[x][y] {
			routes[x][y] += dfs(nx, ny)
		}
	}
	return routes[x][y]

}

func main() {
	defer writer.Flush()
	fmt.Fscanf(reader, "%d %d\n", &m, &n)
	graph = make([][]int, m)
	routes = make([][]int, m)

	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
		routes[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
			routes[i][j] = -1
		}
	}
	fmt.Fprintf(writer, "%d\n", dfs(0, 0))
}
