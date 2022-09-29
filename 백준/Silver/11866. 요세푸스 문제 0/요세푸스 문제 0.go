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
func (q *Queue) Dequeue() (res int) {
	if len(q.Node) <= 0 {
		return -1
	}
	res = q.Node[0]
	if len(q.Node) > 1 {
		q.Node = q.Node[1:]
	} else {
		q.Node = []int{}
	}
	return
}
func (q *Queue) Len() int {
	return len(q.Node)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, k, p, cnt int
	var res = []int{}
	var q = new(Queue)
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	for i := 1; i < n+1; i++ {
		q.Enqueue(i)
	}
	for q.Len() > 0 {
		p = q.Dequeue()
		cnt++
		if cnt == k {
			res = append(res, p)
			cnt = 0
		} else {
			q.Enqueue(p)
		}
	}
	fmt.Fprintf(writer, "<")
	for i := 0; i < n-1; i++ {
		fmt.Fprintf(writer, "%d, ", res[i])
	}
	fmt.Fprintf(writer, "%d>", res[n-1])
}
