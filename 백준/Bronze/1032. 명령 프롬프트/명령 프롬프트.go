package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	first := []string{}

	for i := 0; i < t; i++ {

		var s string
		fmt.Fscanln(reader, &s)
		if i == 0 {
			first = strings.Split(s, "")
		} else {
			stringsArr := strings.Split(s, "")
			if len(first) != len(stringsArr) {
				diff := 0
				if len(first) < len(stringsArr) {
					diff = len(stringsArr) - len(first)
				} else {
					diff = len(first) - len(stringsArr)
				}
				for i := 0; i < diff; i++ {
					first = append(first, "?")
				}
			}
			for idx, word := range stringsArr {
				if first[idx] != word {
					first[idx] = "?"
				}

			}
		}
	}
	fmt.Println(strings.Join(first, ""))
}
