package stringss

/*
LEETCODE
771. Jewels and Stones

*/

func NumJewelsInStones(jewels string, stones string) int {

	jewelSets := make(map[rune]bool)

	for _, j := range jewels {
		jewelSets[j] = true
	}

	count := 0

	for _, i := range stones {
		if jewelSets[i] {
			count++
		}
	}
	return count
}
