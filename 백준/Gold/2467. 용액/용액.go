package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	var liquid = make([]int, tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d ", &liquid[i])
	}

	start := 0
	end := tc - 1
	min := liquid[start] + liquid[end]
	minStart := start
	minEnd := end

	sort.Ints(liquid)
	for start < end {
		if min == 0 {
			break
		}
		//0에 근접 = 0-절댓값(SUM)이 더 크다
		if 0-math.Abs(float64(min)) >= 0-math.Abs(float64(liquid[start]+liquid[end])) { //기존게 더 0에 가까움
			//start나 end에 가공을 한다.
			if (liquid[start] + liquid[end]) < 0 {
				start++
			} else {
				end--
			}

		} else { //바뀐게 0에 더 가까움
			minStart = start
			minEnd = end
			min = liquid[start] + liquid[end]

		}
	}
	fmt.Println(liquid[minStart], liquid[minEnd])

}
