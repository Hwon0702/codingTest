package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	Node []int
}

func (q *Queue) Enqueue(n int) {
	q.Node = append(q.Node, n)
}

func (q *Queue) Dequeue() (ret int) {
	if len(q.Node) <= 0 {
		return -1
	} else {
		ret = q.Node[0]
		if len(q.Node) > 1 {
			q.Node = q.Node[1:]
		} else {
			q.Node = []int{}
		}
	}
	return ret
}

func (q *Queue) Len() int {
	return len(q.Node)
}

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	visited []bool
	graph   [][]int
)

func bfs(start int, group bool) bool {
	var q = new(Queue)
	var now int
	q.Enqueue(start)
	visited[start] = group
	for q.Len() > 0 {
		now = q.Dequeue()
		for _, v := range graph[now] {
			if !visited[v] {
				q.Enqueue(v)
				visited[v] = !visited[now]
			} else if visited[v] == visited[now] {
				return false
			}
		}
	}
	return true
}

func main() {
	defer writer.Flush()
	var K, V, E, u, v int

	fmt.Fscanf(reader, "%d\n", &K)
	for i := 0; i < K; i++ {
		var res bool
		fmt.Fscanf(reader, "%d %d\n", &V, &E)
		graph = make([][]int, V)
		visited = make([]bool, V)
		for j := 0; j < E; j++ {
			fmt.Fscanf(reader, "%d %d\n", &u, &v)
			graph[u-1] = append(graph[u-1], v-1)
			graph[v-1] = append(graph[v-1], u-1)
		}
		for j := 0; j < V; j++ {
			if !visited[j] {
				res = bfs(j, false)
				if !res {
					break
				}
			}
		}
		if res {
			fmt.Fprintf(writer, "YES\n")
		} else {
			fmt.Fprintf(writer, "NO\n")
		}
	}
}
