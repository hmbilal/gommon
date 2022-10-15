package random

import (
	"math/rand"
	"strings"
	"time"
)

type Random struct {
}

const Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Lowercase = "abcdefghijklmnopqrstuvwxyz"
const Alphabetic = Uppercase + Lowercase
const Numeric = "0123456789"
const Alphanumeric = Alphabetic + Numeric
const Symbols = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
const Hex = Numeric + "abcdef"

func New() *Random {
	rand.Seed(time.Now().UnixNano())
	return new(Random)
}

func (r *Random) String(length uint8, charsets ...string) string {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Int63()%int64(len(charset))]
	}

	return string(b)
}

func String(length uint8, charsets ...string) string {
	return New().String(length, charsets...)
}
