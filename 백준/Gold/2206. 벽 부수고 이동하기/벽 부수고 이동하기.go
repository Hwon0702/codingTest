package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var w, h int
	var cx, cy, cz int
	var nx, ny int
	var res int
	var s string
	res = -1
	fmt.Fscanf(reader, "%d %d\n", &h, &w)
	dx := []int{0, 0, -1, 1}
	dy := []int{1, -1, 0, 0}
	var graph = make([][]string, h)
	var visited = make([][][]int, h)
	for i := 0; i < h; i++ {
		visited[i] = make([][]int, w)
		for j := 0; j < w; j++ {
			visited[i][j] = make([]int, 2)
		}
		fmt.Fscanf(reader, "%s\n", &s)
		graph[i] = strings.Split(s, "")
	}
	var q = new(Queue)
	visited[0][0][0] = 1
	q.Enqueue(0, 0, 0)
	for q.Len() > 0 {
		cx, cy, cz = q.Dequeue()
		if cx == h-1 && cy == w-1 {
			res = visited[cx][cy][cz]
			break
		}
		for i := 0; i < 4; i++ {
			nx = cx + dx[i]
			ny = cy + dy[i]
			if 0 <= nx && nx < h && 0 <= ny && ny < w && visited[nx][ny][cz] == 0 {
				if graph[nx][ny] == "0" {
					visited[nx][ny][cz] = visited[cx][cy][cz] + 1
					q.Enqueue(nx, ny, cz)
				} else if graph[nx][ny] == "1" && cz == 0 {
					visited[nx][ny][cz+1] = visited[cx][cy][cz] + 1
					q.Enqueue(nx, ny, cz+1)
				}
			}
		}
	}
	fmt.Fprintf(writer, "%d", res)
}
