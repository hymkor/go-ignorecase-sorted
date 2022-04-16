//go:build ignore
// +build ignore

package main

import (
	"github.com/hymkor/go-ignorecase-sorted"
)

func main() {
	var dic ignoreCaseSorted.Dictionary[int]

	dic.Store("a", 1)
	dic.Store("b", 2)
	dic.Store("c", 3)
	dic.Store("d", 3)

	if val, ok := dic.Load("A"); ok {
		println("dic[`A`]=", val)
	}
	dic.Delete("D")
	println("------")

	dic.Range(func(key string, val int) {
		println("dic[`"+key+"`]=", val)
	})
}
