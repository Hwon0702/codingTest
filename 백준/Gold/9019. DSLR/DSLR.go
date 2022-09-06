package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	N []int
	M []string
}

func (q *Queue) Enqueue(n int, m string) {
	q.N = append(q.N, n)
	q.M = append(q.M, m)
}

func (q *Queue) Dequeue() (n int, m string) {
	if len(q.N) <= 0 {
		return -1, ""
	} else {
		n = q.N[0]
		m = q.M[0]
		if len(q.N) == 0 {
			q.N = []int{}
			q.M = []string{}
		} else {
			q.N = q.N[1:]
			q.M = q.M[1:]
		}
		return n, m
	}
}

func (q *Queue) Len() int {
	return len(q.N)
}

func D(n int) int {
	if n*2 > 9999 {
		return n * 2 % 10000
	}
	return n * 2
}
func S(n int) int {
	if n == 0 {
		return 9999
	}
	return n - 1
}
func L(n int) int {
	return (10*n + int(n/1000)) % 10000
}
func R(n int) int {
	return int((n/10)+(n%10)*1000) % 10000
}

func bfs(s, e int) {
	var visited = make([]bool, 10000)
	var queue = new(Queue)
	queue.Enqueue(s, "")
	var r int
	for queue.Len() > 0 {
		n, m := queue.Dequeue()
		visited[n] = true
		if n == e {
			fmt.Println(m)
			break
		}
		r = D(n)
		if !visited[r] {
			visited[r] = true
			queue.Enqueue(r, m+"D")
		}
		r = S(n)
		if !visited[r] {
			visited[r] = true
			queue.Enqueue(r, m+"S")
		}
		r = L(n)
		if !visited[r] {
			visited[r] = true
			queue.Enqueue(r, m+"L")
		}
		r = R(n)
		if !visited[r] {
			visited[r] = true
			queue.Enqueue(r, m+"R")
		}
	}
	return
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	var a, b int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		bfs(a, b)
	}
}
