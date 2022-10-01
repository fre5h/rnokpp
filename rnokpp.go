// Package rnokpp реалізує функції для роботи з РНОКПП (Реєстраційний номер облікової картки платника податків).
//
// З моменту впровадження державного реєстру фізичних осіб України у 1994 році мав назву «індивідуальний ідентифікаційний номер».
// З 2012 року набув чинності Податковий кодекс України, у якому використовується термін реєстраційний номер облікової картки платника податків (РНОКПП)
// — як десятизначний номер з Державного реєстру фізичних осіб — платників податків.
package rnokpp

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	mathRand "math/rand"
	"strconv"
	"time"
)

var maleDigits = [5]int{1, 3, 5, 7, 9}
var femaleDigits = [5]int{0, 2, 4, 6, 8}

func init() {
	var b [10]byte

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
	numberOfDaysSinceBaseDate--

	details := Details{
		Valid:    true,
		Gender:   gender,
		Birthday: BaseDate.AddDate(0, 0, numberOfDaysSinceBaseDate),
	}

	return &details, nil
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
func GetGender(rnokpp string) (*Gender, error) {
	details, err := GetDetails(rnokpp)

	if err != nil {
		return nil, err
	}

	return &details.Gender, nil
}

// GenerateRnokpp generates valid RNOKPP by date and gender
func GenerateRnokpp(date time.Time, gender Gender) (rnokpp string, err error) {
	if date.Before(BaseDate) {
		err = fmt.Errorf("the allowed dates start from 01.01.1900, but your date is %s", date.Format("02.04.2006"))

		return
	}

	if date.After(time.Now()) {
		err = fmt.Errorf("it is allowed to use only dates in past or current date, but your date is in the future %s", date.Format("02.04.2006"))

		return
	}

	diff := date.Sub(BaseDate)
	numberOfDays := int(diff.Hours() / 24)
	numberOfDays--
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

// parseRnokpp parses RNOKPP from string into array of integers
func parseRnokpp(rnokpp string) (result [10]int, err error) {
	lengthRnokpp := len(rnokpp)

	if lengthRnokpp > 10 {
		return result, fmt.Errorf("more than 10 symbols, expects exactly 10 symbols")
	}

	if lengthRnokpp < 10 {
		return result, fmt.Errorf("less than 10 symbols, expects exactly 10 symbols")
	}

	for i := 0; i < 10; i++ {
		result[i], err = strconv.Atoi(string(rnokpp[i]))

		if err != nil {
			return result, fmt.Errorf("string does not consist of digits")
		}
	}

	return result, nil
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
