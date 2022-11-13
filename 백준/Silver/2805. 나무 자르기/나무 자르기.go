package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m) //항상 k<=n
	var trees = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &trees[i])
	}
	sort.Ints(trees)
	fmt.Fprintf(writer, "%d", getTree(m, trees))
}

func getTree(m int, tree []int) (max int) {
	start := 1
	end := tree[len(tree)-1]
	for start <= end {
		mid := (start + end) / 2
		var totalLength int
		for i := 0; i < len(tree); i++ {
			if tree[i]-mid >= 0 {
				totalLength += tree[i] - mid
			}
		}
		if totalLength >= m {
			if max < mid {
				max = mid
				start = mid + 1
			}
		} else {
			end = mid - 1
		}
	}
	return
}
