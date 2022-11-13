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
	defer writer.Flush()
	var n int

	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		fmt.Println(getResult(str))
	}
}

func getResult(str string) string {
	var open, close int
	strArr := strings.Split(str, "")
	for _, p := range strArr {
		if p == "(" {
			open++
		} else if p == ")" {
			close++
		}
		if open < close {
			return "NO"
		}
	}
	if open != close {
		return "NO"
	} else {
		return "YES"
	}

}
