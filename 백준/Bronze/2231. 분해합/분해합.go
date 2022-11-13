package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	var creator int
	fmt.Fscanf(reader, "%d\n", &n, &n)
	for i := 0; i < n; i++ {
		if getCreator(i) == n {
			creator = i
			break
		}
	}
	fmt.Println(creator)
}

func getCreator(num int) int {
	var sum = num
	for num != 0 {
		sum += num % 10
		num /= 10

	}
	return sum

}
