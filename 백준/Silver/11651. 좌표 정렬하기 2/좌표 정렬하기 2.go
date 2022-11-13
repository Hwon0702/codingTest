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
	var n int
	fmt.Fscanln(reader, &n)
	var coordinates = make([][]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		coordinates[i] = []int{a, b}
	}
	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i][1] == coordinates[j][1] {
			return coordinates[i][0] < coordinates[j][0]
		}
		return coordinates[i][1] < coordinates[j][1]
	})
	for _, nums := range coordinates {
		fmt.Fprintln(writer, nums[0], nums[1])
	}
	writer.Flush()
}
