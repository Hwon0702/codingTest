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
	var str string
	fmt.Fscanln(reader, &str)
	defer writer.Flush()
	wordToMap := map[string]int{
		"A": 3,
		"B": 3,
		"C": 3,
		"D": 4,
		"E": 4,
		"F": 4,
		"G": 5,
		"H": 5,
		"I": 5,
		"J": 6,
		"K": 6,
		"L": 6,
		"M": 7,
		"N": 7,
		"O": 7,
		"P": 8,
		"Q": 8,
		"R": 8,
		"S": 8,
		"T": 9,
		"U": 9,
		"V": 9,
		"W": 10,
		"X": 10,
		"Y": 10,
		"Z": 10,
	}
	strArr := strings.Split(str, "")
	var sum int
	for i := 0; i < len(strArr); i++ {
		sum += wordToMap[strArr[i]]
	}
	fmt.Fprintln(writer, sum)
}
