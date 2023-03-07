package rnokpp

import (
	"errors"
	"fmt"
	"time"
)

var ErrInvalidControlDigit = errors.New("invalid control digit")
var ErrNumberGreaterThanZero = errors.New("number of rnokpp should be greater than 0")
var ErrMoreThan10Digits = errors.New("more than 10 digits, expects exactly 10 digits")
var ErrLessThan10Digits = errors.New("less than 10 symbols, expects exactly 10 symbols")
var ErrStringDoesNotConsistOfDigits = errors.New("string does not consist of digits")

type ErrNotAllowedDate struct {
	Date time.Time
}

func (e *ErrNotAllowedDate) Error() string {
	return fmt.Sprintf("the allowed dates start from 01.01.1900, but your date %s is earlier", e.Date.Format("02.04.2006"))
}

type ErrDateInFuture struct {
	Date time.Time
}

func (e *ErrDateInFuture) Error() string {
	return fmt.Sprintf("it is allowed to use only dates in past or current date, but your date is in the future %s", e.Date.Format("02.04.2006"))
}
