package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 조합
// 순열은
var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func checkWidth(graphs [][]string, n int) int {
	cnt := 0
	graph := make([][]string, n)
	for i := 0; i < n; i++ {
		graph[i] = append(graph[i], graphs[i]...)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if graph[i][j] == graph[i][j+1] && graph[i][j] == "." {
				cnt++
				for k := j; k < n; k++ {
					if graph[i][k] == "X" {
						break
					}
					graph[i][k] = "B"
				}
			}
		}
	}
	return cnt
}

func checkHeight(graphs [][]string, n int) int {
	cnt := 0

	for i := 0; i < n-1; i++ {
		for j := 0; j < n; j++ {
			if graphs[i][j] == graphs[i+1][j] && graphs[i][j] == "." {
				cnt++
				for k := i; k < n; k++ {
					if graphs[k][j] == "X" {
						break
					}
					graphs[k][j] = "B"
				}
			}
		}
	}
	return cnt
}
func main() {
	var n int
	defer writer.Flush()
	fmt.Fscanf(reader, "%d\n", &n)
	var graphs = make([][]string, n)
	var str string
	for i := 0; i < n; i++ {
		str, _ = reader.ReadString('\n')
		graphs[i] = strings.Split(strings.TrimSpace(str), "")
	}
	a := checkWidth(graphs, n)
	b := checkHeight(graphs, n)
	fmt.Fprintf(writer, "%d %d", a, b)

}
