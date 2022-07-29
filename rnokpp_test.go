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

var testVariantsForTestGetDetails = []struct {
	rnokpp string
	valid  bool
	gender rnokpp.Gender
	date   string
}{
	{
		rnokpp: "3652504575",
		valid:  true,
		gender: rnokpp.Male,
		date:   "01.01.2000",
	},
}

func TestGetDetails(t *testing.T) {
	for _, data := range testVariantsForTestGetDetails {
		result, _ := rnokpp.GetDetails(data.rnokpp)

		if result.Valid != data.valid {
			t.Error("invalid", result.Valid, data.valid)
		}

		if result.Gender != data.gender {
			t.Error("different gender", result.Gender, data.gender)
		}

		if result.Birthday.Format("02.01.2006") != data.date {
			t.Error("different birthday", result.Birthday.Format("02.01.2006"), data.date)
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
	result, err := rnokpp.IsMale("3652504575")
	if err == nil {
		fmt.Print(result)
	}
	// Output:
	// true
}

func ExampleIsFemale() {
	result, err := rnokpp.IsFemale("3652504575")
	if err == nil {
		fmt.Print(result)
	}
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

func ExampleGenerateRnokpp() {
	birthday, _ := time.Parse("02.01.2006", "01.01.2000")
	generatedRnokppMale := rnokpp.GenerateRnokpp(birthday, rnokpp.Male)     // string with valid random RNOKPP
	generatedRnokppFemale := rnokpp.GenerateRnokpp(birthday, rnokpp.Female) // string with valid random RNOKPP

	fmt.Print(rnokpp.IsValid(generatedRnokppMale), rnokpp.IsValid(generatedRnokppFemale))
	// Output:
	// true true
}
