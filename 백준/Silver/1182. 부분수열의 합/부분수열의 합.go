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

	var n, s, sum, cnt int
	fmt.Fscanf(reader, "%d %d\n", &n, &s)
	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	for i := 1; i <= n; i++ {
		nCr(0, numbers, []int{}, i)

	}
	for _, v := range resultCombination {
		sum = 0
		for _, num := range v {
			sum += num
		}
		if sum == s {
			cnt++
		}
	}
	//fmt.Println(resultCombination)
	fmt.Fprintf(writer, "%d\n", cnt)
}
