package internal_test

import (
	"testing"
	"time"

	"github.com/fre5h/rnokpp/internal"
)

const reset = "\033[0m"
const red = "\033[31m"

func TestGenerateRandomDate(t *testing.T) {
	for i := 0; i < 1000; i++ {
		date := internal.GenerateRandomDate()
		if date.Before(internal.BaseDate) || date.After(time.Now()) {
			t.Errorf("Invalid generated date %s\"%s\"%s", red, date.Format("02.01.2006"), reset)
		}
	}
}
