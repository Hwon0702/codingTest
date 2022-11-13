package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	Idx       []int
	Important []int
}

func (q *Queue) Push(n, important int) {
	q.Idx = append(q.Idx, n)
	q.Important = append(q.Important, important)
}

func (q *Queue) Size() int {
	return len(q.Idx)
}
func (q *Queue) Pop() (idx, important int) {
	if len(q.Idx) <= 0 {
		return -1, -1
	} else {
		idx = q.Idx[0]
		important = q.Important[0]
		if len(q.Idx) == 1 {
			q.Idx = []int{}
			q.Important = []int{}
		} else {
			q.Idx = q.Idx[1:]
			q.Important = q.Important[1:]
		}
	}
	return idx, important
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		var n, docNumber int
		var queue = new(Queue)
		fmt.Fscanf(reader, "%d %d\n", &n, &docNumber)
		queue.Important = make([]int, n)
		queue.Idx = make([]int, n)
		for j := 0; j < n; j++ {
			queue.Idx[j] = j
			fmt.Fscanf(reader, "%d ", &queue.Important[j])
		}
		var printingCnt = 1
		for queue.Size() != 0 {
			idx, important := queue.Pop()
			var printing = true
			for i := 0; i < len(queue.Important); i++ {
				if queue.Important[i] > important {
					queue.Push(idx, important)
					printing = false
					break
				}
			}
			if printing {
				if idx == docNumber {
					fmt.Fprintf(writer, "%d\n", printingCnt)
					break
				}
				printingCnt++

			}
		}
	}

}
