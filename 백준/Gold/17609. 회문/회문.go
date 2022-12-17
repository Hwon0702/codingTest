package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func SimilarPalindrome(target []string, s, e int) bool {
	for s < e {
		if target[s] == target[e] {
			s++
			e--
		} else {
			return false
		}
	}
	return true
}
func Palindrome(target []string, s, e int) int {
	for s < e {
		if target[s] == target[e] {
			s++
			e--
		} else {
			rmLeft := SimilarPalindrome(target, s+1, e)
			rmRight := SimilarPalindrome(target, s, e-1)
			if rmLeft || rmRight {
				return 1
			} else {
				return 2
			}
		}
	}
	return 0
}
func main() {
	defer writer.Flush()
	var n int
	var s string
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		sArr := strings.Split(strings.TrimSpace(s), "")
		fmt.Fprintf(writer, "%d\n", Palindrome(sArr, 0, len(sArr)-1))
	}
}
