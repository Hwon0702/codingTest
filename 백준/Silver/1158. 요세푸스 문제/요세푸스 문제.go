package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Queue struct {
	Nodes []string
}

func (q *Queue) Push(n string) {
	q.Nodes = append(q.Nodes, n)
}
func (q *Queue) Pop() (n string) {
	if len(q.Nodes) <= 0 {
		return ""
	}
	n = q.Nodes[0]
	if len(q.Nodes) == 1 {
		q.Nodes = []string{}
	} else {
		q.Nodes = q.Nodes[1:]
	}
	return
}

func (q *Queue) Len() int {
	return len(q.Nodes)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, k, cnt int
	var c string
	var res = []string{}
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var people = new(Queue)
	for i := 1; i <= n; i++ {
		people.Push(fmt.Sprint(i))
	}
	for people.Len() > 0 {
		c = people.Pop()
		cnt += 1
		if cnt == k {
			res = append(res, c)
			cnt = 0
			continue
		}
		people.Push(c)
	}
	fmt.Fprintf(writer, "<%s>", strings.Join(res, ", "))
}
