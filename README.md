[![Go Report Card](https://goreportcard.com/badge/github.com/SonicRoshan/sentinal)](https://goreportcard.com/report/github.com/SonicRoshan/sentinal) [![GoDoc](https://godoc.org/github.com/SonicRoshan/sentinal?status.svg)](https://godoc.org/github.com/SonicRoshan/sentinal) [![GoCover](https://gocover.io/_badge/github.com/SonicRoshan/sentinal)](https://gocover.io/github.com/SonicRoshan/sentinal)

# Sentinal
Data Validation Library In Go

## Basic Example
```go
type user struct {
    age int
}

func main() {

    schema := map[string]map[string]string{
        "age" : map[string]string{
            "min" : "14",
            "max" : "100",
        },
    }

    data := user{15}
    valid, msg, err := sentinal.Validate(user, schema)
    // valid will be true

    data := user{10}
    valid, msg, err = sentinal.Validate(user, schema)
    // valid will be false
    // msg = {"age" : ["min is 14"]}

    data := user{101}
    valid, msg, err = sentinal.Validate(user, schema)
    // valid will be false
    // msg = {"age" : ["max is 100"]}
}
```


## Custom Validation Functions
```go
type user struct {
    name string
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


func main() {
    data := user{"Name"}

    schema := map[string]map[string]string{
        "name" : map[string]string{
            "isName" : "true", // isName is custom function
        },
    }


    sentinal.Validate(data, schema, map[string]func(reflect.Value, string) (bool, string, error){
        "isName" : isName,
    })
}


```

## Functions List

### max
This will check if value in struct is ```<=``` than value provided in struct tag.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "max" : "8",
    },
}
```

### maxExclusive
This will check if value in struct is ```<``` than value provided in struct tag.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "maxExclusive" : "8",
    },
}
```

### min
This will check if value in struct is ```>=``` than value provided in struct tag.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "min" : "8",
    },
}
```

### minExclusive
This will check if value in struct is ```>``` than value provided in struct tag, which in this case is 15.5. This would also work with floats.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "minExclusive" : "8",
    },
}
```

### From
Checks if a value is in a list.
Note - there should not be a space between commas.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "from" : "value1,value2,value3",
    },
}
```

### notFrom
Checks if a value is not in a list.
Note - there should not be a space between commas.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "notFrom" : "value1,value2,value3",
    },
}
```

### notEmpty
Checks if a field is not empty. Works with any data type.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "notEmpty" : "true",
    },
}
```

### maxLength
Checks if length of data is less than max. Works with any data type.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "maxLen" : "8",
    },
}
```

### minLength
Checks if length of data is greater than min. Works with any data type.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "minLen" : "8",
    },
}
```

### contains
Checks if a string contains certain items. Seperated by comma
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "contains" : "test,test2",
    },
}
```

### notContains
Checks if a string does not contain certain items. Seperated by comma
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "notContains" : " ,test2",
    },
}
```

### isEmail
Checks if field is an email.
```go
schema := map[string]map[string]string{
    "fieldName" : map[string]string{
        "isEmail" : "true",
    },
}
```