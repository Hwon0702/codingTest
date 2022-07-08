package main

import (
	"bufio"
	"fmt"
	"os"
)

func mod(n int) int {
	return n / 2
}

func minus(n int) int {
	if n%10 == 1 {
		return n / 10
	} else {
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var start, result, count int
	count = 1
	fmt.Fscanf(reader, "%d %d\n", &start, &result)
	for result == start || result > 0 {
		count++
		if result%2 == 0 {
			result = mod(result)
		} else {
			result = minus(result)
		}
		if result == start || result == 0 {
			break
		}
	}
	if result == 0 {
		count = -1
	}
	fmt.Println(count)
}
