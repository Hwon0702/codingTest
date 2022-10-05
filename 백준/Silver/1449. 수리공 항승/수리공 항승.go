package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, l int
	fmt.Fscanf(reader, "%d %d\n", &n, &l)
	var leak = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &leak[i])
	}
	sort.Ints(leak)
	start := leak[0]
	end := leak[0] + l
	cnt := 1
	for i := 0; i < n; i++ {
		if start <= leak[i] && leak[i] < end {
			continue
		} else {
			start = leak[i]
			end = leak[i] + l
			cnt++
		}
	}
	fmt.Println(cnt)
}
