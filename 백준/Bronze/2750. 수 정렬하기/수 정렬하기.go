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
	defer writer.Flush()
	coordinates := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &coordinates[i])
	}

	sort.Slice(coordinates, func(i, j int) bool {

		return coordinates[j] > coordinates[i]
	})
	for _, num := range coordinates {
		fmt.Fprintf(writer, "%d\n", num)
	}
}
