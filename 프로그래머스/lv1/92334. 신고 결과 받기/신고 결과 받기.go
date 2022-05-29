import(
    "strings"
)

func makeSliceUnique(s []string) []string {
	keys := make(map[string]struct{})
	res := make([]string, 0)
	for _, val := range s {
		if _, ok := keys[val]; ok {
			continue
		} else {
			keys[val] = struct{}{}
			res = append(res, val)
		}
	}
	return res
}

func findIdx(user string, idList []string) int {
	for k, v := range idList {
		if user == v {
			return k
		}
	}
	return -1
}
func solution(id_list []string, report []string, k int) []int {
	
	idToReport := map[string]int{}
	userToReport := map[string][]string{}
	userToMailing := map[string]int{}
	userMailing := []int{}

	for _, id := range id_list {
		userToMailing[id] = 0
		userMailing = append(userMailing, 0)
	}
	
	blockedUser := map[string]bool{}
	for _, id := range report {
		reports := strings.Split(id, " ")
		if _, ok := userToReport[reports[0]]; ok {
			userToReport[reports[0]] = append(userToReport[reports[0]], reports[1])
		} else {
			userToReport[reports[0]] = append(userToReport[reports[0]], reports[1])
		}

	}
	for idx, val := range userToReport {
		userToReport[idx] = makeSliceUnique(val)
		for _, report := range userToReport[idx] {

			if _, ok := idToReport[report]; ok {
				idToReport[report] += 1
			} else {
				idToReport[report] = 1

			}

			if idToReport[report] >= k {
				blockedUser[report] = true
			}
		}
	}
	for user, targets := range userToReport {
		for _, reportTarget := range targets {
			if _, ok := blockedUser[reportTarget]; ok {
				userToMailing[user]++
			}
		}
	}

	for user, val := range userToMailing {
		idx := findIdx(user, id_list)
		userMailing[idx] = val
	}

	return userMailing
}