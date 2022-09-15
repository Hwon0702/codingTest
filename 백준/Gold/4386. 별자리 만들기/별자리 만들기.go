package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Node struct {
	Idx int
	Val float64
}

type Nodes []Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Nodes) Less(i, j int) bool {
	return n[i].Val < n[j].Val
}

func (n *Nodes) Push(num interface{}) {
	myNum := num.(Node)
	*n = append(*n, myNum)
}

func (n *Nodes) update(item *Node, idx int, val float64) {
	item.Idx = idx
	item.Val = val
	heap.Fix(n, int(item.Val))
}

func (n *Nodes) Pop() interface{} {
	tmp := *n
	l := len(tmp)
	var res interface{} = tmp[l-1]
	*n = tmp[:l-1]
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, visitCnt int
	var sum float64
	fmt.Fscanf(reader, "%d\n", &n)
	var stars = make([][]float64, n)
	var dir = make([][]float64, n)
	var pq = new(Nodes)
	var temp = new(Node)
	var visit = make([]bool, n)
	heap.Init(pq)
	for i := 0; i < n; i++ {
		stars[i] = make([]float64, 2)
		dir[i] = make([]float64, n)
		fmt.Fscanf(reader, "%f %f\n", &stars[i][0], &stars[i][1])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			} else {
				dir[i][j] = math.Sqrt((math.Pow(stars[i][0]-stars[j][0], 2) + (math.Pow(stars[i][1]-stars[j][1], 2))))
			}
		}
	}
	visitCnt = n
	temp.Idx = 0
	temp.Val = 0
	heap.Push(pq, *temp)
	for visitCnt > 0 && pq.Len() > 0 {
		current := heap.Pop(pq).(Node)
		if visit[current.Idx] {
			continue
		} else {
			visit[current.Idx] = true
			visitCnt--
			sum += current.Val
			for next := 0; next < n; next++ {
				if current.Idx != next && !visit[next] {
					temp.Idx = next
					temp.Val = dir[current.Idx][next]
					heap.Push(pq, *temp)
				}
			}
		}

	}
	fmt.Printf("%.2f", sum)
}
