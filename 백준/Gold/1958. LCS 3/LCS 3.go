package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s1, s2, s3 string
	fmt.Fscanf(reader, "%s\n%s\n%s", &s1, &s2, &s3)
	var res = make([][][]float64, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		res[i] = make([][]float64, len(s2)+1)
		for j := 0; j <= len(s2); j++ {
			res[i][j] = make([]float64, len(s3)+1)
		}
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			for k := 1; k <= len(s3); k++ {
				if s1[i-1] == s2[j-1] && s2[j-1] == s3[k-1] && s1[i-1] == s3[k-1] {
					res[i][j][k] = res[i-1][j-1][k-1] + 1
				} else {
					res[i][j][k] = math.Max(res[i-1][j][k], math.Max(res[i][j-1][k], res[i][j][k-1]))
				}
			}
		}
	}
	fmt.Printf("%0.f\n", res[len(s1)][len(s2)][len(s3)])
}
