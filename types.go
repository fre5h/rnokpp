package rnokpp

import (
	"fmt"
	"time"
)

type Gender string

const Male = Gender("male")
const Female = Gender("female")
const Unknown = Gender("unknown")

func (g Gender) IsMale() bool {
	return g == Male
}

func (g Gender) IsFemale() bool {
	return g == Female
}

func (g Gender) String() string {
	return string(g)
}

var BaseLocation, _ = time.LoadLocation("Europe/Kiev")
var BaseDate = time.Date(1899, 12, 31, 0, 0, 0, 0, BaseLocation)

type Details struct {
	Valid    bool
	Gender   Gender
	Birthday time.Time
}

func NewDetails(valid bool, gender Gender, date string) *Details {
	birthday, err := time.ParseInLocation("02.01.2006", date, BaseLocation)
	if err != nil {
		panic(err)
	}

	return &Details{valid, gender, birthday}
}

func (d Details) String() string {
	v := "valid"
	if !d.Valid {
		v = "invalid"
	}

	return fmt.Sprintf("%s, %s, %s", v, d.Gender, d.Birthday.Format("02.01.2006"))
}
