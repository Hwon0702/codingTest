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
	var n, k, now int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var visited = make([]int, 100001)
	for i := 0; i < 100001; i++ {
		visited[i] = -1
	}
	var q = new(Queue)
	visited[n] = 0
	var cnt int
	q.Enqueue(n)
	for q.Size() > 0 {
		now = q.Dequeue()
		if now == k {
			cnt += 1
		}
		if 0 <= now*2 && now*2 < 100001 {
			if visited[now*2] == -1 || visited[now*2] == visited[now]+1 {
				visited[now*2] = visited[now] + 1
				q.Enqueue(now * 2)
			}
		}
		if 0 <= now+1 && now+1 < 100001 {
			if visited[now+1] == -1 || visited[now+1] == visited[now]+1 {
				visited[now+1] = visited[now] + 1
				q.Enqueue(now + 1)
			}
		}
		if 0 <= now-1 && now-1 < 100001 {
			if visited[now-1] == -1 || visited[now-1] == visited[now]+1 {
				visited[now-1] = visited[now] + 1
				q.Enqueue(now - 1)
			}
		}
	}
	fmt.Printf("%d\n%d", visited[k], cnt)
}
