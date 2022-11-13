package main

import (
	"bufio"
	"fmt"
	"os"
)

var visited []bool
var count int

func dfs(graph [][]int, start int) {
	visited[start] = true // dfs에 들어오면 방문했다고 판단
	for i := 0; i < len(graph[start]); i++ {
		if graph[start][i] == 1 && i < len(visited) && !visited[i] { // 연결되어 있고 방문하지 않은 경우
			dfs(graph, i) // 재귀함수 호출
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	var a, b int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var graphs = make([][]int, n+1)
	visited = make([]bool, n+1)
	for i := 0; i <= n; i++ {
		graphs[i] = make([]int, n+1)
	}
	visited[0] = true
	for i := 1; i <= m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		graphs[a][b] = 1
		graphs[b][a] = 1
	}
	for i := 1; i <= n; i++ {
		if !visited[i] {
			count++
			dfs(graphs, i)
		}
	}
	fmt.Fprintf(writer, "%d\n", count)
}
