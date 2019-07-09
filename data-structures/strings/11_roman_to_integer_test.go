package strings

import "testing"

func Test_RomanToInt(t *testing.T) {
	s := RomanToInt("I")

	if s != 1 {
		t.Fail()
	}

	s = RomanToInt("XIV")

	if s != 14 {
		t.Fail()
	}

	s = RomanToInt("MMMCMXCIX")

	if s != 3999 {
		t.Fail()
	}
}
