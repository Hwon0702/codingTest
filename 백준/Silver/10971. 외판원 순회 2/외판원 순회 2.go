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
	graphs  = [][]int{}
	visited = []bool{}
	res     = math.MaxInt64
	n       = 0
)

func dfs(start, now, cost, cnt int) {
	if cnt == n && start == now {
		if res > cost {
			res = cost
		}
		return
	}

	if cost > res {
		return
	}

	for i := 0; i < n; i++ {
		if visited[i] || graphs[now][i] == 0 {
			continue
		}

		visited[i] = true
		dfs(start, i, cost+graphs[now][i], cnt+1)
		visited[i] = false
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscanf(reader, "%d\n", &n)
	graphs = make([][]int, n)
	visited = make([]bool, n)
	for i := 0; i < n; i++ {
		graphs[i] = make([]int, n)
		visited[i] = false
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &graphs[i][j])

		}
	}
	for i := 0; i < n; i++ {
		dfs(i, i, 0, 0)
	}
	fmt.Fprintf(writer, "%d\n", res)
}
