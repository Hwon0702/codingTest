package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	Weight []int
	Remain []int
}

func (q *Queue) Push(n, i int) {
	q.Weight = append(q.Weight, n)
	q.Remain = append(q.Remain, i)
}

func (q *Queue) Pop() (wei, rem int) {
	if len(q.Weight) <= 0 {
		return -1, -1
	} else if len(q.Weight) == 1 {
		wei = q.Weight[0]
		rem = q.Remain[0]

		q.Weight = []int{}
		q.Remain = []int{}
	} else {
		wei = q.Weight[0]
		rem = q.Remain[0]
		q.Weight = q.Weight[1:]
		q.Remain = q.Remain[1:]
	}
	return wei, rem
}

func (q *Queue) First() (wei, rem int) {
	if len(q.Weight) <= 0 {
		return -1, -1
	} else {
		wei = q.Weight[0]
		rem = q.Remain[0]
	}
	return wei, rem
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, w, l int

	var all = new(Queue)
	var moving = new(Queue)
	var count = 1
	var car, weight, sum int
	fmt.Fscanf(reader, "%d %d %d\n", &n, &w, &l)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &car)
		all.Push(car, w)
	}
	moving.Push(all.Pop())
	sum += moving.Weight[0]
	for len(moving.Remain) > 0 {
		count += 1

		for i, _ := range moving.Remain {
			moving.Remain[i] = moving.Remain[i] - 1
		}
		_, rem := moving.First()
		if rem <= 0 {
			weight, _ = moving.Pop()
			sum -= weight
		}
		if sum < l || len(moving.Weight) < w {
			weight, _ = all.First()
			if weight < 0 || sum+weight > l || len(moving.Weight)+1 > w {
				continue
			} else {
				sum += weight
				moving.Push(all.Pop())
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", count)
}
