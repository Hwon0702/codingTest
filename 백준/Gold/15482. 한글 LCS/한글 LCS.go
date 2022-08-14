package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var s1, s2 string
	fmt.Fscanf(reader, "%s\n%s\n", &s1, &s2)
	s1Arr := strings.Split(s1, "")
	s2Arr := strings.Split(s2, "")
	var res = make([][]float64, len(s1)+1)
	for i := 0; i <= len(s1Arr); i++ {
		res[i] = make([]float64, len(s2)+1)
	}
	for i := 1; i <= len(s1Arr); i++ {
		for j := 1; j <= len(s2Arr); j++ {
			if s1Arr[i-1] == s2Arr[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = math.Max(res[i-1][j], res[i][j-1])
			}
		}
	}
	fmt.Fprintf(writer, "%.0f\n", res[len(s1Arr)][len(s2Arr)])
}
