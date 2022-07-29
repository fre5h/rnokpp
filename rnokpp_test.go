package rnokpp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/fre5h/rnokpp"
)

// const reset = "\033[0m"
// const red = "\033[31m"
// const green = "\033[32m"
// const yellow = "\033[33m"

var testVariants = []struct {
	rnokpp  string
	details rnokpp.Details
}{
	{
		rnokpp:  "3652504575",
		details: *rnokpp.NewDetails(true, rnokpp.Male, "01.01.2000"),
	},
}

func TestGetDetails(t *testing.T) {
	for _, data := range testVariants {
		result, _ := rnokpp.GetDetails(data.rnokpp)

		if result.Valid != data.details.Valid {
			t.Error("invalid", result.Valid, data.details.Valid)
		}

		if result.Gender != data.details.Gender {
			t.Error("mismatch gender", result.Gender, data.details.Gender)
		}

		if result.Birthday.Format("02.01.2006") != data.details.Birthday.Format("02.01.2006") {
			t.Error("mismatch birthday", result.Birthday.Format("02.01.2006"), data.details.Birthday.Format("02.01.2006"))
		}
	}
}

func ExampleIsValid() {
	fmt.Print(rnokpp.IsValid("3652504575"), rnokpp.IsValid("1234567890"))
	// Output:
	// true false
}

func ExampleGetDetails() {
	details, _ := rnokpp.GetDetails("3652504575")
	fmt.Print(details)
	// Output:
	// valid, male, 01.01.2000
}

func ExampleGetGender() {
	result, _ := rnokpp.GetGender("3652504575")
	fmt.Print(result)
	// Output:
	// male
}

func ExampleIsMale() {
	result, _ := rnokpp.IsMale("3652504575")
	fmt.Print(result)
	// Output:
	// true
}

func ExampleIsFemale() {
	result, _ := rnokpp.IsFemale("3652504575")
	fmt.Print(result)
	// Output:
	// false
}

func TestGenerateRnokpp(t *testing.T) {
	birthday, _ := time.Parse("02.01.2006", "01.01.2000")
	generatedRnokpp := rnokpp.GenerateRnokpp(birthday, rnokpp.Male)

	if !rnokpp.IsValid(generatedRnokpp) {
		t.Error("generated RNOKPP is not valid", generatedRnokpp)
	}
}
