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
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		var a, b int //k층 n호
		fmt.Fscanln(reader, &a)
		fmt.Fscanln(reader, &b)
		fmt.Fprintf(writer, "%d\n", getPersons(a, b))

	}
}

func getPersons(a, b int) (count int) {
	if a == 1 {
		for i := 1; i <= b; i++ {
			count += i
		}
		return count
	}
	for i := 1; i <= b; i++ {
		count += getPersons(a-1, i)
	}
	return count
}
