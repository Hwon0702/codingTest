package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	result = 0
	n      = 0
	graph  = [][]string{}
)

func checkBombo() {
	for i := 0; i < n; i++ {
		cnt := 1
		for j := 0; j < (n - 1); j++ {
			if graph[i][j] == graph[i][j+1] {
				cnt += 1
				result = max(result, cnt)
			} else {
				cnt = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		cnt := 1
		for j := 0; j < (n - 1); j++ {
			if graph[j][i] == graph[j+1][i] {
				cnt += 1
				result = max(result, cnt)
			} else {
				cnt = 1
			}
		}
	}
}

func main() {
	defer writer.Flush()

	var str string
	fmt.Fscanf(reader, "%d\n", &n)
	graph = make([][]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &str)
		graph[i] = strings.Split(str, "")
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			graph[i][j], graph[i][j+1] = graph[i][j+1], graph[i][j] //가로 교체 
			checkBombo()
			graph[i][j+1], graph[i][j] = graph[i][j], graph[i][j+1] //복구

			graph[j][i], graph[j+1][i] = graph[j+1][i], graph[j][i] //세로 교체
			checkBombo()
			graph[j+1][i], graph[j][i] = graph[j][i], graph[j+1][i] //복구
		}
	}
	fmt.Fprintf(writer, "%d\n", result)
}
