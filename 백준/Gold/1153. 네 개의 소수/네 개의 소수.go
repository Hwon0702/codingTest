package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var n, j int
	defer writer.Flush()
	fmt.Fscanf(reader, "%d\n", &n)
	if n < 8 {
		fmt.Fprintf(writer, "%d\n", -1)
	} else {
		var nums = make([]bool, n+1)
		for i := 2; i < n+1; i++ {
			nums[i] = true
		}
		for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
			if nums[i] {
				j = 2
				for i*j < n {
					nums[i*j] = false
					j++
				}
			}
		}
		if n%2 == 0 {
			var ans = []int{2, 2}
			n -= 4
			for i := 2; i < n+1; i++ {
				if nums[i] && nums[n-i] {
					ans = append(ans, i)
					ans = append(ans, n-i)
					break
				}
			}
			for _, v := range ans {
				fmt.Fprintf(writer, "%d ", v)
			}
		} else {
			var ans = []int{2, 3}
			n -= 5
			for i := 2; i < n+1; i++ {
				if nums[i] && nums[n-i] {
					ans = append(ans, i)
					ans = append(ans, n-i)
					break
				}
			}
			for _, v := range ans {
				fmt.Fprintf(writer, "%d ", v)
			}
		}
	}
}
