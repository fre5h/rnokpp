package rnokpp

import (
	"fmt"
	"time"
)

// Gender is a string type representing a gender (male/female
type Gender string

const Male = Gender("male")
const Female = Gender("female")
const Unknown = Gender("unknown")

// IsMale checks if gender is male
func (g Gender) IsMale() bool {
	return g == Male
}

// IsFemale checks if gender is female
func (g Gender) IsFemale() bool {
	return g == Female
}

// String returns string representation of gender
func (g Gender) String() string {
	return string(g)
}

var baseLocation, _ = time.LoadLocation("Europe/Kiev")
var baseDate = time.Date(1899, 12, 31, 0, 0, 0, 0, baseLocation)

// Details is a struct representing details of RNOKPP
type Details struct {
	Valid    bool
	Gender   Gender
	Birthday time.Time
}

// NewDetails creates Details, and returns the pointer to it.
func NewDetails(valid bool, gender Gender, date string) *Details {
	birthday, err := time.ParseInLocation("02.01.2006", date, baseLocation)
	if err != nil {
		panic(err)
	}

	return &Details{valid, gender, birthday}
}

// String returns string representation of RNKOPP details
func (d Details) String() string {
	if !d.Valid {
		return "invalid"
	}

	return fmt.Sprintf("valid, %s, %s", d.Gender, d.Birthday.Format("02.01.2006"))
}
