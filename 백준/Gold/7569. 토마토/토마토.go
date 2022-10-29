package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type Queue struct {
	X []int
	Y []int
	Z []int
}

func (q *Queue) Push(x, y, z int) {
	q.X = append(q.X, x)
	q.Y = append(q.Y, y)
	q.Z = append(q.Z, z)
}
func (q *Queue) Pop() (x, y, z int) {
	if len(q.X) <= 0 {
		return -1, -1, -1
	}
	x = q.X[0]
	y = q.Y[0]
	z = q.Z[0]
	if len(q.X) == 1 {
		q.X = []int{}
		q.Y = []int{}
		q.Z = []int{}
	} else {
		q.X = q.X[1:]
		q.Y = q.Y[1:]
		q.Z = q.Z[1:]
	}
	return
}
func (q *Queue) Size() int {
	return len(q.X)
}
func main() {
	defer writer.Flush()
	var m, n, h int
	var cx, cy, cz int
	var nx, ny, nz int
	var cnt float64
	fmt.Fscanf(reader, "%d %d %d\n", &m, &n, &h)
	var q = new(Queue)
	var graph = make([][][]float64, h)
	for i := 0; i < h; i++ {
		graph[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			graph[i][j] = make([]float64, m)
			for k := 0; k < m; k++ {
				fmt.Fscanf(reader, "%f ", &graph[i][j][k])
				if graph[i][j][k] == 1 {
					q.Push(i, j, k)
				}
			}
		}
	}

	var dx = []int{-1, 1, 0, 0, 0, 0}
	var dy = []int{0, 0, -1, 1, 0, 0}
	var dz = []int{0, 0, 0, 0, -1, 1}

	for q.Size() > 0 {
		cx, cy, cz = q.Pop()
		for i := 0; i < 6; i++ {
			nx = cx + dx[i]
			ny = cy + dy[i]
			nz = cz + dz[i]
			if 0 <= nx && nx < h && 0 <= ny && ny < n && 0 <= nz && nz < m && graph[nx][ny][nz] == 0 && graph[cx][cy][cz] > 0 {
				q.Push(nx, ny, nz)
				graph[nx][ny][nz] = graph[cx][cy][cz] + 1
			}
		}

	}
	for i := 0; i < h; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				if graph[i][j][k] == 0 {
					fmt.Println(-1)
					os.Exit(0)
				} else {
					cnt = math.Max(cnt, graph[i][j][k])
				}
			}
		}
	}
	fmt.Fprintf(writer, "%0.f", cnt-1)
}
