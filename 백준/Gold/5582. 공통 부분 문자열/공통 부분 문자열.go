package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s1, s2 string
	var cnt float64
	fmt.Fscanf(reader, "%s\n%s\n", &s1, &s2)
	var res = make([][]float64, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		res[i] = make([]float64, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
				cnt = math.Max(cnt, res[i][j])
			}
		}
	}
	fmt.Println(cnt)
}
