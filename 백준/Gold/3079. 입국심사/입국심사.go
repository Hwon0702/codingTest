package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var left, right, n, m, res, person, mid int
	left = math.MaxInt
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var times = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &times[i])
		if left > times[i] {
			left = times[i]
		}
		if right < times[i] {
			right = times[i]
		}
	}
	right = right * m
	res = right
	for left <= right {
		person = 0
		mid = (left + right) / 2
		for i := 0; i < n; i++ {
			person += mid / times[i]
		}
		if person >= m {
			right = mid - 1
			if res > mid {
				res = mid
			}
		} else {
			left = mid + 1
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
