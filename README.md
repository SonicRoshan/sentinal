# Sentinal
Data Validation Library In Go

## Basic Example
```go
type user struct {
    age `min:"14" max:"100"` // simply add function name as tag and validation data
}

func main() {

    data := user{15}
    valid, err := sentinal.Validate(user)
    //valid will be true

    data := user{10}
    valid, err = sentinal.Validate(user)
    //valid will be false

    data := user{101}
    valid, err = sentinal.Validate(user)
    //valid will be false
}
```

## Functions List

### Max
#### Use case
```go
type user struct {
    age `max:"100"`
}
```
This will check if value in struct is ```<=``` than value provided in struct tag, which in this case is 100

### Max Exclusive
#### Use case
```go
type user struct {
    favNumber float64 `maxExclusive:"50.5"`
}
```
This will check if value in struct is ```<``` than value provided in struct tag, which in this case is 50.5. This would also work with floats.

### Min
#### Use case
```go
type user struct {
    age `min:"14"`
}
```
This will check if value in struct is ```>=``` than value provided in struct tag, which in this case is 14. This would also work with floats.

### Min Exclusive
#### Use case
```go
type user struct {
    favNumber `min:"15.5"`
}
```
This will check if value in struct is ```>``` than value provided in struct tag, which in this case is 15.5. This would also work with floats.