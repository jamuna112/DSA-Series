package search

func IndexOfAllOccurences(input []int, target int) []int {
	numOfInd := make([]int, 0)

	for i := 0; i < len(input); i++ {
		if input[i] == target {
			numOfInd = append(numOfInd, i)
		}
	}

	return numOfInd
}
