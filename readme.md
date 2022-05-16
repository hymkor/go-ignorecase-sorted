go-ignorecase-sorted.Dictionary
===============================

- Container library like map[string]T using Generics
- Ignore cases of string-keys
- Can iterate with ascending

```go
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

	for p := dic.Front(); p != nil; p = p.Next() {
		println("dic[`"+p.Key+"`]=", p.Value)
	}
}
```

```
$ go run example.go
dic[`A`]= 1
------
dic[`a`]= 1
dic[`b`]= 2
dic[`c`]= 3
```
