package stack

func ReverseStack(input []int) []int {

	st := []int{}
	rev := []int{}

	for _, val := range input {
		st = append(st, val)

	}

	for len(st) > 0 {
		top := st[len(st)-1] //peek
		st = st[:len(st)-1]
		rev = append(rev, top)
	}

	return rev
}
