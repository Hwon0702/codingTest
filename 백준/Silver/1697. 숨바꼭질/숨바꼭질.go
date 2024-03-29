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
	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var queue = []int{}
	var visited = make([]int, 100001)
	queue = append(queue, n)
	for len(queue) > 0 {
		now := queue[0]
		queue = queue[1:]
		if now >= 0 {
			if now == k {
				fmt.Fprintf(writer, "%d\n", visited[now])
				break
			} else {
				if 0 <= now-1 && now-1 <= 100000 && visited[now-1] == 0 {
					visited[now-1] = visited[now] + 1
					queue = append(queue, now-1)
				}
				if 0 <= now+1 && now+1 <= 100000 && visited[now+1] == 0 {
					visited[now+1] = visited[now] + 1
					queue = append(queue, now+1)
				}
				if 0 <= now*2 && now*2 <= 100000 && visited[now*2] == 0 {
					visited[now*2] = visited[now] + 1
					queue = append(queue, now*2)
				}
			}
		}
	}
}
