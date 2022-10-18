package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type Queue struct {
	Now []int
	Cnt []int
}

func (q *Queue) Enqueue(now, cnt int) {
	q.Now = append(q.Now, now)
	q.Cnt = append(q.Cnt, cnt)
}
func (q *Queue) Dequeue() (int, int) {
	if len(q.Now) <= 0 {
		return -1, -1
	} else {
		n := q.Now[0]
		c := q.Cnt[0]
		if len(q.Now) == 1 {
			q.Now = []int{}
			q.Cnt = []int{}
		} else {
			q.Now = q.Now[1:]
			q.Cnt = q.Cnt[1:]
		}
		return n, c
	}
}
func (q *Queue) Len() int {
	return len(q.Now)
}
func main() {
	defer writer.Flush()
	var n, k, now, cnt, idx int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var queue = new(Queue)
	var visited = make([]bool, 100001)
	var route = make([]int, 100001)
	var routeRes = []int{}
	queue.Enqueue(n, 0)
	for queue.Len() > 0 {
		now, cnt = queue.Dequeue()
		if now == k {
			break
		}
		if 0 <= now*2 && now*2 < 100001 && !visited[now*2] {
			visited[now*2] = true
			route[now*2] = now
			queue.Enqueue(now*2, cnt+1)
		}
		if 0 <= now+1 && now+1 < 100001 && !visited[now+1] {
			visited[now+1] = true
			route[now+1] = now
			queue.Enqueue(now+1, cnt+1)
		}
		if 0 <= now-1 && now-1 < 100001 && !visited[now-1] {
			visited[now-1] = true
			route[now-1] = now
			queue.Enqueue(now-1, cnt+1)
		}
	}
	idx = k
	for i := 0; i < cnt; i++ {
		routeRes = append(routeRes, route[idx])
		idx = route[idx]
	}
	fmt.Fprintf(writer, "%d\n", cnt)
	for i := len(routeRes) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%d ", routeRes[i])
	}
	fmt.Fprintf(writer, "%d ", k)
}
