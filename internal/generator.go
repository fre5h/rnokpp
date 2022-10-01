package internal

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	mathRand "math/rand"
	"time"
)

func init() {
	var b [10]byte

	_, err := cryptoRand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// GenerateRandomDate generates random datetime in the interval [01.01.1900..current time]
func GenerateRandomDate() time.Time {
	var now = time.Now().In(BaseLocation)
	var randomMonth, randomDay int

	randomYear := mathRand.Intn(now.Year()-BaseYear+1) + BaseYear

	if randomYear == now.Year() {
		randomMonth = mathRand.Intn(int(now.Month())) + 1
	} else {
		randomMonth = mathRand.Intn(11) + 1
	}

	if randomMonth == int(now.Month()) {
		randomDay = mathRand.Intn(now.Day()) + 1
	} else {
		randomDay = mathRand.Intn(getNumberOfDaysInMonth(randomYear, time.Month(randomMonth))) + 1
	}

	return time.Date(randomYear, time.Month(randomMonth), randomDay, 0, 0, 0, 0, BaseLocation)
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func getNumberOfDaysInMonth(year int, month time.Month) int {
	switch month {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if isLeapYear(year) {
			return 29
		} else {
			return 28
		}
	}

	panic(fmt.Sprintf("could not get number of days for month number %d, allowed values are from 1 to 12", month))
}
