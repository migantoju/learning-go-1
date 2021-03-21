# Loops

In real-world applications, you're often going to need to run the same logic repeteadly.
`Loops` are the simplest way of repeating your logic.

Go only has one looping statement, `for`, but it's a flexible one. There are two distinct 
forms: the first is used a lot for ordered collections such as arrays and slices.

The sort of loop used for ordered collections looks as follows:

```go
for <initial statement>; <condition>; <post statement> {
    <statements>
}
```

The `initial`, `condition`, and `post` statements are all optional, and it's possible to write a `for` loop like this:

```go
for {
    <statements>
}
```
This form would result in a loop that would run forever, also known as an infinite loop.