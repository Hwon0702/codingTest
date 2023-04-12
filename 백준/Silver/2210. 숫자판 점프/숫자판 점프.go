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
	nums   = [][]string{}
	dx     = []int{-1, 1, 0, 0}
	dy     = []int{0, 0, -1, 1}
	res    = make(map[string]bool)
)

func dfs(x, y int, s string) {
	if len(s) == 6 {
		res[s] = true
		return
	}
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if 0 <= nx && nx < 5 && 0 <= ny && ny < 5 {
			dfs(nx, ny, s+nums[nx][ny])
		}
	}
}

func main() {
	defer writer.Flush()
	nums = make([][]string, 5)
	var s string
	for i := 0; i < 5; i++ {
		nums[i] = make([]string, 5)
		s, _ = reader.ReadString('\n')
		nums[i] = strings.Split(strings.TrimSpace(s), " ")
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			dfs(i, j, nums[i][j])
		}
	}
	fmt.Fprintf(writer, "%d\n", len(res))
}
