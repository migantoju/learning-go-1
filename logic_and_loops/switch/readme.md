# Switch Statements

While it's possible to add as many `else if` statements to an `if` as you want, at some point, it'll
get hard to read.

When this happens, you can use Go's logic alternative: `switch`. For situations where you would need a 
big `if` statement, `switch` can be a mora compact alternative.

`Switch` notation code snippet:

```go
switch <initial statement>; <expression> {
    case <expression>:
        <statements>
    case <expression>, <expression>:
        <statements>
    default:
        <statements>
}
```
