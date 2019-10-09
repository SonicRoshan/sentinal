[![Go Report Card](https://goreportcard.com/badge/github.com/SonicRoshan/sentinal)](https://goreportcard.com/report/github.com/SonicRoshan/sentinal) [![GoDoc](https://godoc.org/github.com/SonicRoshan/sentinal?status.svg)](https://godoc.org/github.com/SonicRoshan/sentinal) [![GoCover](https://gocover.io/_badge/github.com/SonicRoshan/sentinal)](https://gocover.io/github.com/SonicRoshan/sentinal)

# Sentinal
Data Validation Library In Go

## Basic Example
```go
type user struct {
    age `min:"14" max:"100"` // simply add function name as tag and validation data
}

func main() {

    data := user{15}
    valid, data, err := sentinal.Validate(user)
    //valid will be true

    data := user{10}
    valid, data, err = sentinal.Validate(user)
    //valid will be false
    //data = {"age" : ["min is 14"]}

    data := user{101}
    valid, data, err = sentinal.Validate(user)
    //valid will be false
    //data = {"age" : ["max is 100"]}
}
```


## Custom Validation Functions
```go
type user struct {
    name string `isName:"true"` // isName is a custom functions
}

func isName(value reflect.Value, validationData string) (bool, string, error) {
    // value is the reflect.Value of the field.
    // validationData is the data provided in struct tag, which in this case is true.
    
    //SOME LOGIC HERE

    /*
    first thing to return is if the data is valid.
    second thing to return is a string message when data is invalid.
    This helps you to see what was invalid.
    third thing to return is an error if any.
    */
    
    return true, "", nil
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

### From
Checks if a value is in a list.
Note - there should not be a space between commas.
```go
type address struct {
    country string `from:"england,france,spain"`
}
```

### Not From
Checks if a value is not in a list.
Note - there should not be a space between commas.
```go
type user struct {
    favColor string `notFrom:"black,white,blue"`
}
```

### Not Empty
Checks if a field is not empty. Works with any data type
```go
type user struct {
    username string `notEmpty:"true"`
    age int `notEmpty:"true"`
}
```