package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack struct {
	Node []string
}

func (s *Stack) Push(n string) {
	s.Node = append(s.Node, n)
}
func (s *Stack) Pop() (ret string) {
	if len(s.Node) <= 0 {
		return ""
	} else {
		ret = s.Node[len(s.Node)-1]
		if len(s.Node) == 1 {
			s.Node = []string{}
		} else {
			s.Node = s.Node[:len(s.Node)-1]
		}
	}
	return ret
}

func (s *Stack) Last() (ret string) {
	if len(s.Node) <= 0 {
		return ""
	}
	ret = s.Node[len(s.Node)-1]
	return ret
}

func (s *Stack) Size() int {

	return len(s.Node)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var str string
	var tcCnt = 1

	for true {

		fmt.Fscanf(reader, "%s\n", &str)
		if strings.Contains(str, "-") {
			break
		}
		var stack = new(Stack)
		strArr := strings.Split(str, "")
		var changeCnt int
		for _, v := range strArr {
			if v == "{" {
				stack.Push(v)
			} else if v == "}" {
				open := stack.Last()
				if open == "{" {
					stack.Pop()
				} else {
					stack.Push(v)
				}
			}
		}
		for stack.Size() > 0 {
			if stack.Pop() == "}" {
				if stack.Last() == "}" {
					changeCnt++
				} else {
					changeCnt += 2
				}
				stack.Pop()
			} else {
				if stack.Last() == "{" {
					changeCnt++
				} else {
					changeCnt += 2
				}
				stack.Pop()
			}

		}
		fmt.Fprintf(writer, "%d. %d\n", tcCnt, changeCnt)
		tcCnt++
	}
}
