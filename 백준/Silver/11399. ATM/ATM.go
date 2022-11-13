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
	fmt.Fscanf(reader, "%d\n", &n)
	var atm = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &atm[i])
	}
	sort.Ints(atm)
	var waiting int
	var totalWaiting int
	for i := 0; i < n; i++ {
		waiting = waiting + atm[i]
		totalWaiting += waiting
	}
	fmt.Fprintf(writer, "%d\n", totalWaiting)
}
