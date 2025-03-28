package data

import "errors"

var (
	ErrRecordNotFound = errors.New("no record found")
	ErrDuplicateEmail = errors.New("duplicate email")
)
