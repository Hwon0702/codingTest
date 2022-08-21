package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func remove(s []int, index int) []int {
	ret := []int{}
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, good int
	fmt.Fscanf(reader, "%d\n", &n)
	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	sort.Ints(numbers)
	var left, right, sum int
	for i := 0; i < n; i++ {
		targets := remove(numbers, i)
		left, right = 0, len(targets)-1
		for left < right {
			sum = targets[left] + targets[right]
			if sum == numbers[i] {
				good += 1
				break
			} else if sum < numbers[i] {
				left++
			} else {
				right--
			}
		}
	}
	fmt.Println(good)
}
