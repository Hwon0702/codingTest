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
	var arrLength int
	fmt.Fscanf(reader, "%d\n", &arrLength)
	arrA := make([]int, arrLength)
	arrB := make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		fmt.Fscanf(reader, "%d ", &arrA[i])
	}
	for i := 0; i < arrLength; i++ {
		fmt.Fscanf(reader, "%d ", &arrB[i])
	}
	sort.Ints(arrA)
	sort.Slice(arrB, func(a, b int) bool {
		return arrB[a] > arrB[b]
	})
	var sum int
	for i := 0; i < arrLength; i++ {
		sum += arrA[i] * arrB[i]
	}
	fmt.Fprintf(writer, "%d\n", sum)
}
