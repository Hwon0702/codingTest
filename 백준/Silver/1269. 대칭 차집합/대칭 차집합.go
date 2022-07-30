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
	var lenA, lenB int
	fmt.Fscanf(reader, "%d %d\n", &lenA, &lenB)
	var arrA = make(map[int]bool, lenA)
	var arrB = make(map[int]bool, lenB)
	for i := 0; i < lenA; i++ {
		var n int
		fmt.Fscanf(reader, "%d ", &n)
		arrA[n] = false
	}
	for i := 0; i < lenB; i++ {
		var n int
		fmt.Fscanf(reader, "%d ", &n)
		arrB[n] = false
	}

	ABandBA := []int{}

	for k, _ := range arrA {
		if _, exists := arrB[k]; !exists {
			ABandBA = append(ABandBA, k)
		}
	}
	for k, _ := range arrB {
		if _, exists := arrA[k]; !exists {
			ABandBA = append(ABandBA, k)
		}
	}
	fmt.Fprintf(writer, "%d", len(ABandBA))
}
