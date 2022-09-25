package main

import (
	"bufio"
	"fmt"
	"os"
)

func palindrome(target []int) int {
	for i := 0; i < int(len(target)/2); i++ {
		if target[i] != target[len(target)-i-1] {
			return 0
		}
	}
	return 1

}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, tc, s, e int
	fmt.Fscanf(reader, "%d\n", &n)
	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d %d\n", &s, &e)
		fmt.Fprintf(writer, "%d\n", palindrome(numbers[s-1:e]))
	}
}
