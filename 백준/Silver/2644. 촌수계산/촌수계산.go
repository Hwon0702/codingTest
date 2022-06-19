package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, a, b, m int

	fmt.Fscanf(reader, "%d\n%d %d\n%d\n", &n, &a, &b, &m)
	var persons = make([][]int, n+1)
	var visit = make([]bool, n+1)
	var result = make([]int, n+1)
	var success = false
	for i := 0; i <= n; i++ {
		visit[i] = false
	}
	visit[0] = true
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscanf(reader, "%d %d\n", &x, &y)
		persons[x] = append(persons[x], y)
		persons[y] = append(persons[y], x)

	}
	var queue = persons[a]
	queue = append(queue, a) //첫번째 노드 방문 처리 후 큐에 담음
	if len(visit) <= 0 {
		visit[a] = true
		result[a] += 1
	}

	for len(queue) != 0 {
		now := queue[0]
		if now == b {
			success = true
			fmt.Println(result[b] + 1)
			break
		}
		queue = queue[1:]

		for _, visitNode := range persons[now] {
			if visit[visitNode] != true {
				visit[visitNode] = true
				result[visitNode] = result[now] + 1
				queue = append(queue, visitNode)
			}
		}

	}
	if !success {
		fmt.Println(-1)
	}

}
