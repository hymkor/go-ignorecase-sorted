go-ignorecase-sorted.Dictionary
===============================

- Container library like map[string]T using Generics
- Ignore cases of string-keys
- Iterators can read elements by sorted-order
- It can be used with `GOEXPERIMENT=rangefunc` of Go 1.22
- This package is the version 2 now. Therefore, please import `github.com/hymkor/go-ignorecase-sorted/v2`

```example.go
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
```

**env GOEXPERIMENT=rangefunc go run example.go**

```env GOEXPERIMENT=rangefunc go run example.go|
------ test the case of keys are ignored ------
dic[`A`]= 1 whose key is set as `a`

------ iterate with callback function -----
dic[`a`]= 1
dic[`b`]= 2
dic[`c`]= 3

------ iterate with rangefunc of Go 1.22 -----
dic[`a`]= 1
dic[`b`]= 2
dic[`c`]= 3

```
