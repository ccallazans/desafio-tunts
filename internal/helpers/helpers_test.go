package helpers

import (
	"sort"
	"testing"
)

func TestGetKeys(t *testing.T) {

	data := map[string]interface{}{
		"Key1": nil,
		"Key2": 1,
		"Key3": true,
		"key4": "test",
	}

	out := GetKeys(data)
	sort.Strings(out)

	expected := []string{"Key1", "Key2", "Key3", "key4"}

	if len(out) != len(expected) {
		t.Log("expected: ", expected, " got: ", out)
		t.Fail()
		return
	}
	for i, v := range out {
		if v != expected[i] {
			t.Log("expected: ", expected, " got: ", out)
			t.Fail()
			return
		}
	}
}
