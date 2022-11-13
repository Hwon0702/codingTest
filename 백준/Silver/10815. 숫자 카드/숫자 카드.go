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
	var n int
	var sn int
	fmt.Fscanf(reader, "%d\n", &n)
	var searchTarget = make([]int, n)
	var number int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &searchTarget[i])
	}
	sort.Ints(searchTarget) //정렬
	fmt.Fscanf(reader, "%d\n", &sn)
	for i := 0; i < sn; i++ {

		fmt.Fscanf(reader, "%d ", &number)
		fmt.Fprintf(writer, "%d ", BinarySearch(searchTarget, number, 0, n-1))
	}
}

func BinarySearch(targetArray []int, target, start, end int) int {
	if start > end {
		return 0
	}
	mid := (start + end) / 2
	if targetArray[mid] == target {
		return 1
	} else if targetArray[mid] > target { //왼쪽
		return BinarySearch(targetArray, target, start, mid-1)
	} else if targetArray[mid] < target { //오른쪽
		return BinarySearch(targetArray, target, mid+1, end)
	}

	return 0
}
