package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	min  = math.MaxFloat64
	max  = math.MaxFloat64 * (-1)
	num  int
	nums []float64
)

func dfs(depth int, res float64, plus, minus, multiply, divide int) {
	if depth == num {
		min = math.Min(res, min)
		max = math.Max(res, max)
		return
	}
	if plus > 0 {
		dfs(depth+1, res+nums[depth], plus-1, minus, multiply, divide)
	}
	if minus > 0 {
		dfs(depth+1, res-nums[depth], plus, minus-1, multiply, divide)
	}
	if multiply > 0 {
		dfs(depth+1, res*nums[depth], plus, minus, multiply-1, divide)
	}
	if divide > 0 {
		dfs(depth+1, float64(int(res/nums[depth])), plus, minus, multiply, divide-1)
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanf(reader, "%d\n", &num)
	nums = make([]float64, num)
	var methods = make([]int, 4)
	for i := 0; i < num; i++ {
		fmt.Fscanf(reader, "%f ", &nums[i])
	}
	for i := 0; i < 4; i++ {
		fmt.Fscanf(reader, "%d ", &methods[i])
	}
	dfs(1, nums[0], methods[0], methods[1], methods[2], methods[3])
	fmt.Fprintf(writer, "%0.f\n%0.f", max, min)
}
