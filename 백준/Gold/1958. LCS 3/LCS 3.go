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
	var s1, s2, s3 string
	fmt.Fscanf(reader, "%s\n%s\n%s", &s1, &s2, &s3)
	s1Arr := strings.Split(s1, "")
	s2Arr := strings.Split(s2, "")
	s3Arr := strings.Split(s3, "")
	var res = make([][][]float64, len(s1Arr)+1)
	for i := 0; i <= len(s1Arr); i++ {
		res[i] = make([][]float64, len(s2Arr)+1)
		for j := 0; j <= len(s2Arr); j++ {
			res[i][j] = make([]float64, len(s3Arr)+1)
		}
	}
	for i := 1; i <= len(s1Arr); i++ {
		for j := 1; j <= len(s2Arr); j++ {
			for k := 1; k <= len(s3Arr); k++ {
				if s1Arr[i-1] == s2Arr[j-1] && s2Arr[j-1] == s3Arr[k-1] && s1Arr[i-1] == s3Arr[k-1] {
					res[i][j][k] = res[i-1][j-1][k-1] + 1
				} else {
					res[i][j][k] = math.Max(res[i-1][j][k], math.Max(res[i][j-1][k], res[i][j][k-1]))
				}
			}
		}
	}
	fmt.Printf("%0.f\n", res[len(s1Arr)][len(s2Arr)][len(s3Arr)])
}
