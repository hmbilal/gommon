package cast

import "testing"

func TestToString(t *testing.T) {
	tests := []struct {
		input     interface{}
		expected  string
		precision int
	}{
		{"abc", "abc", 0},
		{1, "1", 0},
		{2.3, "2", 0},
		{2.3, "2.3", 1},
		{3.544444444444444444444, "4", 0},
		{3.544444444444444444444, "3.544", 3},
		{3.20, "3.2", 1},
		{3.20, "3.20", 2},
		{3.21, "3.21", 2},
		{float32(3.21), "3.21", 2},
		{1, "1", 0},
		{uint(1), "1", 0},
		{int8(1), "1", 0},
		{uint8(1), "1", 0},
		{false, "false", 0},
	}

	for _, test := range tests {
		actual := ToString(test.input, test.precision)
		if test.expected != actual {
			t.Errorf("ToString(%v, %d)=%s, expected:%s", test.input, test.precision, actual, test.expected)
		}
	}
}
