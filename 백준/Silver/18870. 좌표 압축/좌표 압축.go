package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Unique(slice []int) []int {
	uniqMap := make(map[int]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	uniqSlice := make([]int, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanln(reader, &n)

	defer writer.Flush()
	coordinates := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &coordinates[i])
	}
	ret := make([]int, n)
	uniqueArr := Unique(coordinates)
	sort.Ints(uniqueArr)
	for i := 0; i < len(coordinates); i++ {
		idx := sort.SearchInts(uniqueArr, coordinates[i])
		ret[i] = idx
	}
	for _, num := range ret {
		fmt.Fprintf(writer, "%d ", num)
	}
}
