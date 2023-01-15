package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	graph  = [][]int{}
	n      int
)

type Queue struct {
	Node []int
}

func (q *Queue) Push(n int) {
	q.Node = append(q.Node, n)
}

func (q *Queue) Pop() (r int) {
	if len(q.Node) <= 0 {
		return -1
	} else if len(q.Node) == 1 {
		r = q.Node[0]
		q.Node = []int{}
	} else {
		r = q.Node[0]
		q.Node = q.Node[1:]
	}
	return
}

func (q *Queue) Len() int {
	return len(q.Node)
}

func find(s int) float64 {
	var q = new(Queue)
	var cost = make([]float64, n+1)
	var visited = make([]bool, n+1)
	var now int
	var res float64
	for i := 0; i <= n; i++ {
		cost[i] = 999999
	}
	visited[0] = true
	visited[s] = true
	cost[s] = 0
	q.Push(s)
	for q.Len() > 0 {
		now = q.Pop()
		for _, v := range graph[now] {
			if !visited[v] {
				q.Push(v)
				visited[v] = true
				cost[v] = math.Min(cost[v], cost[now]+1)
			}
		}
	}
	for i := 1; i < n+1; i++ {
		res += cost[i]
	}
	return res

}

func main() {
	defer writer.Flush()
	var m, s, e, ans int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	graph = make([][]int, n+1)
	var res = make([]float64, n+1)
	var v = math.MaxFloat64
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &s, &e)
		graph[s] = append(graph[s], e)
		graph[e] = append(graph[e], s)
	}
	for i := 1; i < n+1; i++ {
		res[i] = find(i)
		if v > res[i] {
			ans = i
			v = res[i]
		}
	}
	fmt.Fprintf(writer, "%d\n", ans)

}
