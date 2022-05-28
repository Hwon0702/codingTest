package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(elem interface{}) {
	*h = append(*h, elem.(int))
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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, sum int
	fmt.Fscanf(reader, "%d\n", &n)
	var cardsHeap = new(IntHeap)
	heap.Init(cardsHeap)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscanf(reader, "%d\n", &a)
		heap.Push(cardsHeap, a)
	}
	for cardsHeap.Len() > 0 {
		a := heap.Pop(cardsHeap).(int)
		if cardsHeap.Len() > 0 {
			b := heap.Pop(cardsHeap).(int)
			sum += a + b
			heap.Push(cardsHeap, a+b)
		}
	}
	fmt.Fprintf(writer, "%d", sum)
}
