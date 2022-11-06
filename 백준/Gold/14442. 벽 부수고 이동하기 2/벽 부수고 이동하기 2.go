package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	graph  = [][]string{}
)

type Queue struct {
	X []int
	Y []int
	Z []int
}

func (q *Queue) Enqueue(x, y, z int) {
	q.X = append(q.X, x)
	q.Y = append(q.Y, y)
	q.Z = append(q.Z, z)
}

func (q *Queue) Dequeue() (x, y, z int) {
	if len(q.X) <= 0 {
		return -1, -1, -1
	} else {
		x = q.X[0]
		y = q.Y[0]
		z = q.Z[0]
		if len(q.X) > 1 {
			q.X = q.X[1:]
			q.Y = q.Y[1:]
			q.Z = q.Z[1:]
		} else {
			q.X = []int{}
			q.Y = []int{}
			q.Z = []int{}
		}
	}
	return
}

func (q *Queue) Len() int {
	return len(q.X)
}
func bfs(n, m, k int) int {
	var visited = make([][][]int, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			visited[i][j] = make([]int, k+1)
		}
	}

	var cx, cy, cz int
	var nx, ny int
	var dx = []int{-1, 1, 0, 0}
	var dy = []int{0, 0, -1, 1}
	var q = new(Queue)
	q.Enqueue(0, 0, 0)
	visited[0][0][0] = 1
	for q.Len() > 0 {
		cx, cy, cz = q.Dequeue()
		if cx == n-1 && cy == m-1 {
			return visited[cx][cy][cz]
		}
		for i := 0; i < 4; i++ {
			nx = cx + dx[i]
			ny = cy + dy[i]
			if 0 <= nx && nx < n && 0 <= ny && ny < m {
				if graph[nx][ny] == "0" && visited[nx][ny][cz] == 0 {
					visited[nx][ny][cz] = visited[cx][cy][cz] + 1
					q.Enqueue(nx, ny, cz)
				} else if graph[nx][ny] == "1" && cz < k && visited[nx][ny][cz+1] == 0 {
					visited[nx][ny][cz+1] = visited[cx][cy][cz] + 1
					q.Enqueue(nx, ny, cz+1)
				}
			}
		}
	}
	return -1
}
func main() {
	defer writer.Flush()
	var n, m, k int
	var s string
	fmt.Fscanf(reader, "%d %d %d\n", &n, &m, &k)
	graph = make([][]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &s)
		graph[i] = strings.Split(s, "")
	}
	fmt.Fprintf(writer, "%d\n", bfs(n, m, k))
}
