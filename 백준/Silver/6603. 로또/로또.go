package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var resultCombination = [][]int{}

func nCr(n int, nums, ans []int, r int) {
	if n == len(nums) {
		if len(ans) == r {
			var tmp = []int{}
			for _, v := range ans {
				tmp = append(tmp, v)
			}

			resultCombination = append(resultCombination, tmp)
		}
		return
	}
	ans = append(ans, nums[n])
	nCr(n+1, nums, ans, r)
	ans = ans[:len(ans)-1]
	nCr(n+1, nums, ans, r)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	n = 1
	for true {
		fmt.Fscanf(reader, "%d ", &n)
		if n == 0 {
			break
		}
		resultCombination = [][]int{}
		var numbers = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscanf(reader, "%d ", &numbers[i])
		}
		sort.Ints(numbers)
		nCr(0, numbers, []int{}, 6)
		for i := 0; i < len(resultCombination); i++ {
			for _, v := range resultCombination[i] {
				fmt.Fprintf(writer, "%d ", v)
			}
			fmt.Fprintln(writer, "")
		}
		fmt.Fprintln(writer, "")

	}
}
