package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	if n > 0 {
		var numbers = make([]int, n)

		var max = -math.MaxInt
		for i := 0; i < n; i++ {
			fmt.Fscanf(reader, "%d ", &numbers[i])
		}
		var sumArr = make([]int, len(numbers))
		for i, _ := range sumArr {
			sumArr[i] = 0
		}
		sumArr[0] = numbers[0]
		for i := 1; i < len(numbers); i++ {
			sumArr[i] = int(math.Max(float64(sumArr[i-1]+numbers[i]), float64(numbers[i])))
		}
		for _, v := range sumArr {
			if v > max {
				max = v
			}
		}
		fmt.Fprintf(writer, "%d", max)
	} else {
		fmt.Println(0)
	}
}
