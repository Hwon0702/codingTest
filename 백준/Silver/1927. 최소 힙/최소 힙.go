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
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	var intHeap = new(IntHeap)
	heap.Init(intHeap)
	for i := 0; i < tc; i++ {
		var num int
		fmt.Fscanf(reader, "%d\n", &num)
		if num > 0 {
			heap.Push(intHeap, num)
		} else if num == 0 {
			if len(*intHeap) >= 1 {
				fmt.Fprintf(writer, "%d\n", heap.Pop(intHeap).(int))
			} else {
				fmt.Fprintf(writer, "%d\n", 0)
			}

		}
	}

}
