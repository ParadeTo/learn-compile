package tools

type List []interface{}

func (l *List) Add(item interface{}) {
	*l = append(*l, item)
}

func (l *List) Remove(item interface{}) {
	index := -1
	for i, _item := range *l {
		if _item == item {
			index = i
			break
		}
	}
	if index > -1 {
		*l = append(
			(*l)[:index], (*l)[index+1:]...)
	}
}

func NewList() List {
	return make(List, 0)
}
