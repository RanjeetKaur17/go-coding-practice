package strings

import (
	"testing"
)

func Test_CountAndSay(t *testing.T) {
	s := CountAndSay(1)

	if s != "1" {
		t.Fail()
	}

	s = CountAndSay(2)

	if s != "11" {
		t.Fail()
	}

	s = CountAndSay(3)

	if s != "21" {
		t.Fail()
	}

	s = CountAndSay(4)

	if s != "1211" {
		t.Fail()
	}
}
