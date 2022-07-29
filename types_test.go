package rnokpp_test

import (
	"testing"
	"time"

	"github.com/fre5h/rnokpp"
)

const reset = "\033[0m"
const red = "\033[31m"
const green = "\033[32m"

func TestNewDetails(t *testing.T) {
	details := rnokpp.NewDetails(true, rnokpp.Male, "01.01.2000")

	expectedDate, _ := time.ParseInLocation("02.01.2006", "01.01.2000", rnokpp.BaseLocation)

	if details.Birthday.Unix() != expectedDate.Unix() {
		t.Errorf(
			"wrong timestamp for birthday, expected: %s%d%s, actual: %s%d%s.",
			green, expectedDate.Unix(), reset,
			red, details.Birthday.Unix(), reset,
		)
	}
}

func TestNewDetailsWithPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	rnokpp.NewDetails(true, rnokpp.Male, "not date")
}

var testVariantsTestDetailsString = []struct {
	string  string
	details rnokpp.Details
}{
	{
		string:  "valid, male, 01.01.2000",
		details: *rnokpp.NewDetails(true, rnokpp.Male, "01.01.2000"),
	},
	{
		string:  "valid, female, 01.01.2001",
		details: *rnokpp.NewDetails(true, rnokpp.Female, "01.01.2001"),
	},
	{
		string:  "invalid",
		details: *rnokpp.NewDetails(false, rnokpp.Male, "01.01.2002"),
	},
	{
		string:  "invalid",
		details: *rnokpp.NewDetails(false, rnokpp.Female, "01.01.2003"),
	},
}

func TestDetailsString(t *testing.T) {
	for _, data := range testVariantsTestDetailsString {
		if data.string != data.details.String() {
			t.Errorf(
				"invalid string representation of details, expected: %s%s%s, actual: %s%s%s.",
				green, data.string, reset,
				red, data.details.String(), reset,
			)
		}
	}
}
