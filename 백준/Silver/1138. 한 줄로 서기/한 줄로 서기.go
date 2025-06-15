package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	people := make([]int, n)
	alignPeople := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &people[i])
	}
	for i, v := range people {
		cnt := 0
		for j, res := range alignPeople {
			if res == 0 && cnt < v {
				cnt += 1
			} else if res == 0 && cnt == v {
				alignPeople[j] = i + 1
				break
			}
		}
	}
	for _, v := range alignPeople {
		fmt.Fprintf(writer, "%d ", v)
	}
}
