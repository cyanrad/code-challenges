package leet

func minWinSubs(s string, t string) string {
	tMap := map[rune]int{}
	wMap := map[rune]int{}
	wMapCondi := 0
	temp := ""
	final := ""

	for _, c := range t {
		tMap[c] += 1
	}

	for _, c := range s {
		if wMapCondi < len(tMap) {
			if n, ok := tMap[c]; ok {
				wMap[c] += 1
				if n == wMap[c] {
					wMapCondi++
				}
			}
			temp += string(c)
		} else {

		}

	}
	return final
}
