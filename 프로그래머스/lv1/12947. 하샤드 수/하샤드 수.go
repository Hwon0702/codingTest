import(
"strconv"
"strings")
func solution(x int) bool {
	strX := strconv.Itoa(x)
	arrX := strings.Split(strX, "")
	sumNums := 0
	for _, num := range arrX {
		number, _ := strconv.Atoi(num)
		sumNums += number

	}
	return (x%sumNums == 0)

}