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
	var n int
	defer writer.Flush()
	fmt.Fscanf(reader, "%d\n", &n)

	var sum int
	for i := 0; i < n; i++ {
		var word string
		defer writer.Flush()
		fmt.Fscanf(reader, "%s\n", &word)
		var prev string
		wordToCnt := make(map[string]bool)
		wordArr := strings.Split(word, "")
		add := true
		for j := 0; j < len(wordArr); j++ {

			if prev != wordArr[j] {
				prev = wordArr[j]
				if !wordToCnt[wordArr[j]] {
					wordToCnt[wordArr[j]] = true
				} else {
					add = false
					break

				}
			}
		}
		if add {
			sum++
		}
	}
	fmt.Fprintln(writer, sum)
}
