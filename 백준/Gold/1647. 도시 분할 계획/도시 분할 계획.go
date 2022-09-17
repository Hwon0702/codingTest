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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m, s, e, current int
	var c, cost, maxCost, minCost float64
	var hq = new(Nodes)
	var t = new(Node)
	heap.Init(hq)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var lines = make([]map[int]float64, n)
	var visited = make([]bool, n)
	var visitedCnt = n
	for i := 0; i < n; i++ {
		lines[i] = make(map[int]float64)
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %f\n", &s, &e, &c)
		if _, exists := lines[s-1][e-1]; exists {
			lines[s-1][e-1] = math.Min(lines[s-1][e-1], c)
		} else {
			lines[s-1][e-1] = c
		}
		if _, exists := lines[e-1][s-1]; exists {
			lines[e-1][s-1] = math.Min(lines[e-1][s-1], c)
		} else {
			lines[e-1][s-1] = c
		}
	}
	t.Idx = 0
	t.Val = 0
	heap.Push(hq, *t)
	for visitedCnt > 0 && hq.Len() > 0 {
		currentVal := heap.Pop(hq).(Node)
		current = currentVal.Idx
		cost = currentVal.Val
		if visited[current] {
			continue
		} else {
			visited[current] = true
			visitedCnt--
			maxCost = math.Max(maxCost, cost)
			minCost += cost
			for idx, v := range lines[current] {
				if !visited[idx] && idx != current {
					t.Idx = idx
					t.Val = v
					heap.Push(hq, *t)
				}
			}
		}
	}
	fmt.Fprintf(writer, "%0.f", minCost-maxCost)
}
