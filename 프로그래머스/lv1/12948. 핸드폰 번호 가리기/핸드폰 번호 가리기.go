import "strings"
func solution(phone_number string) string {
	ret := ""
	arr := strings.Split(phone_number, "")
	for i, val := range arr {
		if i < len(arr)-4 {
			ret += "*"
		} else {
			ret += val
		}
	}
	return ret
}