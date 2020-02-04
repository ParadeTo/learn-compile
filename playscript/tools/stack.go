package tools

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	len := len(*s)
	if len > 0 {
		item := (*s)[len-1]
		*s = (*s)[:len-1]
		return item
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	len := len(*s)
	if len > 0 {
		item := (*s)[len-1]
		return item
	}
	return nil
}

func NewStack() Stack {
	return make(Stack, 0)
}
