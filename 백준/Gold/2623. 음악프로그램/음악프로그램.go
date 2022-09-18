package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	Node []int
}

func (q *Queue) Enqueue(n int) {
	q.Node = append(q.Node, n)
}

func (q *Queue) Dequeue() (ret int) {
	if len(q.Node) <= 0 {
		return -1
	} else {
		ret = q.Node[0]
		if len(q.Node) > 1 {
			q.Node = q.Node[1:]
		} else {
			q.Node = []int{}
		}
	}
	return ret
}

func (q *Queue) Len() int {
	return len(q.Node)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m, now int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var singers = make([][]int, n+1)
	var cnt = make([]int, n+1)
	var res = []int{}
	var q = new(Queue)
	for i := 0; i < n+1; i++ {
		singers[i] = []int{}
	}
	cnt[0] = -1
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &cnt)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strArr := strings.Split(str, " ")
		numArr := make([]int, len(strArr)-1)
		for j := 1; j < len(strArr); j++ {
			num, _ := strconv.Atoi(strArr[j])
			numArr[j-1] = num
		}
		for j := 0; j < len(numArr)-1; j++ {
			singers[numArr[j]] = append(singers[numArr[j]], numArr[j+1])
			cnt[numArr[j+1]] += 1
		}
	}
	for i := 1; i < n+1; i++ {
		if cnt[i] == 0 {
			q.Enqueue(i)
			res = append(res, i)
		}
	}
	for q.Len() > 0 {
		now = q.Dequeue()
		for _, next := range singers[now] {
			cnt[next]--
			if cnt[next] == 0 {
				q.Enqueue(next)
				res = append(res, next)
			}
		}
	}
	if len(res) == n {
		for _, v := range res {
			fmt.Println(v)
		}
	} else {
		fmt.Println(0)
	}
}
