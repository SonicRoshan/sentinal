package sentinal

import "github.com/asaskevich/govalidator"

var functions = map[string]functionType{
	"max":             maxInclusive,
	"min":             minInclusive,
	"maxExclusive":    maxExclusive,
	"minExclusive":    minExclusive,
	"from":            from,
	"notFrom":         notFrom,
	"notEmpty":        notEmpty,
	"maxLen":          maxLength,
	"minLen":          minLength,
	"contains":        contains,
	"notContains":     notContains,
	"isEmail":         isEmail,
	"hasUpperCase":    simpleOverlay(govalidator.HasUpperCase, "Field does not have uppercase characters"),
	"notHasUpperCase": simpleOverlayReverse(govalidator.HasUpperCase, "Field has uppercase characters"),
}
