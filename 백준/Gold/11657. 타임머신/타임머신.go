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
	dist   = []int{}
	n      int
	graph  = [][]int{}
)

func bellmanFord() bool {
	var current, next, cost int
	dist[1] = 0
	for i := 0; i <= n; i++ {
		for j := 0; j < len(graph); j++ {
			current = graph[j][0]
			next = graph[j][1]
			cost = graph[j][2]
			if dist[current] != math.MaxInt && dist[next] > dist[current]+cost {
				dist[next] = dist[current] + cost
				if i == n-1 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	defer writer.Flush()
	var m, s, e, c int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	graph = make([][]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &s, &e, &c)
		graph[i] = []int{s, e, c}
	}

	dist = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dist[i] = math.MaxInt
	}
	if bellmanFord() {
		fmt.Fprintf(writer, "%d\n", -1)
	} else {
		for i := 2; i <= n; i++ {
			if dist[i] != math.MaxInt {
				fmt.Fprintf(writer, "%d\n", dist[i])
			} else {
				fmt.Fprintf(writer, "%d\n", -1)
			}
		}
	}
}
