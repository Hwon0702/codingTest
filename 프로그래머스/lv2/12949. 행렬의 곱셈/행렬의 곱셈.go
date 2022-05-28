func calcArr(val, colnum int, arr2 [][]int) []int {
	ret := []int{}
	for i := 0; i < len(arr2[0]); i++ {
		ret = append(ret, val*arr2[colnum][i])
		//fmt.Println("val : ", val, "arr2 : ", arr2[colnum][i])
	}
	//	for i := 0; i < len(arr2); i++ {
	//		for j := 0; j < len(arr2[0]); j++ {
	//			fmt.Println("value:", val, "arr2: ", arr2[j][i])
	//		}
	//	}
	return ret
}

func calc(appendArr, tmp []int) []int {
	if len(appendArr) != len(tmp) {
		return tmp
	}
	for idx, _ := range appendArr {
		appendArr[idx] += tmp[idx]
	}
	return appendArr

}

func solution(arr1 [][]int, arr2 [][]int) [][]int {
	ret := [][]int{}
	for _, arr := range arr1 {
		appendArr := []int{}
		for colnum, val := range arr {
			tmp := calcArr(val, colnum, arr2)
			appendArr = calc(appendArr, tmp)
		}
		ret = append(ret, appendArr)
	}
	return ret
}