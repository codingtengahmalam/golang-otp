package helper

import (
	"math/rand"
	"time"
)

const (
	Numeric = "0123456789"
)

func RandNumeric(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = Numeric[rand.Intn(len(Numeric))]
	}

	return string(b)
}
