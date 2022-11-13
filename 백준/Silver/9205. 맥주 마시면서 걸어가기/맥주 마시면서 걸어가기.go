package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Queue struct {
	X []float64
	Y []float64
}

func (q *Queue) Enqueue(x, y float64) {
	q.X = append(q.X, x)
	q.Y = append(q.Y, y)
}
func (q *Queue) Dequeue() (error, x float64, y float64) {
	if len(q.X) <= 0 {
		return 1, -1, -1
	} else {
		x = q.X[0]
		y = q.Y[0]
		if len(q.X) == 0 {
			q.X = []float64{}
			q.Y = []float64{}
		} else {
			q.X = q.X[1:]
			q.Y = q.Y[1:]
		}
		return 0, x, y
	}
}
func (q *Queue) Len() int {
	return len(q.X)
}

func bfs(home []float64, store [][]float64, rock []float64) string {
	var visited = make([]bool, len(store))
	var queue = new(Queue)
	var ret = "sad"
	var nx, ny float64
	queue.Enqueue(home[0], home[1])
	for queue.Len() > 0 {
		e, x, y := queue.Dequeue()
		if e > 0 {
			break
		}
		if math.Abs(x-rock[0])+math.Abs(y-rock[1]) <= 1000 {
			ret = "happy"
			break
		}
		for i := 0; i < len(store); i++ {
			if !visited[i] {
				nx, ny = store[i][0], store[i][1]
				if math.Abs(x-nx)+math.Abs(y-ny) <= 1000 {
					queue.Enqueue(nx, ny)
					visited[i] = true
				}
			}
		}
	}
	return ret
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc, storeCnt int
	var home = make([]float64, 2)
	var rock = make([]float64, 2)
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d\n%f %f\n", &storeCnt, &home[0], &home[1])
		var store = make([][]float64, storeCnt)
		for j := 0; j < storeCnt; j++ {
			store[j] = make([]float64, 2)
			fmt.Fscanf(reader, "%f %f\n", &store[j][0], &store[j][1])
		}
		fmt.Fscanf(reader, "%f %f\n", &rock[0], &rock[1])
		fmt.Println(bfs(home, store, rock))
	}
}
