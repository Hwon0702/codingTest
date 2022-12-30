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
	flag   int
	start  []string
)

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func dfs(target []string) {
	if len(start) == len(target) {
		if strings.Join(start, "") == strings.Join(target, "") {
			flag = 1
		}
		return
	}
	if target[len(target)-1] == "A" {
		target = target[:len(target)-1]
		dfs(target)
		target = append(target, "A")
	}
	if target[0] == "B" {
		target = reverse(target)
		target = target[:len(target)-1]
		dfs(target)
		target = append(target, "B")
		target = reverse(target)
	}

}

func main() {
	defer writer.Flush()
	var s, t string
	fmt.Fscanf(reader, "%s\n%s\n", &s, &t)
	target := strings.Split(t, "")
	start = strings.Split(s, "")
	dfs(target)
	fmt.Fprintf(writer, "%d\n", flag)
}
