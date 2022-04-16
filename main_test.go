package ignoreCaseSorted_test

import (
	"testing"

	"github.com/hymkor/go-ignorecase-sorted"
)

func Test1(t *testing.T) {
	var dic ignoreCaseSorted.Dictionary[int]

	dic.Store("A", 1)
	dic.Store("B", 2)
	dic.Store("c", 3)

	if val, ok := dic.Load("A"); !ok || val != 1 {
		t.Fatalf("expect Load('A')==1 but %d", val)
		return
	}
	dic.Delete("C")
	if val, ok := dic.Load("b"); !ok || val != 2 {
		t.Fatalf("Load('b')==2 but %d", val)
		return
	}
	dic.Store("C", 3)

	i := 0
	dic.Range(func(key string, val int) {
		var result bool
		switch i {
		case 0:
			result = (key == "A" && val == 1)
		case 1:
			result = (key == "B" && val == 2)
		case 2:
			result = (key == "C" && val == 3)
		}
		if !result {
			t.Fatalf("Range fails at %d (%s and %d)", i, key, val)
		}
		i++
	})
}
