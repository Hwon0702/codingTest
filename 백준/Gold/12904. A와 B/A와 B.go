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

func main() {
	defer writer.Flush()
	var start, target []string
	var s, t string
	var flag int
	flag = 0
	fmt.Fscanf(reader, "%s\n%s\n", &s, &t)
	start = strings.Split(s, "")
	target = strings.Split(t, "")
	for len(target) > 0 {
		if target[len(target)-1] == "A" {
			target = target[:len(target)-1]
		} else if target[len(target)-1] == "B" {
			target = target[:len(target)-1]
			for i := len(target)/2 - 1; i >= 0; i-- {
				tmp := len(target) - 1 - i
				target[i], target[tmp] = target[tmp], target[i]
			}

		}
		if strings.Join(start, "") == strings.Join(target, "") {
			flag = 1
			break
		}
	}
	fmt.Fprintf(writer, "%d\n", flag)
}
