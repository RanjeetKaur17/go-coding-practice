package hashing

import (
	"testing"
)

func Test_FindSubstring(t *testing.T) {
	res := FindSubstring("", []string{"", ""})

	if len(res) != 0 {
		t.Fail()
	}

	res = FindSubstring("barfoothefoobarman", []string{"foo", "bar"})

	if len(res) != 2 || res[0] != 0 || res[1] != 9 {
		t.Fail()
	}

	res = FindSubstring("abbaccaaabcabbbccbabbccabbacabcacbbaabbbbbaaabaccaacbccabcbababbbabccabacbbcabbaacaccccbaabcabaabaaaabcaabcacabaa", []string{"cac", "aaa", "aba", "aab", "abc"})

	if len(res) != 1 || res[0] != 97 {
		t.Fail()
	}
}