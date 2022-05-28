import "sort"

func findGCD(a, b int) (int, int) {
	var num1, num2 int
	module := a % b
	if module != 0 {
		num1, num2 = findGCD(b, module)

	} else {
		return a, b
	}
	return num1, num2

}

func solution(a []int) int{
	sort.Sort(sort.IntSlice(a))
	var gcd = 0
	var max = 0
	_, max = findGCD(a[0], a[1])
	gcd = a[0] * a[1] / max

	for idx := 1; idx < len(a); idx++ {
		_, max = findGCD(gcd, a[idx])
		gcd = gcd * a[idx] / max

	}
    return gcd
}