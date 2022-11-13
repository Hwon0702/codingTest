package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	var sum int
	var numbers string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &numbers)
	numArr := strings.Split(numbers, "")
	for i := 0; i < n; i++ {
		number, _ := strconv.Atoi(numArr[i])
		sum += number
	}
	fmt.Fprintf(writer, "%d\n", sum)
	writer.Flush()
}