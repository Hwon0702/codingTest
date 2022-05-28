package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
https://www.acmicpc.net/problem/2504
*/
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

/*
온전한 괄호가 아닐 때
(과 )의 갯수가 다를 때
[과 ]의 갯수가 다를 때
[다음)가 올 때. (다음]가 올 때
쌍으로 다 빼냈는데도 남은게 있을 때->괄호의 갯수가 홀수일 때

(가 나오면 -> 2곱함
[가 나오면 ->3곱함
[뒤에 바로 ]가 나오면 -> 3으로 처리함
(뒤에 바로)가 나오면 -> 2로 처리함

( 2 [ 3 ] ) ( 3 )
[3] =9

*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var str string
	var stack = new(Stack)
	fmt.Fscanf(reader, "%s\n", &str)
	strArr := strings.Split(str, "")
	var result = 0
	var piv = 1
	var valid string
	if len(strArr)%2 == 0 {
		strArr = strings.Split(str, "")
		for i, v := range strArr {
			if v == "[" {
				piv *= 3
				stack.Push("[")
			}
			if v == "(" {
				piv *= 2
				stack.Push("(")
			}
			if v == "]" {
				if i > 0 && strArr[i-1] == "[" {
					result += piv
				}
				piv /= 3
				valid = stack.Pop()
				if valid != "[" {
					result = 0
					break
				}

			}
			if v == ")" {
				if i > 0 && strArr[i-1] == "(" {
					result += piv
				}
				piv /= 2
				valid = stack.Pop()
				if valid != "(" {
					result = 0
					break
				}
			}
		}
		if len(stack.Node) > 0 {
			result = 0
		}
	}
	fmt.Fprintf(writer, "%d", result)
}
