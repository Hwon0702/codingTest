import(
"strings"
"fmt")
func solution(s string) string {
   
	arr := strings.Split(s, "")
	if len(s)%2 == 0 {
		return fmt.Sprintf("%s", arr[len(s)/2-1]+arr[len(s)/2])
	}
	return fmt.Sprintf("%s", arr[len(s)/2])
}