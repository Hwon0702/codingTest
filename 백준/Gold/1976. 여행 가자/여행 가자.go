package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func floydWarshal(graph [][]string, n int) [][]string {
	var connect = make([][]string, n)
	copy(connect, graph)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					connect[i][j] = "1"
				}
				if graph[i][k] == "1" && graph[k][j] == "1" {
					connect[i][j] = "1"
				}
			}
		}
	}
	return connect
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m, d int
	var s string
	var race = make([]int, 0)
	var flag = true
	fmt.Fscanf(reader, "%d\n%d\n", &n, &m)
	var graph = make([][]string, n)
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		sArr := strings.Split(strings.TrimSpace(s), " ")
		graph[i] = sArr
	}
	s, _ = reader.ReadString('\n')
	sArr := strings.Split(strings.TrimSpace(s), " ")
	for _, v := range sArr {
		d, _ = strconv.Atoi(v)
		race = append(race, d-1)
	}
	connected := floydWarshal(graph, n)
	for i := 1; i < len(race); i++ {
		if connected[race[i-1]][race[i]] == "1" {
			continue
		} else {
			flag = false
		}
	}
	if flag {
		fmt.Fprintf(writer, "YES")
	} else {
		fmt.Fprintf(writer, "NO")
	}
}
