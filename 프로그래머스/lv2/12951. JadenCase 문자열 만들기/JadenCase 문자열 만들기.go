import (
	"strings"
)

func change(val string) string {
	ret := strings.ToLower(val)
	tmp := strings.Split(ret, "")
	tmp[0] = strings.ToUpper(tmp[0])
	return strings.Join(tmp, "")

}

func solution(s string) string{
   splits := strings.Split(s, " ")

	finish := []string{}
	for _, val := range splits {
		if len(val) > 0 {
			a := change(val)
			finish = append(finish, a)
		} else {
			finish = append(finish, "")
		}
	}

	return strings.Join(finish, " ")
}