package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type Jewel struct {
	weight uint64
	cost   uint64
}

type IntHeap []uint64

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {

	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(elem interface{}) {
	*h = append(*h, elem.(uint64))
}

func (h *IntHeap) Pop() interface{} {

	old := *h
	n := len(old)

	elem := old[n-1]
	*h = old[0 : n-1]

	return elem
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	var cost, weight, res uint64
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var jewels = make([]Jewel, n)
	var bags = make([]uint64, k)
	var maxHeap = new(IntHeap)
	heap.Init(maxHeap)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &weight, &cost)
		jewels[i] = Jewel{weight, cost}
	}
	sort.Slice(jewels, func(i, j int) bool { return jewels[i].weight < jewels[j].weight })
	for i := 0; i < k; i++ {
		fmt.Fscanf(reader, "%d\n", &bags[i])
	}
	sort.Slice(bags, func(i, j int) bool { return bags[i] < bags[j] })
	i := 0
	for _, bag := range bags {
		for ; i < n && jewels[i].weight <= bag; i++ {
			heap.Push(maxHeap, jewels[i].cost)
		}
		if maxHeap.Len() > 0 {
			res += heap.Pop(maxHeap).(uint64)
		}
	}
	fmt.Println(res)
}
