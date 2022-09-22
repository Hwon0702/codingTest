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
		if len(q.Node) <= 0 {
			q.Node = []int{}

		} else {
			q.Node = q.Node[1:]
		}
	}
	return ret
}

func (q *Queue) Size() int {
	return len(q.Node)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m, b, a, v int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var persons = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		persons[i] = []int{}
	}
	var cnt = make([]int, n+1)
	var res = []int{}
	var q = new(Queue)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d\n", &b, &a)
		persons[b] = append(persons[b], a)
		cnt[a]++
	}
	for i := 1; i <= n; i++ {
		if cnt[i] == 0 {
			q.Enqueue(i)
			res = append(res, i)
		}
	}
	for q.Size() > 0 {
		v = q.Dequeue()
		for _, v := range persons[v] {
			cnt[v]--
			if cnt[v] == 0 {
				q.Enqueue(v)
				res = append(res, v)
			}
		}
	}
	for _, v := range res {
		fmt.Fprintf(writer, "%d ", v)
	}
}
