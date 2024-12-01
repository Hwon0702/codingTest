package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var countries = make([][]int, n)
	for i := 0; i < n; i++ {
		countries[i] = make([]int, 5)
		fmt.Fscanf(reader, "%d %d %d %d\n", &countries[i][0], &countries[i][1], &countries[i][2], &countries[i][3])
	}
	sort.Slice(countries, func(i, j int) bool {
		if countries[i][1] > countries[j][1] {
			return true
		} else if countries[i][1] == countries[j][1] {
			return countries[i][2] > countries[j][2]
		} else if countries[i][2] == countries[j][2] {
			return countries[i][1]+countries[i][2]+countries[i][3] > countries[j][1]+countries[j][2]+countries[j][3]
		}
		return false
	})
	var rank = 1
	//금으로 순위
	for i := 0; i < n-1; i++ {
		if countries[i][1] > countries[i+1][1] {
			countries[i][4] = rank
			countries[i+1][4] = rank + 1
			rank++
			continue
		} else if countries[i][1] == countries[i+1][1] {
			countries[i+1][4] = countries[i][4]
			rank++
			continue
		}
	}
	//은으로 순위
	for i := 0; i < n-1; i++ {
		var rank = countries[i][4]
		if countries[i][1] > countries[i+1][1] {
			continue
		}
		if countries[i][2] > countries[i+1][2] {
			countries[i][4] = rank
			countries[i+1][4] = rank + 1
			continue
		} else if countries[i][2] == countries[i+1][2] {
			countries[i+1][4] = countries[i][4]
			continue
		}
	}
	for i := 0; i < n-1; i++ {
		var rank = countries[i][4]
		if countries[i][1] >= countries[i+1][1] || countries[i][2] >= countries[i+1][2] {
			continue
		}
		if countries[i][3] > countries[i+1][3] {
			countries[i][4] = rank
			countries[i+1][4] = rank + 1
			continue
		} else if countries[i][3] == countries[i+1][3] {
			sum1 := countries[i][1] + countries[i][2] + countries[i][3]
			sum2 := countries[i+1][1] + countries[i+1][2] + countries[i+1][3]
			if sum1 > sum2 {
				countries[i+1][4] = countries[i][4] + 1
			} else {
				countries[i+1][4] = countries[i][4]
			}
		}
	}

	for i := 0; i < n; i++ {
		if countries[i][0] == m {
			fmt.Fprintf(writer, "%d\n", countries[i][4])
		}
	}
}
