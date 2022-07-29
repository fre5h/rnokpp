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

type Details struct {
	Valid    bool
	Gender   Gender
	Birthday time.Time
}

func (d Details) String() string {
	v := "valid"
	if !d.Valid {
		v = "invalid"
	}

	return fmt.Sprintf("%s, %s, %s", v, d.Gender, d.Birthday.Format("02.01.2006"))
}
