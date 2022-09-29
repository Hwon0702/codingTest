package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, a int
	var s string
	var max float64
	fmt.Fscanf(reader, "%d\n", &n)
	var graph = [][]float64{}
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSpace(s)
		strArr := strings.Split(s, " ")
		var tmp = []float64{}
		for j := 0; j < len(strArr); j++ {
			a, _ = strconv.Atoi(strArr[j])
			tmp = append(tmp, float64(a))
		}
		graph = append(graph, tmp)
	}
	var res = graph
	for i := 1; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				res[i][j] = res[i-1][j] + graph[i][j]
			} else if j == i {
				res[i][j] = res[i-1][j-1] + graph[i][j]
			} else {
				res[i][j] = math.Max(res[i-1][j-1], res[i-1][j]) + graph[i][j]
			}
		}
	}
	for _, v := range res[len(res)-1] {
		max = math.Max(v, max)
	}
	fmt.Printf("%0.f", max)
}
