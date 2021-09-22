package helpers

import (
	"math/rand"
	"time"
)

// The symbols team codes are created from.
// Confusing letters such as I and L, O and Q have one pair removed.
var symbols = []rune("ABCDEFGHJKLMNPRSTUVWXYZ")

// NewCode generates an alpha string of easily recognisable characters
func NewCode(length int) string {
	b := make([]rune, length)
	rand.Seed(time.Now().UnixNano() & 999999999)
	for i := 0; i < length; i++ {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}
