package validator

import "regexp"

// EmailRX is a regular expression for matching email addresses.
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validator is a type which contains map of validation errors.
type Validator struct {
	Errors map[string]string
}

// New is a helper method which create new Validator instance with empty errors map.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid return true if the errors map does not contain any entry.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds an error message to the map (so long as no entry already exists for
// the given key).
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// Matches returns true if a string value matches a specific regexp pattern.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique is a generic function which returns true if all values in a slice
// are unique.
func Unique[T comparable](values []T) bool {
	uniqueValues := map[T]bool{}
	for _, v := range values {
		uniqueValues[v] = true
	}
	return len(uniqueValues) == len(values)
}