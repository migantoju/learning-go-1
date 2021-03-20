# ENUMS

Enums are a way to defining a fixed list of values that are all related.  
Go doesn't have a built-in type for enums, but it does provide tools such as `iota`
to let you define your own using `constants`.

### Example
In the following code, we've the days of the week defined as `constants`.  

```go
...
const (
    Sunday = 0
    Monday = 1
    Tuesday = 2
    Wednesday = 3
    Thursday = 4
    Friday = 5
    Saturday = 6 
)
...
```
With `iota`, GO, helps us to manage lists justs like this. Using `iota`, the following code is equal to the preceding code.

```go
...
const = (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
...
```

Now, we have `iota` assinging the numbers for us. Using `iota` makes enums easier to create and mantain, especially if you need a new value to the middle of the code later.