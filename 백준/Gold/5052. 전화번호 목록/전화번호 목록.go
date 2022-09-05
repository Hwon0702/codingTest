package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func calculate(n int) {
	var ret = "YES"
	var strs = make([]string, n)
	var breakFlag = false
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &strs[i])
	}
	sort.Strings(strs)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if strings.HasPrefix(strs[j], strs[i]) {
				ret = "NO"
				breakFlag = true
				break
			}
		}
		if breakFlag {
			break
		}
	}
	fmt.Println(ret)

}
func main() {
	var tc, n int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d\n", &n)
		calculate(n)
	}
}
