package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	x []int
	y []int
	z []int
}

func (q *Queue) Enqueue(x, y, z int) {
	q.x = append(q.x, x)
	q.y = append(q.y, y)
	q.z = append(q.z, z)
}

func (q *Queue) Dequeue() (x, y, z int) {
	if len(q.x) <= 0 {
		return -1, -1, -1
	} else {
		x = q.x[0]
		y = q.y[0]
		z = q.z[0]
		if len(q.x) == 1 {
			q.x = []int{}
			q.y = []int{}
			q.z = []int{}
		} else {
			q.x = q.x[1:]
			q.y = q.y[1:]
			q.z = q.z[1:]
		}
	}
	return
}

func (q *Queue) Len() int {
	return len(q.x)
}

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var k, w, h int
	var cx, cy, cz int
	var nx, ny int
	var res int
	res = -1
	fmt.Fscanf(reader, "%d\n%d %d\n", &k, &w, &h)
	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}
	hx := []int{-2, -2, -1, 1, 2, 2, -1, 1}
	hy := []int{-1, 1, -2, -2, -1, 1, 2, 2}
	var graph = make([][]int, h)
	var visited = make([][][]int, h)
	for i := 0; i < h; i++ {
		graph[i] = make([]int, w)
		visited[i] = make([][]int, w)
		for j := 0; j < w; j++ {
			visited[i][j] = make([]int, k+1)
			fmt.Fscanf(reader, "%d ", &graph[i][j])
		}
	}
	var q = new(Queue)
	visited[0][0][0] = 1
	q.Enqueue(0, 0, 0)
	for q.Len() > 0 {
		cx, cy, cz = q.Dequeue()
		if cx == h-1 && cy == w-1 {
			res = visited[cx][cy][cz] - 1
			q = new(Queue)
			break
		}
		for i := 0; i < 4; i++ {
			nx = cx + dx[i]
			ny = cy + dy[i]
			if 0 <= nx && nx < h && 0 <= ny && ny < w && visited[nx][ny][cz] == 0 && graph[nx][ny] == 0 {
				visited[nx][ny][cz] = visited[cx][cy][cz] + 1
				q.Enqueue(nx, ny, cz)
			}
		}
		if cz < k {
			for i := 0; i < 8; i++ {
				nx = cx + hx[i]
				ny = cy + hy[i]
				if 0 <= nx && nx < h && 0 <= ny && ny < w {
					if graph[nx][ny] == 0 {
						if visited[nx][ny][cz+1] == 0 {
							visited[nx][ny][cz+1] = visited[cx][cy][cz] + 1
							q.Enqueue(nx, ny, cz+1)
						}
					}
				}
			}
		}
	}
	fmt.Fprintf(writer, "%d", res)
}
