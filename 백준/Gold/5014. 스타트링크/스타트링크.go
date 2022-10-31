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
func (q *Queue) Dequeue() (n int) {
	if len(q.Node) <= 0 {
		return -1
	} else {
		n = q.Node[0]
		if len(q.Node) == 1 {
			q.Node = []int{}
		} else {
			q.Node = q.Node[1:]
		}
	}
	return n
}

func (q *Queue) Size() int {
	return len(q.Node)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var f, s, g, u, d int
	fmt.Fscanf(reader, "%d %d %d %d %d\n", &f, &s, &g, &u, &d)
	var visited = make([]bool, f+1)
	var count = make([]int, f+1)

	var q = new(Queue)
	q.Enqueue(s)
	visited[s] = true
	for q.Size() > 0 {
		now := q.Dequeue()
		if now == g {
			break
		}
		if 0 < now-d && now-d <= f && visited[now-d] == false {
			visited[now-d] = true
			q.Enqueue(now - d)
			count[now-d] = count[now] + 1
		}
		if 0 < now+u && now+u <= f && visited[now+u] == false {
			visited[now+u] = true
			q.Enqueue(now + u)
			count[now+u] = count[now] + 1
		}
	}
	if count[g] > 0 || s == g {
		fmt.Fprintf(writer, "%d\n", count[g])
	} else {

		fmt.Fprintf(writer, "%s\n", "use the stairs")
	}
}
