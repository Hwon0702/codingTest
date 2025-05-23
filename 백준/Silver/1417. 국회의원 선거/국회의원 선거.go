package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type PriorityQueue struct {
	Nodes []int
}

func NewQueue() *PriorityQueue {
	return new(PriorityQueue)
}
func (p *PriorityQueue) Push(n int) {
	p.Nodes = append(p.Nodes, n)
	sort.Slice(p.Nodes, func(i, j int) bool {
		return p.Nodes[i] > p.Nodes[j]
	})
}

func (p *PriorityQueue) Pop() (ret int) {
	ret = -1
	if p.Empty() {
		return
	}
	ret = p.Front()
	if p.Length() > 1 {
		p.Nodes = p.Nodes[1:]
	} else {
		p.Nodes = []int{}
	}
	return
}

func (p *PriorityQueue) Empty() bool {
	return len(p.Nodes) < 0
}

func (p *PriorityQueue) Front() int {
	return p.Nodes[0]
}
func (p *PriorityQueue) Last() int {
	return p.Nodes[p.Length()-1]
}

func (p *PriorityQueue) Length() int {
	return len(p.Nodes)
}
func main() {
	defer writer.Flush()
	var n, t, target, cnt int
	cnt = 0
	fmt.Fscanf(reader, "%d\n", &n)
	//다솜이는 기호 1번
	fmt.Fscanf(reader, "%d\n", &target)
	if n > 1 {
		q := NewQueue()
		for i := 1; i < n; i++ {
			fmt.Fscanf(reader, "%d\n", &t)
			q.Push(t)
		}
		first := 0
		for true {
			if q.Empty() {
				break
			}
			first = q.Pop()
			if target > first {
				break
			}
			target += 1
			cnt += 1
			q.Push(first - 1)
		}
	}
	fmt.Fprintf(writer, "%d\n", cnt)
}
