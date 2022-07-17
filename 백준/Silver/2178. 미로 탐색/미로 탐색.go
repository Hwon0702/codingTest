package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	Nodes [][]int
}

func (q *Queue) Size() int {
	return len(q.Nodes)
}
func (q *Queue) Push(n []int) {
	q.Nodes = append(q.Nodes, n)
}
func (q *Queue) Pop() (ret []int) {
	if q.Size() <= 0 {
		return []int{}
	}
	ret = q.Nodes[0]
	if q.Size() == 1 {
		q.Nodes = [][]int{}
	} else {
		q.Nodes = q.Nodes[1:]
	}
	return ret
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var maze = make([][]int, n)
	for i := 0; i < n; i++ {
		maze[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscanf(reader, "%s\n", &s)
		strArr := strings.Split(s, "")
		for idx, v := range strArr {
			n, _ := strconv.Atoi(v)
			maze[i][idx] = n
		}

	}
	var dx = []int{1, -1, 0, 0}
	var dy = []int{0, 0, 1, -1}
	var queue = new(Queue)
	queue.Push([]int{0, 0}) //시작 노드 삽입
	for queue.Size() >= 1 {
		q := queue.Pop()
		for i := 0; i < 4; i++ {
			nx, ny := q[0]+dx[i], q[1]+dy[i]
			if 0 <= nx && nx < n && 0 <= ny && ny < m {
				if maze[nx][ny] == 1 {
					maze[nx][ny] = maze[q[0]][q[1]] + 1
					queue.Push([]int{nx, ny})
				}
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", maze[n-1][m-1])
}
