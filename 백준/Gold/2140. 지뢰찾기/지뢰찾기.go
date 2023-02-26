package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	dx     = []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	dy     = []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
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
	if len(q.X) > 1 {
		q.X = q.X[1:]
		q.Y = q.Y[1:]
	} else {
		q.X = []int{}
		q.Y = []int{}
	}
	return
}

func (q *Queue) Len() int {
	return len(q.X)
}
func main() {
	defer writer.Flush()
	var n, res int
	var cx, cy int
	var nx, ny int
	var q = new(Queue)
	var tmp string
	fmt.Fscanf(reader, "%d\n", &n)
	var bomb = make([][]float64, n)
	var boards = make([][]string, n)
	for i := 0; i < n; i++ {
		bomb[i] = make([]float64, n)
		tmp, _ = reader.ReadString('\n')
		tmpArr := strings.Split(tmp, "")
		boards[i] = make([]string, n)
		for j := 0; j < n; j++ {
			if tmpArr[j] == "#" {
				q.Enqueue(i, j)
				bomb[i][j] = 1
			}
			boards[i][j] = tmpArr[j]
		}
	}
	for q.Len() > 0 {
		cx, cy = q.Dequeue()
		for i := 0; i < 9; i++ {
			nx = cx + dx[i]
			ny = cy + dy[i]
			if 0 <= nx && nx < n && 0 <= ny && ny < n {
				if boards[nx][ny] != "#" {
					v, _ := strconv.Atoi(boards[nx][ny])
					bomb[cx][cy] = math.Min(bomb[cx][cy], float64(v))
				}
			}
		}
		if bomb[cx][cy] > 0 {
			for i := 0; i < 9; i++ {
				nx = cx + dx[i]
				ny = cy + dy[i]
				if 0 <= nx && nx < n && 0 <= ny && ny < n {
					if boards[nx][ny] != "#" {
						v, _ := strconv.Atoi(boards[nx][ny])
						if v > 0 {
							boards[nx][ny] = fmt.Sprintf("%d", v-1)
						}
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if bomb[i][j] > 0 {
				res += 1
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
