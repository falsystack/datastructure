package mymap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashMap(t *testing.T) {
	var h HashMap[int]

	h.Add("tucker", 100)
	v, ok := h.Get("tucker")
	assert.True(t, ok)
	assert.Equal(t, 100, v)

	h.Add("golang", 200)
	v, ok = h.Get("golang")
	assert.True(t, ok)
	assert.Equal(t, 200, v)

	v, ok = h.Get("tucker")
	assert.True(t, ok)
	assert.Equal(t, 100, v)

	h.Add("awesome", 300)
	v, ok = h.Get("awesome")
	assert.True(t, ok)
	assert.Equal(t, 300, v)
}

func TestGoBasicMap(t *testing.T) {
	m := make(map[string]int)
	m["tucker"] = 100
	m["golang"] = 200
	m["awesome"] = 300

	assert.Equal(t, 100, m["tucker"])
	assert.Equal(t, 200, m["golang"])
	assert.Equal(t, 300, m["awesome"])
	assert.Equal(t, 0, m["aaa"]) // ない値を照会すると基本値が出る
	_, ok := m["aaa"]            // 値があるかどうかの確認方法
	assert.False(t, ok)

	delete(m, "tucker")
	v, ok := m["tucker"]
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}
