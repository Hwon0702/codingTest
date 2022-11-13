package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	nodes []string
}

func (S *Stack) Push(s string) {
	S.nodes = append(S.nodes, s)
}

func (S *Stack) Pop() (ret string) {
	if len(S.nodes) <= 0 {
		return ""
	} else {
		ret = S.nodes[len(S.nodes)-1]
		if len(S.nodes) > 1 {
			S.nodes = S.nodes[:len(S.nodes)-1]
		} else {
			S.nodes = []string{}
		}
	}
	return ret
}

func (S *Stack) Size() int {
	return len(S.nodes)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var str string
	for true {
		str, _ = reader.ReadString('\n')
		if len(str) == 2 && strings.TrimSpace(str) == "." {
			break
		} else {
			str = strings.TrimSpace(str)
		}
		var stack = new(Stack)
		strArr := strings.Split(str, "")
		var result = "yes"
		for _, v := range strArr {
			if v == "[" || v == "(" {
				stack.Push(v)
			}
			if v == "]" {
				if stack.Pop() == "[" {
					continue
				} else {
					result = "no"
					break
				}
			} else if v == ")" {
				if stack.Pop() == "(" {
					continue
				} else {
					result = "no"
					break
				}
			}
		}
		if stack.Size() >= 1 {
			result = "no"
		}
		fmt.Fprintf(writer, "%s\n", result)
	}

}
