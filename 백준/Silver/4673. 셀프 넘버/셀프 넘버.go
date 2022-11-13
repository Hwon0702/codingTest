package main

import (
	"fmt"
)

func main() {
	result := S4673(10000)
	for i := 1; i < len(result); i++ {
		if result[i] == false {
			fmt.Println(i)
		}
	}
}

func S4673(n int) map[int]bool {
	result := make(map[int]bool, n+1)
	for i := 0; i < n+1; i++ {
		result[i] = false
	}
	for i := 0; i < n+1; i++ {
		var sum = i
		var number = i
		for j := number; j != 0; j /= 10 {
			sum += j % 10
		}
		if sum <= n {
			result[sum] = true
		}
	}

	return result
}
