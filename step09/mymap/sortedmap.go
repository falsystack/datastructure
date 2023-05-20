package mymap

import (
	"golang.org/x/exp/constraints"
	"sort"
)

type Element[TKey constraints.Ordered, TValue any] struct {
	Key   TKey
	Value TValue
}

type SortedMap[TKey constraints.Ordered, TValue any] struct {
	Arr []Element[TKey, TValue]
}

func (s *SortedMap[TKey, TValue]) Add(key TKey, value TValue) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr[idx].Value = value
		return
	}

	s.Arr = append(s.Arr[:idx], append([]Element[TKey, TValue]{
		{Key: key, Value: value},
	}, s.Arr[idx:]...)...)
}

func (s *SortedMap[TKey, TValue]) Get(key TKey) (TValue, bool) {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if s.Arr[idx].Key == key {
		return s.Arr[idx].Value, true
	}

	var defaultV TValue
	return defaultV, false
}

func (s *SortedMap[TKey, TValue]) Remove(key TKey) bool {
	idx := sort.Search(len(s.Arr), func(i int) bool {
		return s.Arr[i].Key >= key
	})

	if idx < len(s.Arr) && s.Arr[idx].Key == key {
		s.Arr = append(s.Arr[:idx], s.Arr[idx+1:]...)
		return true
	}
	return false
}
