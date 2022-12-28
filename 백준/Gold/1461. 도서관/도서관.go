package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, m, b int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var pos, neg, count = []int{}, []int{}, []int{}

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &b)
		if b > 0 {
			pos = append(pos, b)
		} else {
			neg = append(neg, b*(-1))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pos)))
	sort.Sort(sort.Reverse(sort.IntSlice(neg)))
	for i := 0; i < int(len(pos)/m); i++ {
		count = append(count, pos[i*m])
	}
	if len(pos)%m > 0 {
		count = append(count, pos[int(len(pos)/m)*m])
	}
	for i := 0; i < int(len(neg)/m); i++ {
		count = append(count, neg[i*m])
	}
	if len(neg)%m > 0 {
		count = append(count, neg[int(len(neg)/m)*m])
	}
	sort.Ints(count)
	res := count[len(count)-1]
	for i := 0; i < len(count)-1; i++ {
		res += count[i] * 2
	}
	fmt.Fprintf(writer, "%d", res)
}
