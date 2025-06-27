package stringss

/*
LEETCODE
1662. Check If Two String Arrays are Equivalent
*/

func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	return sum(word1) == sum(word2)

}

func sum(words []string) string {

	var add string

	for i := 0; i < len(words); i++ {
		add += words[i]
	}

	return add
}
