var result [][]int

func solution(n int)[][]int{
    ret :=Hanoi(n, 1, 3, 2)
    return ret
    
}

func Hanoi(n, from, to, aux int) [][]int {
	/*
		처음 들어오면
		n == 3(1번 기둥에 원판이 3개)
		from : 1(1번기둥)
		to : 3번으로
		aux : 2(보조)
			n : 남은 원판 갯수
			from : n번 기둥에서
			to : n 번 기둥으로
			aux : 보조 기둥
	*/
	if n == 1 {
		result = append(result, []int{from, to})

	} else {

		Hanoi(n-1, from, aux, to)
		result = append(result, []int{from, to})
		Hanoi(n-1, aux, to, from)
	}
	return result
}