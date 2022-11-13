package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Queue struct {
	X []int
	Y []int
}

func (q *Queue) Enqueue(x, y int) {
	q.X = append(q.X, x)
	q.Y = append(q.Y, y)
}
func (q *Queue) Dequeue() (x, y int) {
	if len(q.X) <= 0 {
		x = -1
		y = -1
	} else {
		x = q.X[0]
		y = q.Y[0]
		if len(q.X) == 1 {
			q.X = []int{}
			q.Y = []int{}
		} else {
			q.X = q.X[1:]
			q.Y = q.Y[1:]
		}
	}
	return x, y
}
func (q *Queue) Size() int {
	return len(q.X)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var w, h, n int
	fmt.Fscanf(reader, "%d %d\n", &w, &h)
	var graph = make([][]int, h)
	var q = new(Queue)
	for y := 0; y < h; y++ {
		graph[y] = make([]int, w)
		for x := 0; x < w; x++ {
			fmt.Fscanf(reader, "%d ", &n)
			graph[y][x] = n
			if n == 1 {
				q.Enqueue(x, y)
			}
		}
	}
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for len(q.X) > 0 {
		x, y := q.Dequeue()
		if x < 0 || y < 0 {
			break
		}
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if 0 <= nx && nx < w && 0 <= ny && ny < h && graph[ny][nx] == 0 {
				graph[ny][nx] = graph[y][x] + 1

				q.Enqueue(nx, ny)
			}
		}
	}
	var impossible = false
	var count float64
	for i := 0; i < h; i++ {
		if impossible {
			break
		}
		for j := 0; j < w; j++ {
			if graph[i][j] == 0 {
				impossible = true
				break
			} else {
				count = math.Max(count, float64(graph[i][j]))
			}
		}
	}
	if impossible {
		fmt.Println(-1)
	} else {
		fmt.Println(count - 1)
	}

}
