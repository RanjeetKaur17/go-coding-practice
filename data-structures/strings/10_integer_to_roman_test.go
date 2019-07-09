package strings

import "testing"

func Test_IntToRoman(t *testing.T) {
	s := IntToRoman(1)

	if s != "I" {
		t.Fail()
	}

	s = IntToRoman(14)

	if s != "XIV" {
		t.Fail()
	}

	s = IntToRoman(3999)

	if s != "MMMCMXCIX" {
		t.Fail()
	}
}
