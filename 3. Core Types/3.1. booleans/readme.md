# `True` and `False`

`True` and `False` logic is represented using the Boolean type, `bool`. Use this type when need an on/off switch in your code. The value of a `bool` can only ever be `true` or `false`.  
The Zero value of `bool` is `false`.

Wheen using a comparison operator such as `==` or `>`, the result of that comparison is a `bool` value.

Code Example:
```go
package main

import "fmt"

func main() {
    fmt.Println(10 > 5) // result is: true
    fmt.Println(10 == 5) // result is: false
}
```