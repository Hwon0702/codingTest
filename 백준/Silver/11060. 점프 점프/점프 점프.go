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
	Current []int
	Move    []int
	Cnt     []int
}

func (q *Queue) Enqueue(x, m, c int) {
	q.Current = append(q.Current, x)
	q.Move = append(q.Move, m)
	q.Cnt = append(q.Cnt, c)
}
func (q *Queue) Dequeue() (x, m, c int) {
	if len(q.Current) <= 0 {
		return -1, -1, -1
	}
	x = q.Current[0]
	m = q.Move[0]
	c = q.Cnt[0]
	if len(q.Current) == 1 {
		q.Current = []int{}
		q.Move = []int{}
		q.Cnt = []int{}
	} else {
		q.Current = q.Current[1:]
		q.Move = q.Move[1:]
		q.Cnt = q.Cnt[1:]
	}
	return
}
func (q *Queue) Len() int {
	return len(q.Current)
}

func main() {
	defer writer.Flush()
	var n, res, x, move, cnt int
	fmt.Fscanf(reader, "%d\n", &n)
	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	var visited = make([]bool, n)
	var q = new(Queue)
	res = -1
	q.Enqueue(0, numbers[0], 0)
	for q.Len() > 0 {
		x, move, cnt = q.Dequeue()
		if x == n-1 {
			res = cnt
			break
		}
		for i := 1; i <= move; i++ {
			if 0 <= x+i && x+i < n && visited[x+i] == false {
				visited[x+i] = true
				q.Enqueue(x+i, numbers[x+i], cnt+1)
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
