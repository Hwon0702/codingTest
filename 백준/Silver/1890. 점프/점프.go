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
	var n, nx, ny int
	fmt.Fscanf(reader, "%d\n", &n)
	var graph = make([][]int, n)
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
			res[i][j] = 0
		}
	}
	res[0][0] = 1
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if x == n-1 && y == n-1 {
				fmt.Fprintf(writer, "%d", res[x][y])
				break
			}
			nx = x + graph[x][y]
			ny = y + graph[x][y]
			if 0 <= nx && nx < n {
				res[nx][y] += res[x][y]
			}
			if 0 <= ny && ny < n {
				res[x][ny] += res[x][y]
			}
		}
	}
}
