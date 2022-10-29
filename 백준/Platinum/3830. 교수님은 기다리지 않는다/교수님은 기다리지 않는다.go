package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	parents []int
	dist    []int
)

func find(x int) int {
	if x == parents[x] {
		return x
	} else {
		r := find(parents[x])
		dist[x] += dist[parents[x]]
		parents[x] = r
	}
	return parents[x]
}

func union(x, y, r int) {
	rootX := parents[x]
	rootY := parents[y]
	if rootX != rootY {
		parents[rootY] = rootX
		dist[rootY] = (dist[x] + r) - dist[y]
	}

}

func main() {
	defer writer.Flush()
	var n, m int
	var a, b, c int
	var s string
	for {
		fmt.Fscanf(reader, "%d %d\n", &n, &m)
		if n == 0 && m == 0 {
			break
		}
		parents = make([]int, n+1)
		dist = make([]int, n+1)
		for i := 0; i <= n; i++ {
			parents[i] = i
		}
		for i := 0; i < m; i++ {
			s, _ = reader.ReadString('\n')
			s = strings.TrimSpace(s)
			Arr := strings.Split(s, " ")
			a, _ = strconv.Atoi(Arr[1])
			b, _ = strconv.Atoi(Arr[2])
			find(a)
			find(b)
			if Arr[0] == "!" {
				c, _ = strconv.Atoi(Arr[3])
				union(a, b, c)
			} else {
				if parents[a] == parents[b] {
					fmt.Fprintf(writer, "%d\n", dist[b]-dist[a])
				} else {
					fmt.Fprintf(writer, "UNKNOWN\n")
				}
			}
		}
	}
}
