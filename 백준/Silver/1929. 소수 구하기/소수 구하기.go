package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var m, n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &m, &n)
	defer writer.Flush()

	var primeNumber = make(map[int]bool, n-m+1)
	for i := m; i < n+1; i++ {
		primeNumber[i] = true
	}
	if m == 1 {
		primeNumber[m] = false
	}

	for i := m; i < n+1; i++ {
		var divisorCount int
		if primeNumber[i] {
			for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					divisorCount++
				}
				if divisorCount > 1 {
					primeNumber[i] = false
					// 에라토스테네스의 체
					for k := 1; k*i < n+1; k++ {
						primeNumber[k*i] = false
					}
					break
				}
			}
		}
	}

	var keys []int
	for k, v := range primeNumber {
		if v {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Fprintf(writer, "%d\n", k)
	}
}
