go-ignorecase-sorted.Dictionary
===============================

- Container library like map[string]T using Generics
- Ignore cases of string-keys
- Iterators can read elements by sorted-order

```go
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
```

```
$ go run example.go
------ The case of keys are ignored ------
dic[`A`]= 1 whose key is set as `a`
------ type0 iterator -----
dic[`a`]= 1
dic[`b`]= 2
dic[`c`]= 3
------ type1 iterator (ascending) -----
dic[`a`]= 1
dic[`b`]= 2
dic[`c`]= 3
------ type1 iterator (descending) -----
dic[`c`]= 3
dic[`b`]= 2
dic[`a`]= 1
------ type2 iterator (ascending) -----
dic[`a`]= 1
dic[`b`]= 2
dic[`c`]= 3
------ type2 iterator (descending) -----
dic[`c`]= 3
dic[`b`]= 2
dic[`a`]= 1
```
