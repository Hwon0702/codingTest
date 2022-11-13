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
	var length int
	fmt.Fscanln(reader, &length)
	var num int
	intArr := make([]int, length)
	i := 0
	for i < length {
		fmt.Fscan(reader, &num)
		intArr[i] = num
		i++
	}
	sort.Ints(intArr)
	fmt.Fprintf(writer, "%d %d \n", intArr[0], intArr[length-1])
	writer.Flush()
}
