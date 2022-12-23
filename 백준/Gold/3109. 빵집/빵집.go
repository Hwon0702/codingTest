package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader    = bufio.NewReader(os.Stdin)
	writer    = bufio.NewWriter(os.Stdout)
	dx        = []int{-1, 0, 1}
	dy        = []int{1, 1, 1}
	cnt, h, w int
	graph     = [][]string{}
)

func move(x, y int) bool {
	graph[x][y] = "a"
	if y == w-1 {
		cnt++
		return true
	}
	for i := 0; i < 3; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if 0 <= nx && nx < h && 0 < ny && ny < w && graph[nx][ny] == "." {
			if move(nx, ny) {
				return true
			}
		}
	}
	return false
}
func main() {
	defer writer.Flush()
	fmt.Fscanf(reader, "%d %d\n", &h, &w)
	var s string
	graph = make([][]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscanf(reader, "%s\n", &s)
		graph[i] = strings.Split(s, "")
	}
	for i := 0; i < h; i++ {
		move(i, 0)
	}
	fmt.Fprintf(writer, "%d\n", cnt)
}
