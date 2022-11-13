package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	result := 1
	for i := 0; i < 3; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		result = result * num
	}
	numToCnt := map[string]int{
		"0": 0,
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
	}
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	strArr := strings.Split(fmt.Sprintf("%d", result), "")
	for _, num := range strArr {
		numToCnt[num]++
	}
	for _, num := range nums {
		fmt.Fprintf(writer, "%d\n", numToCnt[fmt.Sprintf("%d", num)])
	}
	writer.Flush()

}
