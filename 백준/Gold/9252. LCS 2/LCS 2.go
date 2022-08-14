package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//writer := bufio.NewWriter(os.Stdout)
	//defer writer.Flush()
	var s1, s2 string
	fmt.Fscanf(reader, "%s\n%s\n", &s1, &s2)
	s1Arr := strings.Split(s1, "")
	s2Arr := strings.Split(s2, "")
	var res = make([][]string, len(s1Arr)+1)
	for i := 0; i <= len(s1Arr); i++ {
		res[i] = make([]string, len(s2Arr)+1)
	}
	for i := 1; i <= len(s1Arr); i++ {
		for j := 1; j <= len(s2Arr); j++ {
			if s1Arr[i-1] == s2Arr[j-1] {
				res[i][j] = res[i-1][j-1] + s1Arr[i-1]
			} else {
				if len(res[i-1][j]) > len(res[i][j-1]) {
					res[i][j] = res[i-1][j]
				} else {
					res[i][j] = res[i][j-1]
				}
			}
		}
	}
	if len(res[len(s1)][len(s2)]) > 0 {
		fmt.Printf("%d\n%s\n", len(res[len(s1)][len(s2)]), res[len(s1)][len(s2)])
	} else {
		fmt.Printf("%d\n", 0)
	}
}
