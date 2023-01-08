package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		var n, m, cnt int
		fmt.Fscanf(reader, "%d %d\n", &n, &m)
		var books = make([]bool, n+1)
		var req = make([][2]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d %d\n", &req[j][0], &req[j][1])
		}
		sort.SliceStable(req, func(i, j int) bool {
			return req[i][1] < req[j][1]
		})
		for j := 0; j < len(req); j++ {
			for k := req[j][0]; k <= req[j][1]; k++ {
				if books[k] == false {
					cnt++
					books[k] = true
					break
				}
			}
		}
		fmt.Fprintf(writer, "%d\n", cnt)
	}
}
