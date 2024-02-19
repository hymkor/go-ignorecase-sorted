//go:build ignore

package main

import (
	"github.com/hymkor/go-ignorecase-sorted"
)

func main() {
	var dic ignoreCaseSorted.Dictionary[int]

	dic.Set("a", 1)
	dic.Set("b", 2)
	dic.Set("c", 3)
	dic.Set("d", 3)

	println("------ The case of keys are ignored ------")
	if val, ok := dic.Get("A"); ok {
		println("dic[`A`]=", val, "whose key is set as `a`")
	}
	dic.Delete("D")

	println("------ type0 iterator -----")
	dic.Range(func(key string, value int) bool {
		println("dic[`"+key+"`]=", value)
		return true
	})

	println("------ Go 1.22 rangefunc iterator -----")
	for key, value := range dic.Range {
		println("dic[`"+key+"`]=", value)
	}

	println("------ type1 iterator (ascending) -----")
	for p := dic.Front(); p != nil; p = p.Next() {
		println("dic[`"+p.Key+"`]=", p.Value)
	}

	println("------ type1 iterator (descending) -----")
	for p := dic.Back(); p != nil; p = p.Prev() {
		println("dic[`"+p.Key+"`]=", p.Value)
	}

	println("------ type2 iterator (ascending) -----")
	for p := dic.Ascend(); p.Range(); {
		println("dic[`"+p.Key+"`]=", p.Value)
	}

	println("------ type2 iterator (descending) -----")
	for p := dic.Descend(); p.Range(); {
		println("dic[`"+p.Key+"`]=", p.Value)
	}

}
