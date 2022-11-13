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

func (s *Stack) Size() int {
	return len(s.Node)
}

func (s *Stack) Push(w string) {
	s.Node = append(s.Node, w)
}

func (s *Stack) Pop() (ret string) {
	if s.Size() <= 0 {
		return ""
	} else {
		ret = s.Node[s.Size()-1]
		if s.Size() == 1 {
			s.Node = []string{}
		} else {
			s.Node = s.Node[:s.Size()-1]
		}
	}
	return ret
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var str string
	str, _ = reader.ReadString('\n')
	strings.TrimSpace(str)
	stringsArr := strings.Split(str, "")
	retArr := []string{}
	var revFlag = false
	var stack = new(Stack)
	for _, v := range stringsArr {
		if v == "<" || revFlag && v != ">" {
			for stack.Size() > 0 {
				retArr = append(retArr, stack.Pop())
			}
			revFlag = true
			retArr = append(retArr, v)
		} else if v == ">" {
			revFlag = false
			retArr = append(retArr, v)

		} else if v == " " {
			for stack.Size() > 0 {
				retArr = append(retArr, stack.Pop())
			}
			retArr = append(retArr, " ")
		} else {
			if v != "\n" {
				stack.Push(v)
			}

		}
	}
	if stack.Size() > 0 {
		for stack.Size() > 0 {
			retArr = append(retArr, stack.Pop())
		}
	}
	fmt.Fprintf(writer, "%s\n", strings.Join(retArr, ""))
}
