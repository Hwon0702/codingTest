package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader    = bufio.NewReader(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	visited   = [][]bool{}
	graph     = [][]int{}
	dx        = []int{-1, 1, 0, 0}
	dy        = []int{0, 0, -1, 1}
	islandNum = 1
	n         = 0
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
		return -1, -1
	}
	x = q.X[0]
	y = q.Y[0]
	if len(q.X) == 1 {
		q.X = []int{}
		q.Y = []int{}
	} else {
		q.X = q.X[1:]
		q.Y = q.Y[1:]
	}
	return
}

func (q *Queue) Len() int {
	return len(q.X)
}

func devideIsland(i, j int) {
	island := new(Queue)
	island.Enqueue(i, j)
	var cx, cy int
	var nx, ny int
	for island.Len() > 0 {
		cx, cy = island.Dequeue()
		for i := 0; i < 4; i++ {
			nx = cx + dx[i]
			ny = cy + dy[i]
			if 0 <= nx && nx < n && 0 <= ny && ny < n && visited[nx][ny] == false && graph[nx][ny] == 1 {
				visited[nx][ny] = true
				graph[nx][ny] = islandNum
				island.Enqueue(nx, ny)
			}
		}
	}
}

func makingBridge(startIslandNum int) float64 {
	bridge := new(Queue)
	length := make([][]float64, n)
	for i := 0; i < n; i++ {
		length[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			length[i][j] = math.MaxFloat64
			if graph[i][j] == startIslandNum {
				length[i][j] = 0
				bridge.Enqueue(i, j)
			}
		}
	}

	var cx, cy int
	var nx, ny int
	for bridge.Len() > 0 {
		cx, cy = bridge.Dequeue()
		for i := 0; i < 4; i++ {
			nx = cx + dx[i]
			ny = cy + dy[i]
			if 0 <= nx && nx < n && 0 <= ny && ny < n {
				if graph[nx][ny] > 0 && graph[nx][ny] != startIslandNum && visited[nx][ny] {
					return length[cx][cy]
				} else if graph[nx][ny] == 0 && length[nx][ny] == math.MaxFloat64 {
					length[nx][ny] = length[cx][cy] + 1
					bridge.Enqueue(nx, ny)
				}
			}
		}
	}
	return math.MaxFloat64
}
func main() {
	defer writer.Flush()
	fmt.Fscanf(reader, "%d\n", &n)
	graph = make([][]int, n)
	visited = make([][]bool, n)
	var res float64
	res = math.MaxFloat64
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == 1 && visited[i][j] == false {
				visited[i][j] = true
				graph[i][j] = islandNum
				devideIsland(i, j)
				islandNum++
			}
		}
	}
	for i := 1; i <= islandNum; i++ {
		res = math.Min(res, makingBridge(i))
	}
	fmt.Fprintf(writer, "%0.f\n", res)
}
