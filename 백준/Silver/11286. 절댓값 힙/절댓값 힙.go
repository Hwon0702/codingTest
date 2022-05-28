package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type FloatHeap []float64

func (h FloatHeap) Len() int {
	return len(h)
}

func (h FloatHeap) Less(i, j int) bool {
	if math.Abs(h[i]) < math.Abs(h[j]) {
		return math.Abs(h[i]) < math.Abs(h[j])
	} else if math.Abs(h[i]) == math.Abs(h[j]) {
		return h[i] < h[j]
	}
	return math.Abs(h[i]) < math.Abs(h[j])

}

func (h FloatHeap) Swap(i, j int) {

	h[i], h[j] = h[j], h[i]
}

func (h *FloatHeap) Push(elem interface{}) {
	*h = append(*h, elem.(float64))
}

func (h *FloatHeap) Pop() interface{} {

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
	var floatHeap = new(FloatHeap)
	heap.Init(floatHeap)
	for i := 0; i < tc; i++ {
		var num float64
		fmt.Fscanf(reader, "%f\n", &num)

		if num != 0 {
			heap.Push(floatHeap, num)
		} else if num == 0 {
			if len(*floatHeap) >= 1 {
				fmt.Fprintf(writer, "%d\n", int(heap.Pop(floatHeap).(float64)))
			} else {
				fmt.Fprintf(writer, "%d\n", 0)
			}

		}
	}

}
