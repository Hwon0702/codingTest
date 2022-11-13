package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findStr(c string, n int, str []string) (count int) {
	count = 1
	var drawing = make([]bool, n)
	var index = -1
	for i, v := range str {
		if c == v {
			drawing[i] = true
		}
	}
	for i, _ := range str {
		if drawing[i] == false {
			if index < 0 {
				count += 1
				index = i
			} else if index == i-1 {
				index = i
			} else {
				index = i
				count += 1
			}
		}
	}
	return count
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	var str string
	fmt.Fscanf(reader, "%d\n%s", &n, &str)
	strArr := strings.Split(str, "")
	RCount := findStr("R", n, strArr)
	BCount := findStr("B", n, strArr)
	if RCount < BCount {
		fmt.Fprintf(writer, "%d\n", RCount)
	} else {
		fmt.Fprintf(writer, "%d\n", BCount)
	}

}
