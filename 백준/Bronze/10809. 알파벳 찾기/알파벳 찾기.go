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
	charToNum := make(map[string]int)
	chars := make([]string, 26)
	for i := 97; i < 123; i++ {
		charToNum[string(i)] = -1
		chars[i-97] = string(i)
	}
	strArr := strings.Split(str, "")
	for idx, char := range strArr {
		if charToNum[char] < 0 {
			charToNum[char] = idx
		}
	}
	for _, val := range chars {
		fmt.Fprint(writer, charToNum[val], " ")
	}
	fmt.Fprint(writer, "\n")
	writer.Flush()
}
