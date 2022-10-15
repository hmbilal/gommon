package random

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRandom_String(t *testing.T) {
	assert.Len(t, String(32), 32)
	r := New()
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), r.String(8, Numeric))
	assert.Regexp(t, regexp.MustCompile("[A-Z]+$"), r.String(3, Uppercase))
	assert.Regexp(t, regexp.MustCompile("[a-z]+$"), r.String(18, Lowercase))
	assert.Regexp(t, regexp.MustCompile("[a-zA-Z]+$"), r.String(25, Alphabetic))
	assert.Regexp(t, regexp.MustCompile("[a-zA-Z0-9]+$"), r.String(6, Alphanumeric))
}
