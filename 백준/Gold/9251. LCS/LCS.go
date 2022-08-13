package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func searchIdx(target []string, v string) (ret []int) {
	for idx, t := range target {
		if v == t {
			ret = append(ret, idx)
		}
	}
	return ret
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var s1, s2 string
	fmt.Fscanf(reader, "%s\n%s\n", &s1, &s2)
	var res = make([][]float64, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		res[i] = make([]float64, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				res[i][j] = res[i-1][j-1] + 1
			} else {
				res[i][j] = math.Max(res[i-1][j], res[i][j-1])
			}
		}
	}
	fmt.Println(res[len(s1)][len(s2)])
}
