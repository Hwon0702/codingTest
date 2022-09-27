package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, start, end, mid int
	fmt.Fscanf(reader, "%d\n", &n)
	var numbers = make([]int, n)
	var nums = []int{0}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])

	}
	for _, v := range numbers {
		if nums[len(nums)-1] < v {
			nums = append(nums, v)
		} else {
			start = 0
			end = len(nums)
			for start < end {
				mid = int((start + end) / 2)
				if nums[mid] < v {
					start = mid + 1
				} else {
					end = mid
				}
			}
			nums[end] = v
		}
	}
	fmt.Fprintf(writer, "%d", len(nums) - 1)
}
