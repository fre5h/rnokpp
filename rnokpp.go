/*
	Допоміжні функції для роботи з РНОКПП (Реєстраційний номер облікової картки платника податків).

    З моменту впровадження державного реєстру фізичних осіб України у 1994 році мав назву «індивідуальний ідентифікаційний номер».
    З 2012 року набув чинності Податковий кодекс України, у якому використовується термін реєстраційний номер облікової картки платника податків (РНОКПП)
    — як десятизначний номер з Державного реєстру фізичних осіб — платників податків.

    https://uk.wikipedia.org/wiki/%D0%A0%D0%B5%D1%94%D1%81%D1%82%D1%80%D0%B0%D1%86%D1%96%D0%B9%D0%BD%D0%B8%D0%B9_%D0%BD%D0%BE%D0%BC%D0%B5%D1%80_%D0%BE%D0%B1%D0%BB%D1%96%D0%BA%D0%BE%D0%B2%D0%BE%D1%97_%D0%BA%D0%B0%D1%80%D1%82%D0%BA%D0%B8_%D0%BF%D0%BB%D0%B0%D1%82%D0%BD%D0%B8%D0%BA%D0%B0_%D0%BF%D0%BE%D0%B4%D0%B0%D1%82%D0%BA%D1%96%D0%B2
*/

package rnokpp

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	mathRand "math/rand"
	"strconv"
	"time"
)

func init() {
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// GetDetails returns details about RNOKPP if possible
func GetDetails(rnokpp string) (*Details, error) {
	pRnokpp, err := parseRnokpp(rnokpp)

	if err != nil {
		return nil, err
	}

	genderDigit := pRnokpp[8]  // gender digit
	controlDigit := pRnokpp[9] // control digit

	digits := [9]int{pRnokpp[0], pRnokpp[1], pRnokpp[2], pRnokpp[3], pRnokpp[4], pRnokpp[5], pRnokpp[6], pRnokpp[7], pRnokpp[8]}
	calculatedControlDigit := calculateControlDigit(digits)

	if controlDigit != calculatedControlDigit {
		return nil, fmt.Errorf("invalid")
	}

	var gender Gender

	if genderDigit%2 == 0 {
		gender = Female
	} else {
		gender = Male
	}

	numberOfDaysSinceBaseDate := pRnokpp[0]*10000 + pRnokpp[1]*1000 + pRnokpp[2]*100 + pRnokpp[3]*10 + pRnokpp[4]*1

	details := Details{
		Valid:    true,
		Gender:   gender,
		Birthday: BaseDate.AddDate(0, 0, numberOfDaysSinceBaseDate),
	}

	return &details, nil
}

func parseRnokpp(rnokpp string) (result [10]int, err error) {
	var empty [10]int

	if len(rnokpp) > 10 {
		return empty, fmt.Errorf("more than 10 digits")
	}

	for i := 0; i < len(rnokpp); i++ {
		result[i], err = strconv.Atoi(string(rnokpp[i]))

		if err != nil {
			return empty, fmt.Errorf("string does not consist of digits")
		}
	}

	return result, nil
}

// IsValid checks if RNOKPP is valid
func IsValid(rnokpp string) bool {
	details, err := GetDetails(rnokpp)

	if err != nil {
		return false
	}

	return details.Valid
}

// IsMale checks if RNOKPP belongs to the male gender
func IsMale(rnokpp string) (bool, error) {
	details, err := GetDetails(rnokpp)

	if err != nil {
		return false, err
	}

	return details.Gender.IsMale(), nil
}

// IsFemale checks if RNOKPP belongs to the female gender
func IsFemale(rnokpp string) (bool, error) {
	details, err := GetDetails(rnokpp)

	if err != nil {
		return false, err
	}

	return details.Gender.IsFemale(), nil
}

// GetGender gets gender from RNOKPP
func GetGender(rnokpp string) (Gender, error) {
	details, err := GetDetails(rnokpp)

	if err != nil {
		return Unknown, err
	}

	return details.Gender, nil
}

// func RandomRnokpp() string {
// }
//
// func RandomRnokppN(count int) (result []string) {
// 	for i := 0; i < count; i++ {
// 		result = append(result, RandomRnokpp())
// 	}
//
// 	return
// }

var maleDigits = [5]int{1, 3, 5, 7, 9}
var femaleDigits = [5]int{0, 2, 4, 6, 8}

// GenerateRnokpp generates RNOKPP by date and gender
func GenerateRnokpp(date time.Time, gender Gender) (rnokpp string) {
	diff := date.Sub(BaseDate)
	numberOfDays := int(diff.Hours() / 24)
	rnokpp = fmt.Sprintf("%05d", numberOfDays)

	// three random account number digits
	rnokpp += strconv.Itoa(mathRand.Intn(9))
	rnokpp += strconv.Itoa(mathRand.Intn(9))
	rnokpp += strconv.Itoa(mathRand.Intn(9))

	if gender == Male {
		rnokpp += strconv.Itoa(maleDigits[mathRand.Intn(4)])
	} else {
		rnokpp += strconv.Itoa(femaleDigits[mathRand.Intn(4)])
	}

	var digits [9]int
	for i := 0; i < len(rnokpp); i++ {
		digits[i], _ = strconv.Atoi(string(rnokpp[i]))
	}

	rnokpp += strconv.Itoa(calculateControlDigit(digits))

	return
}

// calculateControlDigit calculates 10th (control digit) from the first 9 digits
func calculateControlDigit(rnokpp [9]int) int {
	checksum := rnokpp[0] * -1
	checksum += rnokpp[1] * 5
	checksum += rnokpp[2] * 7
	checksum += rnokpp[3] * 9
	checksum += rnokpp[4] * 4
	checksum += rnokpp[5] * 6
	checksum += rnokpp[6] * 10
	checksum += rnokpp[7] * 5
	checksum += rnokpp[8] * 7

	return (checksum % 11) % 10
}
