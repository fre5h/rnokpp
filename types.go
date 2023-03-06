package rnokpp

import (
	"fmt"
	mathRand "math/rand"
	"time"

	"github.com/fre5h/rnokpp/internal"
)

// Gender is a string type representing a gender (male/female)
type Gender string

const Male = Gender("male")
const Female = Gender("female")

// String returns string representation of gender
func (g Gender) String() string {
	return string(g)
}

// IsMale checks if gender is male
func (g Gender) IsMale() bool {
	return g == Male
}

// IsFemale checks if gender is female
func (g Gender) IsFemale() bool {
	return g == Female
}

// RandomGender returns random gender
func RandomGender() Gender {
	if mathRand.Intn(1) == 0 {
		return Female
	}

	return Male
}

// Details is a struct representing details of RNOKPP
type Details struct {
	Valid    bool
	Gender   Gender
	Birthday time.Time
}

// String returns string representation of RNOKPP details
func (d Details) String() string {
	if !d.Valid {
		return "invalid"
	}

	return fmt.Sprintf("valid, %s, %s", d.Gender, d.Birthday.Format("02.01.2006"))
}

// NewDetails creates Details, and returns the pointer to it
func NewDetails(valid bool, gender Gender, date string) *Details {
	birthday, err := time.ParseInLocation("02.01.2006", date, internal.BaseLocation)
	if err != nil {
		panic(err)
	}

	return &Details{valid, gender, birthday}
}
