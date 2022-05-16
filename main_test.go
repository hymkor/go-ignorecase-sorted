package ignoreCaseSorted_test

import (
	"testing"

	"github.com/hymkor/go-ignorecase-sorted"
)

func TestAscend(t *testing.T) {
	var dic ignoreCaseSorted.Dictionary[int]

	dic.Set("x", 7)
	dic.Set("y", 8)
	dic.Set("Z", 9)

	expect := []struct {
		key   string
		value int
	}{
		{key: "x", value: 7},
		{key: "y", value: 8},
		{key: "Z", value: 9},
	}
	for p := dic.Front(); p != nil; p = p.Next() {
		if expect[0].key != p.Key {
			t.Fatalf("'%s' != '%s'", expect[0].key, p.Key)
			return
		}
		if expect[0].value != p.Value {
			t.Fatalf("'%d' != '%d'", expect[0].value, p.Value)
			return
		}
		expect = expect[1:]
	}
}

func TestDesend(t *testing.T) {
	var dic ignoreCaseSorted.Dictionary[int]

	dic.Set("x", 7)
	dic.Set("y", 8)
	dic.Set("Z", 9)

	expect := []struct {
		key   string
		value int
	}{
		{key: "Z", value: 9},
		{key: "y", value: 8},
		{key: "x", value: 7},
	}
	for p := dic.Back(); p != nil; p = p.Prev() {
		if expect[0].key != p.Key {
			t.Fatalf("'%s' != '%s'", expect[0].key, p.Key)
			return
		}
		if expect[0].value != p.Value {
			t.Fatalf("'%d' != '%d'", expect[0].value, p.Value)
			return
		}
		expect = expect[1:]
	}
}
