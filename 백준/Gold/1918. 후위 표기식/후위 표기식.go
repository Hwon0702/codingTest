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
	var s, res string
	fmt.Fscanf(reader, "%s\n", &s)
	calculates := strings.Split(s, "")
	var stack = new(Stack)
	var alphabets = make([]string, 26)
	for i := 0; i < 26; i++ {
		alphabets[i] = string(i + 65)
	}

	for _, v := range calculates {
		if v != "(" && v != ")" && v != "+" && v != "-" && v != "*" && v != "/" {
			res += v
		} else if v == "(" {
			stack.Push(v)
		} else if v == "*" || v == "/" {
			for stack.Size() > 0 && (stack.Last() == "*" || stack.Last() == "/") {
				res += stack.Pop()
			}
			stack.Push(v)
		} else if v == "+" || v == "-" {
			for stack.Size() > 0 && stack.Last() != "(" {
				res += stack.Pop()
			}
			stack.Push(v)
		} else if v == ")" {
			for stack.Size() > 0 && stack.Last() != "(" {
				res += stack.Pop()
			}
			stack.Pop()
		}
	}
	for stack.Size() > 0 {
		res += stack.Pop()
	}
	fmt.Fprintf(writer, "%s\n", res)
}
