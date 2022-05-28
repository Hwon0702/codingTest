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
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	var request = make([]int, tc)
	var sum, budget int
	for i := 0; i < tc; i++ {
		var n int
		fmt.Fscanf(reader, "%d ", &n)
		request[i] = n
		sum += n
	}
	sort.Ints(request)
	fmt.Fscanf(reader, "%d\n", &budget)
	if sum <= budget {

		fmt.Fprintf(writer, "%d", request[tc-1]) //모든 지자체가 요구한 금액의 합이 예산보다 작을 경우
	} else { //모든 지자체가 요구한 금액의 합이 예산보다 클 경우
		sumMoney := 0
		min := 0
		max := request[tc-1]
		mid := sum / tc
		for min < max { //start가 end보다 작고 max(모든 지자체의 요구의 합)보다 예산이 크거나 같을 경우
			mid = (min + max) / 2

			sumMoney = 0
			for i := 0; i < tc; i++ {
				if request[i] <= mid {
					sumMoney += request[i]
				} else {
					sumMoney += max * (tc - i)
					break
				}
			}
			if sumMoney == budget {
				break
			} else if sumMoney < budget {
				min = max
				max += 2
			} else if sumMoney > budget {
				max = mid
			}

		}
		fmt.Fprintf(writer, "%d", max)
	}

}
