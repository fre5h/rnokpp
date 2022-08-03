package generator

import (
	"math/rand"
	"time"
)

// GenerateRandomTime generates random datetime in the interval 01.01.1900..current time
func GenerateRandomTime() time.Time {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000 // @todo rewrite method

	return time.Unix(randomTime, 0)
}
