package conmap

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func genFakeInput() (m map[string]interface{}) {
	m = make(map[string]interface{})
	for i := 0; i < 100000; i++ {
		m[fmt.Sprintf("key-%07d", i)] = fmt.Sprintf("val-%07d", i)
	}
	return m
}

func TestConMapCRUD(t *testing.T) {
	fakeInput := genFakeInput()

	cm := NewConMap()
	cm.BatchStore(fakeInput)
	assert.Equal(t, cm.Empty(), false)
	assert.Equal(t, cm.Count(), len(fakeInput))
	assert.Equal(t, cm.StoreIfNotExists("key-0100001", "val-0100001"), true)
	assert.Equal(t, cm.Has("key-0100001"), true)
	cm.Remove("key-0100001")
	assert.Equal(t, cm.Has("key-0100001"), false)
	reflect.DeepEqual(fakeInput, cm.Map())
}
