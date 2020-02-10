package tools

// 空结构体
var Exists = struct{}{}

type Set struct {
	// struct为结构体类型的变量
	m map[interface{}]struct{}
}

func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(s.m, item)
	}
}

func (s *Set) RemoveSubset(subset *Set) {
	for item := range subset.m {
		s.Remove(item)
	}
}

func (s *Set) ForEach(cb func(interface{})) {
	for item := range s.m {
		cb(item)
	}
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

func NewSet(items ...interface{}) *Set {
	// 获取Set的地址
	s := &Set{}
	// 声明map类型的数据结构
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}
