package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Point []int
	Time  []int
}

func (s *Stack) Push(p, t int) {
	s.Point = append(s.Point, p)
	s.Time = append(s.Time, t)
}

func (s *Stack) Pop() (p, t int) {
	if len(s.Point) <= 0 {
		p, t = -1, -1
	} else {
		p = s.Point[len(s.Point)-1]
		t = s.Time[len(s.Time)-1]
		if len(s.Point) == 1 {
			s.Point = []int{}
			s.Time = []int{}
		} else {
			s.Point = s.Point[:len(s.Point)-1]
			s.Time = s.Time[:len(s.Time)-1]
		}
	}
	return p, t
}

func (s *Stack) Size() int {
	return len(s.Point)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc, getPoint int
	fmt.Fscanf(reader, "%d\n", &tc)
	var stack = new(Stack)
	for i := 0; i < tc; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strArr := strings.Split(str, " ")
		if len(strArr) > 1 {
			point, _ := strconv.Atoi(strArr[1])
			time, _ := strconv.Atoi(strArr[2])
			stack.Push(point, time)
		}
		if stack.Size() > 0 {
			point, time := stack.Pop()
			if time-1 == 0 {
				getPoint += point
			} else {
				stack.Push(point, time-1)
			}
		} else {
			continue
		}
	}
	fmt.Println(getPoint)

}
