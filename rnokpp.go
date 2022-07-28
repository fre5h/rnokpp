// Допоміжні функції для роботи з РНОКПП (Реєстраційний номер облікової картки платника податків).
//
// З моменту впровадження державного реєстру фізичних осіб України у 1994 році мав назву «індивідуальний ідентифікаційний номер».
// З 2012 року набув чинності Податковий кодекс України, у якому використовується термін реєстраційний номер облікової картки платника податків (РНОКПП)
// — як десятизначний номер з Державного реєстру фізичних осіб — платників податків.
//
// https://uk.wikipedia.org/wiki/%D0%A0%D0%B5%D1%94%D1%81%D1%82%D1%80%D0%B0%D1%86%D1%96%D0%B9%D0%BD%D0%B8%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80_%D0%BE%D0%B1%D0%BB%D1%96%D0%BA%D0%BE%D0%B2%D0%BE%D1%97_%D0%BA%D0%B0%D1%80%D1%82%D0%BA%D0%B8_%D0%BF%D0%BB%D0%B0%D1%82%D0%BD%D0%B8%D0%BA%D0%B0_%D0%BF%D0%BE%D0%B4%D0%B0%D1%82%D0%BA%D1%96%D0%B2

package rnokpp

import (
	"fmt"
	"time"
)

type Gender string

const Male Gender = Gender("male")
const Female Gender = Gender("female")

var BaseLocation, _ = time.LoadLocation("Europe/Kiev")
var BaseDate = time.Date(1899, 12, 31, 0, 0, 0, 0, BaseLocation)

type Details struct {
	Valid    bool
	Gender   Gender
	Birthday time.Time
}

func IsValid(rnokpp string) bool {
	return true
}

func GetDetails(rnokpp [10]int) (*Details, error) {
	birthdayDigit1 := rnokpp[0] * -1
	birthdayDigit2 := rnokpp[1] * 5
	birthdayDigit3 := rnokpp[2] * 7
	birthdayDigit4 := rnokpp[3] * 9
	birthdayDigit5 := rnokpp[4] * 4
	accountCardDigit1 := rnokpp[5] * 6
	accountCardDigit2 := rnokpp[6] * 10
	accountCardDigit3 := rnokpp[7] * 5
	genderDigit := rnokpp[8] * 7
	controlDigit := rnokpp[9]

	checksum := birthdayDigit1 + birthdayDigit2 + birthdayDigit3 + birthdayDigit4 + birthdayDigit5 + accountCardDigit1 + accountCardDigit2 + accountCardDigit3 + genderDigit
	calculatedControlDigit := (checksum % 11) % 10

	if controlDigit != calculatedControlDigit {
		return nil, fmt.Errorf("invalid")
	}

	details := Details{
		Valid:    true,
		Gender:   Male,
		Birthday: time.Date(),
	}

	return nil
}

func RandomRnokpp() [10]int {

}

func RandomRnokppS() string {

}
