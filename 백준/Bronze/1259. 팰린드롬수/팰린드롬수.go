package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func palindrome(s string) string {
	sArr := strings.Split(s, "")
	for i := 0; i < int(len(sArr)/2); i++ {
		if sArr[i] != sArr[len(sArr)-i-1] {
			return "no"
		}
	}
	return "yes"
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	for {
		fmt.Fscanf(reader, "%s\n", &s)
		if s == "0" {
			break
		}
		fmt.Println(palindrome(s))
	}
}
