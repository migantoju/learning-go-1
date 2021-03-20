# CONSTANTS

Constants are like variables, but you can't change their initial value.  
These are useful for situations where the value of a constant doesn't need 
to or shouldn't change when your code is running.

`constant` declarations are similar to `var` statements. With a constant, the initial vlaue is
required. Types are optional and inferred if left out. 

```go
constant <name> <type> = <value>
```  
or declare multiple constants
```go
constant (
    <name1> <type1> = <value1>
    <name2> <type2> = <value2>
    ...
    <nameN> <typeN> = <valueN>
)
```

In this exercise, we used `constants` to define values that don't need to change while the 
code is running. 