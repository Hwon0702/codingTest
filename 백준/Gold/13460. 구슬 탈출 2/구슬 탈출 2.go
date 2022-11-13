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
	q      = &Queue{}
	n, m   int
)

type Queue struct {
	BX  []int
	BY  []int
	RX  []int
	RY  []int
	Cnt []int
}

func (q *Queue) Enqueue(bx, by, rx, ry, cnt int) {
	q.BX = append(q.BX, bx)
	q.BY = append(q.BY, by)
	q.RX = append(q.RX, rx)
	q.RY = append(q.RY, ry)
	q.Cnt = append(q.Cnt, cnt)
}

func (q *Queue) Dequeue() (bx, by, rx, ry, cnt int) {
	if len(q.RX) <= 0 {
		return -1, -1, -1, -1, -1
	} else {
		bx = q.BX[0]
		by = q.BY[0]
		rx = q.RX[0]
		ry = q.RY[0]
		cnt = q.Cnt[0]
		if len(q.RX) > 1 {
			q.RX = q.RX[1:]
			q.RY = q.RY[1:]
			q.BX = q.BX[1:]
			q.BY = q.BY[1:]
			q.Cnt = q.Cnt[1:]
		} else {
			q.RX = []int{}
			q.RY = []int{}
			q.BX = []int{}
			q.BY = []int{}
			q.Cnt = []int{}
		}
	}
	return
}

func (q *Queue) Len() int {
	return len(q.RX)
}

func move(x, y, dx, dy int) (int, int, int) {
	var cnt = 0
	for graph[x+dx][y+dy] != "#" && graph[x][y] != "O" {
		x += dx
		y += dy
		cnt += 1
	}
	return x, y, cnt
}
func sol() int {
	var visited = make([][][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][][]bool, m)
		for j := 0; j < m; j++ {
			visited[i][j] = make([][]bool, n)
			for k := 0; k < n; k++ {
				visited[i][j][k] = make([]bool, m)
			}
		}
	}

	var rx, ry, bx, by, cnt int
	var nrx, nry, rMove, nbx, nby, bMove int
	var dx = []int{-1, 1, 0, 0}
	var dy = []int{0, 0, -1, 1}

	for q.Len() > 0 {
		bx, by, rx, ry, cnt = q.Dequeue()
		if cnt > 10 {
			return -1
		}
		for i := 0; i < 4; i++ {
			nrx, nry, rMove = move(rx, ry, dx[i], dy[i])
			nbx, nby, bMove = move(bx, by, dx[i], dy[i])
			if graph[nbx][nby] == "O" {
				continue
			}
			if graph[nrx][nry] == "O" {
				return cnt
			}
			if nrx == nbx && nry == nby {
				if rMove > bMove {
					nrx -= dx[i]
					nry -= dy[i]
				} else {
					nbx -= dx[i]
					nby -= dy[i]
				}
			}
			if visited[nrx][nry][nbx][nby] == false {
				visited[nrx][nry][nbx][nby] = true
				q.Enqueue(nbx, nby, nrx, nry, cnt+1)
			}
		}
	}
	return -1
}
func main() {
	defer writer.Flush()
	var s string
	var rx, ry, bx, by int
	q = new(Queue)
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	graph = make([][]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &s)
		sArr := strings.Split(s, "")
		for j := 0; j < len(sArr); j++ {
			if sArr[j] == "B" {
				bx, by = i, j
				sArr[j] = "."
			} else if sArr[j] == "R" {
				rx, ry = i, j
				sArr[j] = "."
			}
		}
		graph[i] = sArr
	}
	q.Enqueue(bx, by, rx, ry, int(1))
	fmt.Fprintf(writer, "%d\n", sol())
}
