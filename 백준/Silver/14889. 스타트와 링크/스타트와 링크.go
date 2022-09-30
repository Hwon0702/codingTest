package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var resultCombination = [][]int{}

func nCr(n int, nums, ans []int, r int) {
	if n == len(nums) {
		if len(ans) == r {
			var tmp = []int{}
			for _, v := range ans {
				tmp = append(tmp, v)
			}
			resultCombination = append(resultCombination, tmp)
		}
		return
	}
	ans = append(ans, nums[n])
	nCr(n+1, nums, ans, r)
	ans = ans[:len(ans)-1]
	nCr(n+1, nums, ans, r)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	var res, teamA, teamB float64
	res = math.MaxFloat64
	fmt.Fscanf(reader, "%d\n", &n)

	var p = make([]int, n)
	var persons = make([][]float64, n)
	for i := 0; i < n; i++ {
		persons[i] = make([]float64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%f ", &persons[i][j])
		}
		p[i] = i
	}

	nCr(0, p, []int{}, int(n/2))
	for i := 0; i < int(len(resultCombination)/2); i++ {
		team := resultCombination[i]
		teamA = 0
		for j := 0; j < int(n/2); j++ {
			mem := team[j] //팀A에 속한 멤버들
			for _, k := range team {
				teamA += persons[mem][k] //팀A에 속한 멤버들끼리의 합
			}
		}
		team = resultCombination[len(resultCombination)-1-i]
		teamB = 0
		for j := 0; j < int(n/2); j++ {
			mem := team[j] //팀B에 속한 멤버들
			for _, k := range team {
				teamB += persons[mem][k] //팀B에 속한 멤버들끼리의 합
			}
		}
		res = math.Min(res, math.Abs(teamA-teamB))
		if res == 0 {
			break
		}
	}
	fmt.Fprintf(writer, "%0.f\n", res)
}
