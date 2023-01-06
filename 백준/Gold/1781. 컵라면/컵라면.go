package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	r := h[i] < h[j]
	return r
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, res int
	h := new(MinHeap)
	fmt.Fscanf(reader, "%d\n", &n)
	var lists = make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &lists[i][0], &lists[i][1])
	}
	sort.SliceStable(lists, func(i, j int) bool {
		return lists[i][0] < lists[j][0]
	})
	for _, v := range lists {
		heap.Push(h, v[1])
		if v[0] < h.Len() {
			heap.Pop(h)
		}
	}
	for _, v := range *h {
		res += v
	}
	fmt.Fprintf(writer, "%d\n", res)
}
