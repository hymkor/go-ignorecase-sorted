//go:build ignore

package main

import (
	"github.com/hymkor/go-ignorecase-sorted/v2"
)

func main() {
	var dic ignoreCaseSorted.Dictionary[int]

	dic.Set("a", 1)
	dic.Set("b", 2)
	dic.Set("c", 3)
	dic.Set("d", 3)

	println("------ test the case of keys are ignored ------")
	if val, ok := dic.Get("A"); ok {
		println("dic[`A`]=", val, "whose key is set as `a`")
	}
	dic.Delete("D")
	println()

	println("------ iterate with callback function -----")
	dic.Each(func(key string, value int) bool {
		println("dic[`"+key+"`]=", value)
		return true
	})
	println()

	println("------ iterate with rangefunc of Go 1.22 -----")
	for key, value := range dic.Each {
		println("dic[`"+key+"`]=", value)
	}
	println()
}
