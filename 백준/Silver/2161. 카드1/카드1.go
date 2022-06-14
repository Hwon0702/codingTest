package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	Nodes []int
}

func (q *Queue) Size() int {
	return len(q.Nodes)
}
func (q *Queue) Push(n int) {
	q.Nodes = append(q.Nodes, n)
}
func (q *Queue) Pop() (ret int) {
	if q.Size() <= 0 {
		return -1
	}
	ret = q.Nodes[0]
	if q.Size() == 1 {
		q.Nodes = []int{}
	} else {
		q.Nodes = q.Nodes[1:]
	}
	return ret
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var queue = new(Queue)
	for i := 1; i <= n; i++ {
		queue.Push(i)
	}
	for queue.Size() > 1 {
		fmt.Fprintf(writer, "%d ", queue.Pop())
		queue.Push(queue.Pop())
	}
	fmt.Fprintf(writer, "%d", queue.Pop())

}
