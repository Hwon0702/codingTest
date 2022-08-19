package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, c, count int
	var start, end int
	fmt.Fscanf(reader, "%d %d\n", &n, &c)
	var distance = make([]int, n)
	var res = []int{}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &distance[i])
	}
	sort.Ints(distance)
	start = 1
	end = distance[n-1] - distance[0]
	var current = distance[0]
	var mid = (start + end) / 2
	for start <= end {
		count = 1
		mid = (start + end) / 2
		current = distance[0]
		for _, x := range distance {
			if current+mid <= x {
				count++
				current = x
			}
		}
		if count >= c {
			start = mid + 1
			res = append(res, mid)
		} else {
			end = mid - 1
		}
	}
	sort.Ints(res)
	fmt.Println(res[len(res)-1])
}
