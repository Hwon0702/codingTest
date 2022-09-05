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
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &strs[i])
	}
	sort.Strings(strs)

	for i := 0; i < n-1; i++ {
		if strings.HasPrefix(strs[i+1], strs[i]) {
			ret = "NO"
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
