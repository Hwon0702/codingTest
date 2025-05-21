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
	var n, record, p int
	var rank = 1
	fmt.Fscanf(reader, "%d %d %d\n", &n, &record, &p)
	var records = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &records[i])
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i] > records[j]
	})
	if len(records) >= p && records[len(records)-1] >= record {
		rank = -1
	} else {
		for i := 0; i < len(records); i++ {
			if records[i] > record {
				rank += 1
			} else {
				break
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", rank)
}
