package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, num, sum int
	positive := []int{}
	negative := []int{}
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &num)
		if num <= 0 {
			negative = append(negative, num)
		} else if num > 1 {
			positive = append(positive, num)
		} else {
			sum += 1
		}
	}
	sort.Ints(positive)
	sort.Ints(negative)
	if len(positive) > 0 {
		if len(positive)%2 == 0 {
			for i := 0; i < len(positive); i += 2 {
				sum += positive[i] * positive[i+1]
			}
		} else {
			sum += positive[0]
			for i := 1; i < len(positive); i += 2 {
				sum += positive[i] * positive[i+1]
			}
		}
	}
	if len(negative) > 0 {
		if len(negative)%2 == 0 {
			for i := 0; i < len(negative); i += 2 {
				sum += negative[i] * negative[i+1]
			}
		} else {
			sum += negative[len(negative)-1]
			for i := 0; i < len(negative)-1; i += 2 {
				sum += negative[i] * negative[i+1]
			}
		}
	}
	fmt.Println(sum)
}
