package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	q := Queue{list.New()}
	return &q
}
func (q *Queue) Push(n []int) {
	q.v.PushBack(n)
}

func (q *Queue) Pop() []int {
	f := q.v.Front()
	if f == nil {
		return nil
	}
	return q.v.Remove(f).([]int)
}

func (q *Queue) Empty() bool {
	if q.v.Len() == 0 {
		return true
	}
	return false
}

var (
	graph [][]int
	dx    = []int{1, -1, 0, 0}
	dy    = []int{0, 0, 1, -1}
	n, m  int
)

func bfs(x, y int) int {

	graph[x][y] = 0
	var width = 1

	q := NewQueue()
	q.Push([]int{x, y})
	for !q.Empty() {
		num := q.Pop()
		for i := 0; i < 4; i++ {
			nx := dx[i] + num[0]
			ny := dy[i] + num[1]
			if 0 <= nx && nx < n && 0 <= ny && ny < m && graph[nx][ny] == 1 {
				q.Push([]int{nx, ny})
				graph[nx][ny] = 0
				width += 1
			}
		}
	}
	return width
}

func main() {
	defer writer.Flush()
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
		}
	}
	cnt := 0
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == 1 {
				cnt += 1
				res = max(bfs(i, j), res)
			}
		}
	}
	fmt.Fprintf(writer, "%d\n%d", cnt, res)
}
