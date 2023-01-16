package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	parent = []int{}
	n      int
)

func find(x int) int {
	if parent[x] == x {
		return x
	}
	p := find(parent[x])
	parent[x] = p
	return p
}
func union(x, y int) {
	x, y = find(x), find(y)
	if x == y {
		return
	} else if x < y {
		parent[y] = x
	} else {
		parent[x] = y
	}
	return
}

func main() {
	defer writer.Flush()
	var m, op, a, b int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	parent = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		parent[i] = i
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &op, &a, &b)
		if op == 0 {
			union(a, b)
		} else if op == 1 {
			if find(a) == find(b) {
				fmt.Fprintf(writer, "YES\n")
			} else {
				fmt.Fprintf(writer, "NO\n")
			}
		}
	}
}
