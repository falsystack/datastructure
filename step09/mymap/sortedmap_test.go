package mymap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedMap_Add(t *testing.T) {
	var s SortedMap[string, int]

	s.Add("aaa", 10)
	v, _ := s.Get("aaa")
	assert.Equal(t, 10, v)

	s.Add("bbb", 20)
	v, _ = s.Get("bbb")
	assert.Equal(t, 20, v)

}
