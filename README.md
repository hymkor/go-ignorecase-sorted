go-ignorecase-sorted.Dictionary
===============================

- Container library like map[string]T using Generics
- Ignore cases of string-keys
- Iterators can read elements by sorted-order
- It can be used with GOEXPERIMENT=rangefunc of Go 1.22

```example.go
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
}
```

**go run example.go**
```go run example.go|
------ The case of keys are ignored ------
dic[`A`]= 1 whose key is set as `a`
------ type0 iterator -----
dic[`a`]= 1
dic[`b`]= 2
dic[`c`]= 3
------ Go 1.22 rangefunc iterator -----
dic[`a`]= 1
dic[`b`]= 2
dic[`c`]= 3
```
