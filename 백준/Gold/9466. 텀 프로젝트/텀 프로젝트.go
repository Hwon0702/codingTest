package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	cnt     = 0
	finish  = []bool{}
	visited = []bool{}
	reader  = bufio.NewReader(os.Stdin)
)

func dfs(v, i int, persons []int) {
	visited[v] = true
	var next = persons[v]
	if !visited[next] {
		dfs(next, i, persons)
	} else {
		if !finish[next] {
			i = next
			for i != v {
				cnt += 1
				i = persons[i]
			}
			cnt += 1
		}
	}
	finish[v] = true
}
func search(n int) {
	var persons = make([]int, n+1)
	visited = make([]bool, n+1)
	finish = make([]bool, n+1)
	visited[0] = true
	finish[0] = true
	for j := 1; j <= n; j++ {
		fmt.Fscanf(reader, "%d ", &persons[j])
	}
	for j := 1; j <= n; j++ {
		if !visited[j] {
			dfs(j, j, persons)
		}
	}
	fmt.Println(n - cnt)
}
func main() {

	var tc, n int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		cnt = 0
		fmt.Fscanf(reader, "%d\n", &n)
		search(n)
	}
}
