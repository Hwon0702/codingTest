package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Find(arr []int, target int) int {
	for i, n := range arr {
		if target == n {
			return i
		}
	}
	return -1
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
		res[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if numbers[i] > numbers[j] {
				if res[i] < res[j]+1 {
					res[i] = res[j] + 1
				}
			}
		}
	}
	for _, v := range res {
		if max < v {
			max = v
		}
	}
	fmt.Fprintf(writer, "%d\n", max)
	idx := Find(res, max)
	var resList = []int{}
	for idx >= 0 {
		if res[idx] == max {
			resList = append(resList, numbers[idx])
			max--
		}
		idx--
	}
	sort.Ints(resList)
	for _, v := range resList {
		fmt.Fprintf(writer, "%d ", v)
	}
}
