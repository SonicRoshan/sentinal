package sentinal

var functions = map[string]functionType{
	"max":          maxInclusive,
	"min":          minInclusive,
	"maxExclusive": maxExclusive,
	"minExclusive": minExclusive,
	"from":         from,
	"notFrom":      notFrom,
	"notEmpty":     notEmpty,
	"maxLen":       maxLength,
	"minLen":       minLength,
	"contains":     contains,
	"notContains":  notContains,
}
