package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Fscanf(reader, "%d\n", &n)
	var numbers = make([]int, n)
	var res = make([]int, n)
	var max int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if numbers[i] > numbers[j] && res[i] < res[j] {
				res[i] = res[j]
			}
		}
		res[i] += numbers[i]

	}
	for _, v := range res {
		if max < v {
			max = v
		}
	}
	fmt.Fprintf(writer, "%d", max)
}
