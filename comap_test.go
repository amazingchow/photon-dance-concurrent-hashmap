package comap

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func genFakeInput() (m map[string]interface{}) {
	m = make(map[string]interface{})
	for i := 0; i < 10000; i++ {
		m[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("val-%d", i)
	}
	return m
}

func TestCoMap(t *testing.T) {
	fakeInput := genFakeInput()
	cm := NewCoMap()
	cm.BatchStore(fakeInput)
	assert.Equal(t, cm.Empty(), false)
	assert.Equal(t, cm.Count(), len(fakeInput))
	assert.Equal(t, cm.StoreIfNotExists("key-10000", "val-10000"), true)
	assert.Equal(t, cm.Has("key-10000"), true)
	cm.Remove("key-10000")
	assert.Equal(t, cm.Has("key-10000"), false)
	reflect.DeepEqual(fakeInput, cm.Map())
}
