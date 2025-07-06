package stringss

import (
	"strconv"
)

func RlEncoding(str string) string {

	if len(str) == 0 {
		return ""
	}

	result := ""
	count := 1
	curCh := str[0]

	for i := 1; i < len(str); i++ {
		if str[i] == curCh {
			count++

		} else {
			result += string(curCh) + strconv.Itoa(count)
			curCh = str[i]
			count = 1

		}

	}
	result += string(curCh) + strconv.Itoa(count)
	return result
}
