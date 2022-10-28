package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader       = bufio.NewReader(os.Stdin)
	writer       = bufio.NewWriter(os.Stdout)
	visited      []bool
	graph        [][]int
	reverseGraph [][]int
)

type Stack struct {
	Node []int
}

func (s *Stack) Push(n int) {
	s.Node = append(s.Node, n)
}
func (s *Stack) Pop() (ret int) {
	if len(s.Node) <= 0 {
		return -1
	} else {
		ret = s.Node[len(s.Node)-1]
		if len(s.Node) == 1 {
			s.Node = []int{}
		} else {
			s.Node = s.Node[:len(s.Node)-1]
		}
	}
	return
}

func (s *Stack) Size() int {
	return len(s.Node)
}
func dfs(start int, stack *Stack) {
	visited[start] = true
	for _, now := range graph[start] {
		if visited[now] == false {
			dfs(now, stack)
		}
	}
	stack.Push(start)
}
func reverseDfs(start int, stack *Stack) {
	visited[start] = true
	stack.Push(start)
	for _, now := range reverseGraph[start] {
		if visited[now] == false {
			reverseDfs(now, stack)
		}
	}
}
func main() {
	defer writer.Flush()
	var v, e, s, f, now int
	fmt.Fscanf(reader, "%d %d\n", &v, &e)
	graph = make([][]int, v+1)
	reverseGraph = make([][]int, v+1)
	for i := 0; i < e; i++ {
		fmt.Fscanf(reader, "%d %d\n", &s, &f)
		graph[s] = append(graph[s], f)
		reverseGraph[f] = append(reverseGraph[f], s)
	}

	var stack = new(Stack)
	visited = make([]bool, v+1)
	for i := 1; i <= v; i++ {
		if visited[i] == false {
			dfs(i, stack)
		}
	}
	visited = make([]bool, v+1)
	var res = [][]int{}
	for stack.Size() > 0 {
		var scc = new(Stack)
		now = stack.Pop()
		if visited[now] == false {
			reverseDfs(now, scc)
			sort.Ints(scc.Node)
			res = append(res, scc.Node)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	fmt.Fprintf(writer, "%d\n", len(res))
	for _, v := range res {
		for _, k := range v {
			fmt.Fprintf(writer, "%d ", k)
		}
		fmt.Fprintf(writer, "-1 \n")
	}
}
