package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var k, n int
	fmt.Fscanf(reader, "%d %d\n", &k, &n) //항상 k<=n
	var lans = make([]int, k)
	var sum int
	for i := 0; i < k; i++ {
		var length int
		fmt.Fscanf(reader, "%d\n", &length)
		lans[i] = length
		sum += length
	}
	sort.Ints(lans)
	fmt.Fprintf(writer, "%d", ModuleLength(n, lans))
}

func ModuleLength(n int, lans []int) (max int) {
	start := 1
	end := lans[len(lans)-1]
	for start <= end {
		mid := (start + end) / 2
		var totalCount int
		for i := 0; i < len(lans); i++ {
			totalCount += lans[i] / mid
		}
		if totalCount >= n {
			if max < mid {
				max = mid
				start = mid + 1
			}
		} else if totalCount < n {
			end = mid - 1
		}
	}
	return
}
