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
	checkVal := make(map[int]bool)
	arr := []int{}
	for i := 0; i < arrLength; i++ {
		var num int
		fmt.Fscanf(reader, "%d ", &num)
		if _, exists := checkVal[num]; !exists {
			checkVal[num] = true
			arr = append(arr, num)

		}
	}

	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {

		fmt.Fprintf(writer, "%d ", arr[i])
	}
}
