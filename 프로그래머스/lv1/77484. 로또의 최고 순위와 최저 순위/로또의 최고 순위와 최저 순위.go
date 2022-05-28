func solution(lottos []int, win_nums []int) []int {
	var countWin int
	var zeroNum int
	for i := 0; i < len(lottos); i++ {
		if lottos[i] == 0 {
			zeroNum++
		}
		for j := 0; j < len(win_nums); j++ {
			if lottos[j] == win_nums[i] {
				countWin++
			}

		}

	}
	maxVal := countWin + zeroNum
	score := map[int]int{
		0: 6,
		1: 6,
		2: 5,
		3: 4,
		4: 3,
		5: 2,
		6: 1,
	}
	ret := []int{}
	ret = append(ret, score[maxVal])
	ret = append(ret, score[countWin])

	return ret
}