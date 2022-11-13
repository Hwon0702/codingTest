package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	Node   []int
	Dist []int
}

func (q *Queue) Enqueue(n, d int) {
	q.Node = append(q.Node, n)
	q.Dist = append(q.Dist, d)
}
func (q *Queue) Dequeue() (n, d int) {
	if len(q.Node) > 0 {
		n = q.Node[0]
		d = q.Dist[0]

		if len(q.Node) > 1 {
			q.Node = q.Node[1:]
			q.Dist = q.Dist[1:]
		} else {
			q.Node = []int{}
			q.Dist = []int{}
		}
		return n, d
	} else {
		return -1, -1
	}
}
func bfs(s, e, n int, graph []map[int]int) int {
	var queue = new(Queue)
	queue.Enqueue(s, 0)
	var visited = make([]bool, n+1)
	visited[s] = true
	var start, dist int
	var ret int
	for len(queue.Node) > 0 {
		start, dist = queue.Dequeue()
		if start == e {
			ret = dist
			break
		} else {
			for k, v := range graph[start] {
				if visited[k] == false {
					visited[k] = true
					queue.Enqueue(k, dist+v)
				}
			}
		}
	}
	return ret
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	var s, e, d int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var graph = make([]map[int]int, n+1)
	for i := 0; i <= n; i++ {
		graph[i] = make(map[int]int)
	}
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &s, &e, &d)
		graph[s][e] = d
		graph[e][s] = d
	}
	for j := 0; j < m; j++ {
		fmt.Fscanf(reader, "%d %d\n", &s, &e)
		fmt.Println(bfs(s, e, n, graph))
	}
}
