package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	var ropes = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &ropes[i])
	}
	sort.Slice(ropes, func(i, j int) bool {
		return ropes[i] > ropes[j]
	})
	for i := 0; i < n; i++ {
		ropes[i] = ropes[i] * (i + 1)
	}
	sort.Ints(ropes)
	fmt.Println(ropes[len(ropes)-1])
}
