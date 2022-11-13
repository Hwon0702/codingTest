package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	Now  []int
	Time []int
}

func (q *Queue) Enqueue(n, t int) {
	q.Now = append(q.Now, n)
	q.Time = append(q.Time, t)
}
func (q *Queue) Dequeue() (n, t int) {
	if len(q.Now) <= 0 {
		n, t = -1, -1
	} else {
		n, t = q.Now[0], q.Time[0]
		if len(q.Now) == 1 {
			q.Now = []int{}
			q.Time = []int{}
		} else {
			q.Now = q.Now[1:]
			q.Time = q.Time[1:]
		}
	}
	return n, t
}

func (q *Queue) Len() int {
	return len(q.Now)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var x, y int
	fmt.Fscanf(reader, "%d %d\n", &x, &y)
	var visited = make([]bool, 100001)
	var q = new(Queue)
	q.Enqueue(x, 0)
	for q.Len() > 0 {
		n, t := q.Dequeue()
		if n == y {
			fmt.Println(t)
			break
		}
		if 0 <= n*2 && n*2 < 100001 && !visited[n*2] {
			q.Enqueue(n*2, t)
			visited[n*2] = true
		}
		if 0 <= n-1 && n-1 < 100001 && !visited[n-1] {
			q.Enqueue(n-1, t+1)
			visited[n-1] = true
		}
		if 0 <= n+1 && n+1 < 100001 && !visited[n+1] {
			q.Enqueue(n+1, t+1)
			visited[n+1] = true
		}
	}

}
