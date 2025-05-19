package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	vowel  = []string{"a", "e", "i", "o", "u"}
)

func main() {
	defer writer.Flush()
	var s string
	var vowelCnt, consonantCnt, totalVowelCnt = 0, 0, 0
	var result = true
	for {
		vowelCnt, consonantCnt, totalVowelCnt = 0, 0, 0
		result = true
		fmt.Fscanf(reader, "%s\n", &s)

		if s == "end" {
			break
		}

		for i := 0; i < len(vowel); i++ {
			if strings.Contains(s, vowel[i]) {
				totalVowelCnt += 1
			}
		}
		if totalVowelCnt <= 0 {
			result = false
		}
		sArr := strings.Split(s, "")

		for i := 1; i < len(sArr); i++ {
			if (sArr[i-1] == sArr[i]) && (sArr[i] != "o" && sArr[i] != "e") {
				result = false
			}
		}
		for i := 0; i < len(sArr); i++ {
			if slices.Contains(vowel, sArr[i]) {
				vowelCnt += 1
				consonantCnt = 0
			} else {
				vowelCnt = 0
				consonantCnt += 1
			}
			if vowelCnt >= 3 || consonantCnt >= 3 {
				result = false
			}
		}
		if result {
			fmt.Fprintf(writer, "<%s> is acceptable.\n", s)
		} else {
			fmt.Fprintf(writer, "<%s> is not acceptable.\n", s)
		}
	}
}
