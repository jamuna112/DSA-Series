package stack

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	//push element on top of the stack
	*s = append(*s, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (interface{}, bool) {

	if len(*s) == 0 {
		return nil, false
	}

	lastIndex := len(*s) - 1
	item := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return item, true
}

//Peak returns the top element without removing it

func (s *Stack) Peek() (interface{}, bool) {

	if len(*s) == 0 {
		return nil, false
	}

	lastIndex := len(*s) - 1
	item := (*s)[lastIndex]

	return item, true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
