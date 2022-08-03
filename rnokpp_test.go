package rnokpp_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/fre5h/rnokpp"
)

const reset = "\033[0m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"

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
	for _, testVariant := range testVariantsForTestGetDetails {
		result, _ := rnokpp.GetDetails(testVariant.rnokpp)

		if result.Valid != testVariant.valid {
			t.Error("Invalid", result.Valid, testVariant.valid)
		}

		if result.Gender != testVariant.gender {
			t.Error("Different gender", result.Gender, testVariant.gender)
		}

		if result.Birthday.Format("02.01.2006") != testVariant.date {
			t.Error("Different birthday", result.Birthday.Format("02.01.2006"), testVariant.date)
		}
	}
}

func TestIsMale(t *testing.T) {
	var err error
	var isMale bool

	isMale, err = rnokpp.IsMale("3652504575") // male RNOKPP
	if isMale == false || err != nil {
		t.Errorf(
			"RNOKPP %s%s%s expected to be a valid for male, but it is not",
			yellow, "3652504575", reset,
		)
	}

	isMale, err = rnokpp.IsMale("3068208400") // female RNOKPP
	if isMale == true || err != nil {
		t.Errorf(
			"RNOKPP %s%s%s expected to be a invalid for male, but it is not",
			yellow, "3068208400", reset,
		)
	}

	_, err = rnokpp.IsMale("invalid") // invalid RNOKPP
	if err == nil {
		t.Error("Expected error for invalid RNOKPP")
	}
}

func TestIsFemale(t *testing.T) {
	var err error
	var isFemale bool

	isFemale, err = rnokpp.IsFemale("3652504575") // male RNOKPP
	if isFemale == true || err != nil {
		t.Errorf(
			"RNOKPP %s%s%s expected to be a valid for female, but it is not",
			yellow, "3652504575", reset,
		)
	}

	isFemale, err = rnokpp.IsFemale("3068208400") // female RNOKPP
	if isFemale == false || err != nil {
		t.Errorf(
			"RNOKPP %s%s%s expected to be a invalid for female, but it is not",
			yellow, "3068208400", reset,
		)
	}

	_, err = rnokpp.IsFemale("invalid") // invalid RNOKPP
	if err == nil {
		t.Error("Expected error for invalid RNOKPP")
	}
}

func TestGetDetailsWithErrors(t *testing.T) {
	var err error

	_, err = rnokpp.GetDetails("1234567890+")
	if err == nil || err.Error() != "more than 10 digits, expects exactly 10 digits" {
		t.Error("Expected error for a string longer than 10 symbols")
	}

	_, err = rnokpp.GetDetails("123456789")
	if err == nil || err.Error() != "less than 10 digits, expects exactly 10 digits" {
		t.Error("Expected error for a string smaller than 10 symbols")
	}

	var testVariantsForTestGetDetailsForNonDigitsInString = []string{
		"123456789X",
		"          ",
		"ABCDEFGHIJ",
		" 234567890",
		"123456789 ",
	}

	for _, invalidRnokpp := range testVariantsForTestGetDetailsForNonDigitsInString {
		_, err := rnokpp.GetDetails(invalidRnokpp)
		if err == nil || err.Error() != "string does not consist of digits" {
			t.Error("Expected error for non digits in invalidRnokpp")
		}
	}
}

func TestGenerateRnokpp(t *testing.T) {
	for i := 0; i < 100; i++ {
		randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
		birthday := time.Unix(randomTime, 0)
		generatedRnokpp := rnokpp.GenerateRnokpp(birthday, rnokpp.Male)

		if !rnokpp.IsValid(generatedRnokpp) {
			t.Errorf(
				"Generated RNOKPP is not valid %s\"%s\"%s",
				red, generatedRnokpp, reset,
			)
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
	result1, err1 := rnokpp.IsMale("3652504575") // male RNOKPP
	if err1 == nil {
		fmt.Println(result1)
	}

	result2, err2 := rnokpp.IsMale("3068208400") // female RNOKPP
	if err2 == nil {
		fmt.Println(result2)
	}

	// Output:
	// true
	// false
}

func ExampleIsFemale() {
	result1, err1 := rnokpp.IsFemale("3652504575") // male RNOKPP
	if err1 == nil {
		fmt.Println(result1)
	}

	result2, err2 := rnokpp.IsFemale("3068208400") // female RNOKPP
	if err2 == nil {
		fmt.Println(result2)
	}

	// Output:
	// false
	// true
}

func ExampleGenerateRnokpp() {
	birthday, _ := time.Parse("02.01.2006", "01.01.2000")
	generatedRnokppMale := rnokpp.GenerateRnokpp(birthday, rnokpp.Male)     // string with valid random RNOKPP for male
	generatedRnokppFemale := rnokpp.GenerateRnokpp(birthday, rnokpp.Female) // string with valid random RNOKPP for female

	fmt.Print(rnokpp.IsValid(generatedRnokppMale), rnokpp.IsValid(generatedRnokppFemale))
	// Output:
	// true true
}
