package rnokpp_test

import (
	"testing"
	"time"

	"github.com/fre5h/rnokpp"
)

func TestErrNotAllowedDate(t *testing.T) {
	date, _ := time.Parse("02.01.2006", "31.12.1899")
	x := &rnokpp.ErrNotAllowedDate{Date: date}

	if x.Error() != "the allowed dates start from 01.01.1900, but your date 31.12.1899 is earlier" {
		t.Error("Wrong error message for ErrNotAllowedDate")
	}
}

func TestErrDateInFuture(t *testing.T) {
	date, _ := time.Parse("02.01.2006", "01.01.3000")
	x := &rnokpp.ErrDateInFuture{Date: date}

	if x.Error() != "it is allowed to use only dates in past or current date, but your date is in the future 01.01.3000" {
		t.Error("Wrong error message for ErrDateInFuture")
	}
}
