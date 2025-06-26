//go:build !solution

package mycheck

import (
	"errors"
	"strings"
	"unicode"
)

type myError []string

func (err *myError) AddError(newErr string) {
	*err = append(*err, newErr)
}

func (err myError) Error() string {
	if len(err) == 0 {
		return ""
	}
	return strings.Join(err, ";")
}

func MyCheck(input string) error {
	var errs myError

	var spaceCounter int
	var foundNumbers, isLong bool

	if len(input) >= 20 {
		isLong = true
	}

	for _, val := range input {
		if unicode.IsDigit(val) {
			foundNumbers = true
		} else if val == ' ' {
			spaceCounter++
		}
	}

	if foundNumbers {
		errs.AddError("found numbers")
	}
	if isLong {
		errs.AddError("line is too long")
	}
	if spaceCounter != 2 {
		errs.AddError("no two spaces")
	}
	return errors.New(errs.Error())
}
