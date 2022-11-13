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
	var n int                         //삽입될 숫자
	var sum float64                   //합계
	var avg float64                   //평균
	var maxVal float64                //최빈값
	var numRange int                  //최댓값 - 최솟값
	numToCnt := make(map[float64]int) //최빈값 찾기 위함
	fmt.Fscanln(reader, &n)
	numArr := make([]float64, n)
	for i := 0; i < n; i++ {
		var num float64
		fmt.Fscanln(reader, &num)
		sum += num
		numToCnt[num]++
		numArr[i] = num
	}
	avg = sum / float64(n)

	sort.Float64s(numArr)
	numRange = int(numArr[n-1] - numArr[0])

	nums := make([]float64, 0, len(numToCnt))
	cnts := make([]int, 0, len(numToCnt))
	var maxCnt int
	for num, cnt := range numToCnt {
		if cnt > maxCnt {
			maxVal = num
			maxCnt = cnt
			nums = []float64{num}
			cnts = []int{cnt}
		} else if cnt == maxCnt {
			nums = append(nums, num)
			cnts = append(cnts, cnt)
		}
	}

	sort.Float64s(nums)
	if len(nums) > 1 {
		maxVal = nums[1]
	}
	if avg > -0.5 && avg < 0.5 {
		avg = 0
	}

	fmt.Fprintf(writer, "%.0f\n", math.Round(avg))
	fmt.Fprintln(writer, numArr[n/2])
	fmt.Fprintln(writer, maxVal)
	fmt.Fprintln(writer, numRange)

	writer.Flush()
}
