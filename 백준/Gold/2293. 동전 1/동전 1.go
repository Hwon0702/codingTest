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
	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	var numbers = make([]int, n)
	var results = make([]int, k+1)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &numbers[i])
	}
	sort.Ints(numbers)
	results[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			if j-numbers[i] >= 0 {
				results[j] += results[j-numbers[i]]
			}
		}
	}
	fmt.Println(results[k])
}
