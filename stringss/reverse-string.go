package stringss

/*
LEETCODE
344. Reverse String
*/

func ReverseString(str []byte) []byte {

	startPtr := 0
	endPtr := len(str) - 1

	for startPtr < endPtr {
		str[startPtr], str[endPtr] = str[endPtr], str[startPtr]
		startPtr++
		endPtr--
	}

	return str
}
