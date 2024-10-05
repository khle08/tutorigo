
package template

import _ "fmt"


var L = 0;


type Stack[T any] struct {
	Items []T
}


func (s *Stack[T]) GetItems() []T {
	return s.Items
}


func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}


func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Items) == 0 {
		var zeroValue T
		return zeroValue, false
	}

	item := s.Items[len(s.Items) - 1]
	s.Items = s.Items[:len(s.Items) - 1]
	return item, true
}


