package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Queue struct {
	Node []int
}

func (q *Queue) Enqueue(n int) {
	q.Node = append(q.Node, n)
}

func (q *Queue) Dequeue() (ret int) {
	if len(q.Node) <= 0 {
		return -1
	} else {
		ret = q.Node[0]
		if len(q.Node) > 1 {
			q.Node = q.Node[1:]
		} else {
			q.Node = []int{}
		}
	}
	return ret
}

func (q *Queue) Len() int {
	return len(q.Node)
}
func makeBuilding(n, w int, cnt []int, times []float64, buildings [][]int) float64 {
	var resTime = make([]float64, n+1)
	var now int
	var q = new(Queue)
	for i, v := range cnt {
		if v == 0 {
			q.Enqueue(i)
			resTime[i] += times[i]
		}
	}
	for q.Len() > 0 {
		now = q.Dequeue()
		for _, v := range buildings[now] {
			cnt[v]--
			resTime[v] = math.Max(resTime[v], resTime[now]+times[v])
			if cnt[v] == 0 {
				q.Enqueue(v)
			}
		}
	}
	return resTime[w]
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc, n, k, start, end, w int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d %d\n", &n, &k)
		var buildings = make([][]int, n+1)
		var times = make([]float64, n+1)
		var cnt = make([]int, n+1)
		for j := 1; j < n+1; j++ {
			buildings[i] = []int{}
			fmt.Fscanf(reader, "%f ", &times[j])
		}
		for j := 0; j < k; j++ {
			fmt.Fscanf(reader, "%d %d\n", &start, &end)
			buildings[start] = append(buildings[start], end)
			cnt[end]++
		}
		fmt.Fscanf(reader, "%d\n", &w)
		fmt.Printf("%0.f\n", makeBuilding(n, w, cnt, times, buildings))
	}
}
